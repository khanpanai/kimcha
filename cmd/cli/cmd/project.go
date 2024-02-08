package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kimcha/cmd/cli/cmd/project"
)

var projectGroup = &cobra.Group{
	ID:    "project",
	Title: "project",
}

var projectCmd = &cobra.Command{
	Use: "project",
}

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project list")
	},
	GroupID: "project",
}

var projectSetActiveCmd = &cobra.Command{
	Use: "sa",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project set active")
	},
	GroupID: "project",
}

var projectRenameCmd = &cobra.Command{
	Use: "rename",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project set active")
	},
	GroupID: "project",
}

var projectDeleteCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project set active")
	},
	GroupID: "project",
}

func init() {
	rootCmd.AddCommand(projectCmd)

	projectCmd.AddGroup(projectGroup)
	projectCmd.AddCommand(projectListCmd)
	projectCmd.AddCommand(project.CreateCmd)
	projectCmd.AddCommand(projectRenameCmd)
	projectCmd.AddCommand(projectDeleteCmd)
	projectCmd.AddCommand(projectSetActiveCmd)

	commandsWithUlidFlag := []*cobra.Command{
		projectRenameCmd,
		projectDeleteCmd,
		projectSetActiveCmd,
	}

	for _, command := range commandsWithUlidFlag {
		command.PersistentFlags().StringP("ulid", "u", "", "ULID of project")
		_ = command.MarkPersistentFlagRequired("ulid")
	}

}
