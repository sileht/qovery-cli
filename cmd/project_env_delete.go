package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"qovery.go/api"
	"qovery.go/util"
)

var projectEnvDeleteCmd = &cobra.Command{
	Use:   "delete <key>",
	Short: "Delete environment variable",
	Long: `DELETE an environment variable to a project. For example:

	qovery project env delete`,
	Run: func(cmd *cobra.Command, args []string) {
		if !hasFlagChanged(cmd) {
			ProjectName = util.CurrentQoveryYML().Application.Project

			if ProjectName == "" {
				fmt.Println("The current directory is not a Qovery project (-h for help)")
				os.Exit(1)
			}
		}

		if len(args) != 1 {
			_ = cmd.Help()
			return
		}

		p := api.GetProjectByName(ProjectName)
		ev := api.ListProjectEnvironmentVariables(p.Id).GetEnvironmentVariableByKey(args[0])
		api.DeleteProjectEnvironmentVariable(ev.Id, p.Id)
		fmt.Println("ok")
	},
}

func init() {
	projectEnvDeleteCmd.PersistentFlags().StringVarP(&ProjectName, "project", "p", "", "Your project name")

	projectEnvCmd.AddCommand(projectEnvDeleteCmd)
}
