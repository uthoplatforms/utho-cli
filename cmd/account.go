package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/uthoplatforms/utho-go/utho"
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
		token := viper.GetString("token")
		if token == "" {
			fmt.Println("No token found. Please login first.")
			os.Exit(1)
		}
		account, err := getAccount(token)
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

func getAccount(token string) (*utho.User, error) {
	clinet, err := utho.NewClient(token)
	if err != nil {
		return nil, err
	}

	return clinet.Account().Read()
}
