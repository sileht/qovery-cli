package cmd

import (
	"fmt"
	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
	"os"
	"qovery.go/api"
	"qovery.go/util"
	"strconv"
)

var storageListCmd = &cobra.Command{
	Use:   "list",
	Short: "List storage",
	Long: `LIST show all available storage within a project and environment. For example:

	qovery storage list`,
	Run: func(cmd *cobra.Command, args []string) {
		if !hasFlagChanged(cmd) {
			BranchName = util.CurrentBranchName()
			ProjectName = util.CurrentQoveryYML().Application.Project

			if BranchName == "" || ProjectName == "" {
				fmt.Println("The current directory is not a Qovery project (-h for help)")
				os.Exit(0)
			}
		}

		ShowStorageList(ProjectName, BranchName)
	},
}

func init() {
	storageListCmd.PersistentFlags().StringVarP(&ProjectName, "project", "p", "", "Your project name")
	storageListCmd.PersistentFlags().StringVarP(&BranchName, "branch", "b", "", "Your branch name")

	storageCmd.AddCommand(storageListCmd)
}

func ShowStorageList(projectName string, branchName string) {
	output := []string{
		"name | status | type | version | endpoint | port | username | password | application",
	}

	services := api.ListStorage(api.GetProjectByName(projectName).Id, branchName)

	if services.Results == nil || len(services.Results) == 0 {
		fmt.Println(columnize.SimpleFormat(output))
		return
	}

	for _, a := range services.Results {
		applicationName := "none"

		if a.Application != nil {
			applicationName = a.Application.Name
		}

		output = append(output, a.Name+" | "+a.Status.CodeMessage+
			" | "+a.Type+" | "+a.Version+" | "+a.FQDN+" | "+strconv.Itoa(*a.Port)+
			" | "+a.Username+" | "+a.Password+" | "+applicationName)
	}

	fmt.Println(columnize.SimpleFormat(output))
}
