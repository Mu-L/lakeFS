// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: actions/actions.proto

package actions

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// message data model for RunResult struct
type RunResultData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RunId         string                 `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	BranchId      string                 `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	CommitId      string                 `protobuf:"bytes,3,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	SourceRef     string                 `protobuf:"bytes,4,opt,name=source_ref,json=sourceRef,proto3" json:"source_ref,omitempty"`
	EventType     string                 `protobuf:"bytes,5,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	StartTime     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Passed        bool                   `protobuf:"varint,8,opt,name=passed,proto3" json:"passed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RunResultData) Reset() {
	*x = RunResultData{}
	mi := &file_actions_actions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RunResultData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResultData) ProtoMessage() {}

func (x *RunResultData) ProtoReflect() protoreflect.Message {
	mi := &file_actions_actions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResultData.ProtoReflect.Descriptor instead.
func (*RunResultData) Descriptor() ([]byte, []int) {
	return file_actions_actions_proto_rawDescGZIP(), []int{0}
}

func (x *RunResultData) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *RunResultData) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *RunResultData) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *RunResultData) GetSourceRef() string {
	if x != nil {
		return x.SourceRef
	}
	return ""
}

func (x *RunResultData) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *RunResultData) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *RunResultData) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *RunResultData) GetPassed() bool {
	if x != nil {
		return x.Passed
	}
	return false
}

// message data model for TaskResult struct
type TaskResultData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RunId         string                 `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	HookRunId     string                 `protobuf:"bytes,2,opt,name=hook_run_id,json=hookRunId,proto3" json:"hook_run_id,omitempty"`
	HookId        string                 `protobuf:"bytes,3,opt,name=hook_id,json=hookId,proto3" json:"hook_id,omitempty"`
	ActionName    string                 `protobuf:"bytes,4,opt,name=action_name,json=actionName,proto3" json:"action_name,omitempty"`
	StartTime     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Passed        bool                   `protobuf:"varint,9,opt,name=passed,proto3" json:"passed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskResultData) Reset() {
	*x = TaskResultData{}
	mi := &file_actions_actions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskResultData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResultData) ProtoMessage() {}

func (x *TaskResultData) ProtoReflect() protoreflect.Message {
	mi := &file_actions_actions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResultData.ProtoReflect.Descriptor instead.
func (*TaskResultData) Descriptor() ([]byte, []int) {
	return file_actions_actions_proto_rawDescGZIP(), []int{1}
}

func (x *TaskResultData) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *TaskResultData) GetHookRunId() string {
	if x != nil {
		return x.HookRunId
	}
	return ""
}

func (x *TaskResultData) GetHookId() string {
	if x != nil {
		return x.HookId
	}
	return ""
}

func (x *TaskResultData) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *TaskResultData) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TaskResultData) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *TaskResultData) GetPassed() bool {
	if x != nil {
		return x.Passed
	}
	return false
}

var File_actions_actions_proto protoreflect.FileDescriptor

const file_actions_actions_proto_rawDesc = "" +
	"\n" +
	"\x15actions/actions.proto\x12\x1bio.treeverse.lakefs.actions\x1a\x1fgoogle/protobuf/timestamp.proto\"\xa8\x02\n" +
	"\rRunResultData\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\x12\x1b\n" +
	"\tbranch_id\x18\x02 \x01(\tR\bbranchId\x12\x1b\n" +
	"\tcommit_id\x18\x03 \x01(\tR\bcommitId\x12\x1d\n" +
	"\n" +
	"source_ref\x18\x04 \x01(\tR\tsourceRef\x12\x1d\n" +
	"\n" +
	"event_type\x18\x05 \x01(\tR\teventType\x129\n" +
	"\n" +
	"start_time\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tstartTime\x125\n" +
	"\bend_time\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\aendTime\x12\x16\n" +
	"\x06passed\x18\b \x01(\bR\x06passed\"\x8b\x02\n" +
	"\x0eTaskResultData\x12\x15\n" +
	"\x06run_id\x18\x01 \x01(\tR\x05runId\x12\x1e\n" +
	"\vhook_run_id\x18\x02 \x01(\tR\thookRunId\x12\x17\n" +
	"\ahook_id\x18\x03 \x01(\tR\x06hookId\x12\x1f\n" +
	"\vaction_name\x18\x04 \x01(\tR\n" +
	"actionName\x129\n" +
	"\n" +
	"start_time\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tstartTime\x125\n" +
	"\bend_time\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\aendTime\x12\x16\n" +
	"\x06passed\x18\t \x01(\bR\x06passedB%Z#github.com/treeverse/lakefs/actionsb\x06proto3"

var (
	file_actions_actions_proto_rawDescOnce sync.Once
	file_actions_actions_proto_rawDescData []byte
)

func file_actions_actions_proto_rawDescGZIP() []byte {
	file_actions_actions_proto_rawDescOnce.Do(func() {
		file_actions_actions_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_actions_actions_proto_rawDesc), len(file_actions_actions_proto_rawDesc)))
	})
	return file_actions_actions_proto_rawDescData
}

var file_actions_actions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_actions_actions_proto_goTypes = []any{
	(*RunResultData)(nil),         // 0: io.treeverse.lakefs.actions.RunResultData
	(*TaskResultData)(nil),        // 1: io.treeverse.lakefs.actions.TaskResultData
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_actions_actions_proto_depIdxs = []int32{
	2, // 0: io.treeverse.lakefs.actions.RunResultData.start_time:type_name -> google.protobuf.Timestamp
	2, // 1: io.treeverse.lakefs.actions.RunResultData.end_time:type_name -> google.protobuf.Timestamp
	2, // 2: io.treeverse.lakefs.actions.TaskResultData.start_time:type_name -> google.protobuf.Timestamp
	2, // 3: io.treeverse.lakefs.actions.TaskResultData.end_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_actions_actions_proto_init() }
func file_actions_actions_proto_init() {
	if File_actions_actions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_actions_actions_proto_rawDesc), len(file_actions_actions_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_actions_actions_proto_goTypes,
		DependencyIndexes: file_actions_actions_proto_depIdxs,
		MessageInfos:      file_actions_actions_proto_msgTypes,
	}.Build()
	File_actions_actions_proto = out.File
	file_actions_actions_proto_goTypes = nil
	file_actions_actions_proto_depIdxs = nil
}
