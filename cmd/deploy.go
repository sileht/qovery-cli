package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"qovery.go/api"
	"qovery.go/util"
)

var deployCmd = &cobra.Command{
	Use:   "deploy <commit id>",
	Short: "Perform deploy actions",
	Long: `DEPLOY performs actions on deploy. For example:

	qovery deploy`,
	Run: func(cmd *cobra.Command, args []string) {
		if !hasFlagChanged(cmd) {
			BranchName = util.CurrentBranchName()
			qoveryYML := util.CurrentQoveryYML()
			ProjectName = qoveryYML.Application.Project
			ApplicationName = qoveryYML.Application.Name

			if BranchName == "" || ProjectName == "" || ApplicationName == "" {
				fmt.Println("The current directory is not a Qovery project (-h for help)")
				os.Exit(0)
			}
		}

		if len(args) != 1 {
			_ = cmd.Help()
			return
		}

		commitId := args[0]

		projectId := api.GetProjectByName(ProjectName).Id
		applicationId := api.GetApplicationByName(projectId, BranchName, ApplicationName).Id

		api.Deploy(projectId, BranchName, applicationId, commitId)

		fmt.Println("deployment in progress...")
		fmt.Println("hint: type \"qovery status --watch\" to track the progression of deployment")
	},
}

func init() {
	deployCmd.PersistentFlags().StringVarP(&ApplicationName, "application", "a", "", "Your application name")
	deployCmd.PersistentFlags().StringVarP(&ProjectName, "project", "p", "", "Your project name")
	deployCmd.PersistentFlags().StringVarP(&BranchName, "branch", "b", "", "Your branch name")

	RootCmd.AddCommand(deployCmd)
}
