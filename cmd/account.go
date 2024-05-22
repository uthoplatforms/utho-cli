package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/uthoplatforms/utho-cli/helper"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Get account info",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var getAccountCmd = &cobra.Command{
	Use:   "get",
	Short: "Get account info",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		account, err := client.Account().Read()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Id", "User Email", "Cloud Limit", "Total Cloud Instance", "K8s limit", "Currency", "Available Credit")
		tbl.AddRow(account.ID, account.Email, account.Cloudlimit, account.TotalCloudservers, account.K8SLimit, account.Currency, account.Availablecredit)
		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(getAccountCmd)
}
