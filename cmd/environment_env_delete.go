package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"qovery.go/api"
	"qovery.go/util"
)

var environmentEnvDeleteCmd = &cobra.Command{
	Use:   "delete <key>",
	Short: "Delete environment variable",
	Long: `DELETE an environment variable to an environment. For example:

	qovery environment env delete`,
	Run: func(cmd *cobra.Command, args []string) {
		if !hasFlagChanged(cmd) {
			BranchName = util.CurrentBranchName()
			ProjectName = util.CurrentQoveryYML().Application.Project

			if BranchName == "" || ProjectName == "" {
				fmt.Println("The current directory is not a Qovery project (-h for help)")
				os.Exit(1)
			}
		}

		if len(args) != 1 {
			_ = cmd.Help()
			return
		}

		p := api.GetProjectByName(ProjectName)
		ev := api.ListEnvironmentEnvironmentVariables(p.Id, BranchName).GetEnvironmentVariableByKey(args[0])
		api.DeleteEnvironmentEnvironmentVariable(ev.Id, p.Id, BranchName)
		fmt.Println("ok")
	},
}

func init() {
	environmentEnvDeleteCmd.PersistentFlags().StringVarP(&ProjectName, "project", "p", "", "Your project name")
	environmentEnvDeleteCmd.PersistentFlags().StringVarP(&BranchName, "branch", "b", "", "Your branch name")

	environmentEnvCmd.AddCommand(environmentEnvDeleteCmd)
}
