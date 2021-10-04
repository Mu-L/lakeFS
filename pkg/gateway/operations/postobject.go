package operations

import (
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/treeverse/lakefs/pkg/gateway/multiparts"

	"github.com/google/uuid"
	"github.com/treeverse/lakefs/pkg/block"
	gatewayErrors "github.com/treeverse/lakefs/pkg/gateway/errors"
	"github.com/treeverse/lakefs/pkg/gateway/path"
	"github.com/treeverse/lakefs/pkg/gateway/serde"
	"github.com/treeverse/lakefs/pkg/graveler"
	"github.com/treeverse/lakefs/pkg/httputil"
	"github.com/treeverse/lakefs/pkg/logging"
	"github.com/treeverse/lakefs/pkg/permissions"
)

const (
	CreateMultipartUploadQueryParam   = "uploads"
	CompleteMultipartUploadQueryParam = "uploadId"
)

type PostObject struct{}

func (controller *PostObject) RequiredPermissions(_ *http.Request, repoID, _, path string) ([]permissions.Permission, error) {
	return []permissions.Permission{
		{
			Action:   permissions.WriteObjectAction,
			Resource: permissions.ObjectArn(repoID, path),
		},
	}, nil
}

func (controller *PostObject) HandleCreateMultipartUpload(w http.ResponseWriter, req *http.Request, o *PathOperation) {
	o.Incr("create_mpu")
	uuidBytes := [16]byte(uuid.New())
	objName := hex.EncodeToString(uuidBytes[:])
	storageClass := StorageClassFromHeader(req.Header)
	opts := block.CreateMultiPartUploadOpts{StorageClass: storageClass}
	uploadID, err := o.BlockStore.CreateMultiPartUpload(req.Context(), block.ObjectPointer{StorageNamespace: o.Repository.StorageNamespace, Identifier: objName}, req, opts)
	if err != nil {
		o.Log(req).WithError(err).Error("could not create multipart upload")
		_ = o.EncodeError(w, req, gatewayErrors.Codes.ToAPIErr(gatewayErrors.ErrInternalError))
		return
	}
	metadata := multiparts.Metadata{}
	metadata.Set("Content-Type", req.Header.Get("Content-Type"))
	err = o.MultipartsTracker.Create(req.Context(), uploadID, o.Path, objName, time.Now(), metadata)
	if err != nil {
		o.Log(req).WithError(err).Error("could not write multipart upload to DB")
		_ = o.EncodeError(w, req, gatewayErrors.Codes.ToAPIErr(gatewayErrors.ErrInternalError))
		return
	}
	o.EncodeResponse(w, req, &serde.InitiateMultipartUploadResult{
		Bucket:   o.Repository.Name,
		Key:      path.WithRef(o.Path, o.Reference),
		UploadID: uploadID,
	}, http.StatusOK)
}

func trimQuotes(s string) string {
	return strings.Trim(s, "\"")
}

func (controller *PostObject) HandleCompleteMultipartUpload(w http.ResponseWriter, req *http.Request, o *PathOperation) {
	var etag *string
	var size int64
	o.Incr("complete_mpu")
	uploadID := req.URL.Query().Get(CompleteMultipartUploadQueryParam)
	req = req.WithContext(logging.AddFields(req.Context(), logging.Fields{logging.UploadIDFieldKey: uploadID}))
	multiPart, err := o.MultipartsTracker.Get(req.Context(), uploadID)
	if err != nil {
		o.Log(req).WithError(err).Error("could not read multipart record")
		_ = o.EncodeError(w, req, gatewayErrors.Codes.ToAPIErr(gatewayErrors.ErrInternalError))
		return
	}
	objName := multiPart.PhysicalAddress
	req = req.WithContext(logging.AddFields(req.Context(), logging.Fields{logging.PhysicalAddressFieldKey: objName}))
	xmlMultipartComplete, err := ioutil.ReadAll(req.Body)
	if err != nil {
		o.Log(req).WithError(err).Error("could not read request body")
		_ = o.EncodeError(w, req, gatewayErrors.Codes.ToAPIErr(gatewayErrors.ErrInternalError))
		return
	}
	var multipartList block.MultipartUploadCompletion
	err = xml.Unmarshal(xmlMultipartComplete, &multipartList)
	if err != nil {
		o.Log(req).WithError(err).Error("could not parse multipart XML on complete multipart")
		_ = o.EncodeError(w, req, gatewayErrors.Codes.ToAPIErr(gatewayErrors.ErrInternalError))
		return
	}
	etag, size, err = o.BlockStore.CompleteMultiPartUpload(req.Context(), block.ObjectPointer{
		StorageNamespace: o.Repository.StorageNamespace,
		Identifier:       objName,
	}, uploadID, &multipartList)
	if err != nil {
		o.Log(req).WithError(err).Error("could not complete multipart upload")
		_ = o.EncodeError(w, req, gatewayErrors.Codes.ToAPIErr(gatewayErrors.ErrInternalError))
		return
	}
	ch := trimQuotes(*etag)
	checksum := strings.Split(ch, "-")[0]
	err = o.finishUpload(req, checksum, objName, size, true, multiPart.Metadata)
	if errors.Is(err, graveler.ErrWriteToProtectedBranch) {
		_ = o.EncodeError(w, req, gatewayErrors.Codes.ToAPIErr(gatewayErrors.ErrWriteToProtectedBranch))
		return
	}
	if err != nil {
		_ = o.EncodeError(w, req, gatewayErrors.Codes.ToAPIErr(gatewayErrors.ErrInternalError))
		return
	}
	err = o.MultipartsTracker.Delete(req.Context(), uploadID)
	if err != nil {
		o.Log(req).WithError(err).Warn("could not delete multipart record")
	}

	scheme := httputil.RequestScheme(req)
	var location string
	if o.MatchedHost {
		location = fmt.Sprintf("%s://%s/%s/%s", scheme, req.Host, o.Reference, o.Path)
	} else {
		location = fmt.Sprintf("%s://%s/%s/%s/%s", scheme, req.Host, o.Repository.Name, o.Reference, o.Path)
	}
	o.EncodeResponse(w, req, &serde.CompleteMultipartUploadResult{
		Location: location,
		Bucket:   o.Repository.Name,
		Key:      path.WithRef(o.Path, o.Reference),
		ETag:     *etag,
	}, http.StatusOK)
}

func (controller *PostObject) Handle(w http.ResponseWriter, req *http.Request, o *PathOperation) {
	// POST is only supported for CreateMultipartUpload/CompleteMultipartUpload
	// https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
	// https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
	_, mpuCreateParamExist := req.URL.Query()[CreateMultipartUploadQueryParam]
	if mpuCreateParamExist {
		controller.HandleCreateMultipartUpload(w, req, o)
		return
	}

	_, mpuCompleteParamExist := req.URL.Query()[CompleteMultipartUploadQueryParam]
	if mpuCompleteParamExist {
		controller.HandleCompleteMultipartUpload(w, req, o)
		return
	}
	// otherwise
	w.WriteHeader(http.StatusMethodNotAllowed)
}
