package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/uthoplatforms/utho-cli/helper"
)

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Get action info",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var listActionCmd = &cobra.Command{
	Use:   "list",
	Short: "List action info",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		actions, err := client.Action().List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Action", "ResourceType", "ResourceID", "StartedAt", "CompletedAt", "Process")
		for _, action := range actions {
			tbl.AddRow(action.ID, action.Action, action.ResourceType, action.ResourceID, action.StartedAt, action.CompletedAt, action.Process)
		}
		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(actionCmd)
	actionCmd.AddCommand(listActionCmd)
}
