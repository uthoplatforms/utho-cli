package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/uthoplatforms/utho-cli/helper"
	"github.com/uthoplatforms/utho-go/utho"
)

var vpcCmd = &cobra.Command{
	Use:   "vpc",
	Short: "Use this command to manage VPCs.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createVpcCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a vpc.",
	Example: "uthoctl vpc create <vpc-name>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dcslug, _ := cmd.Flags().GetString("dcslug")
		planid, _ := cmd.Flags().GetString("planid")
		size, _ := cmd.Flags().GetString("size")
		network, _ := cmd.Flags().GetString("network")

		params := utho.CreateVpcParams{
			Name:    args[0],
			Dcslug:  dcslug,
			Planid:  planid,
			Network: network,
			Size:    size,
		}
		vpc, err := client.Vpc().Create(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Vpc Name", "Vpc Id", "Status")
		tbl.AddRow(args[0], vpc.ID, vpc.Status)
		tbl.Print()
	},
}

var getVpcCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get vpc info",
	Example: "uthoctl vpc get <vpc-id>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		vpc, err := client.Vpc().Read(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Network", "Name", "Size", "Dcslug")
		tbl.AddRow(vpc.ID, vpc.Network, vpc.Name, vpc.Size, vpc.Dcslug)
		tbl.Print()
	},
}

var listVpcCmd = &cobra.Command{
	Use:     "list",
	Short:   "List vpc info",
	Example: "uthoctl vpc list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		vpcs, err := client.Vpc().List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Network", "Name", "Size", "Dcslug")
		for _, vpc := range vpcs {
			tbl.AddRow(vpc.ID, vpc.Network, vpc.Name, vpc.Size, vpc.Dcslug)
		}
		tbl.Print()
	},
}

var deleteVpcCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a vpc from your account.",
	Example: "uthoctl vpc delete <vpc-id>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		confirm := helper.Ask()
		if !confirm {
			fmt.Println("Operation aborted.")
			os.Exit(1)
		}

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		vpc, err := client.Vpc().Delete(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + vpc.Status)
	},
}

func init() {
	rootCmd.AddCommand(vpcCmd)
	// Vpc
	vpcCmd.AddCommand(createVpcCmd)
	createVpcCmd.Flags().String("dcslug", "", "Provide Zone dcslug eg: innoida")
	createVpcCmd.Flags().String("Billing", "", "")
	createVpcCmd.Flags().String("Size", "", "")
	createVpcCmd.Flags().String("Price", "", "")

	vpcCmd.AddCommand(getVpcCmd)
	vpcCmd.AddCommand(listVpcCmd)
	vpcCmd.AddCommand(deleteVpcCmd)
}
