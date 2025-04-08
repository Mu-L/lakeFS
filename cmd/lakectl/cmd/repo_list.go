package cmd

import (
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/treeverse/lakefs/pkg/api/apigen"
	"github.com/treeverse/lakefs/pkg/api/apiutil"
)

var repoListCmd = &cobra.Command{
	Use:   "list",
	Short: "List repositories",
	Args:  cobra.NoArgs,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		prefix, after, amount := getPaginationFlags(cmd)
		clt := getClient()

		resp, err := clt.ListRepositoriesWithResponse(cmd.Context(), &apigen.ListRepositoriesParams{
			Prefix: apiutil.Ptr(apigen.PaginationPrefix(prefix)),
			After:  apiutil.Ptr(apigen.PaginationAfter(after)),
			Amount: apiutil.Ptr(apigen.PaginationAmount(amount)),
		})
		DieOnErrorOrUnexpectedStatusCode(resp, err, http.StatusOK)
		if resp.JSON200 == nil {
			Die("Bad response from server", 1)
		}
		repos := resp.JSON200.Results
		rows := make([][]interface{}, len(repos))
		for i, repo := range repos {
			ts := time.Unix(repo.CreationDate, 0).String()
			rows[i] = []interface{}{repo.Id, ts, repo.DefaultBranch, repo.Id, repo.StorageNamespace}
		}
		pagination := resp.JSON200.Pagination
		PrintTable(rows, []interface{}{"Repository", "Creation Date", "Default Ref Name", "Storage ID", "Storage Namespace"}, &pagination, amount)
	},
}

//nolint:gochecknoinits
func init() {
	withPaginationFlags(repoListCmd)
	repoCmd.AddCommand(repoListCmd)
}
