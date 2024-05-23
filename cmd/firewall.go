package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/uthoplatforms/utho-cli/helper"
	"github.com/uthoplatforms/utho-go/utho"
)

var firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Use this command to manage firewalls.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createFirewallCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a firewall.",
	Example: "uthoctl firewall create <firewall-name>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		params := utho.CreateFirewallParams{
			Name: args[0],
		}
		firewall, err := client.Firewall().Create(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Firewall Name", "Firewall Id", "Status")
		tbl.AddRow(args[0], firewall.ID, firewall.Status)
		tbl.Print()
	},
}

var getFirewallCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get firewall info",
	Example: "uthoctl firewall get <firewall-id>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		firewall, err := client.Firewall().Read(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "CreatedAt", "Rulecount", "Serverscount")
		tbl.AddRow(firewall.ID, firewall.Name, firewall.CreatedAt, firewall.Rulecount, firewall.Serverscount)
		tbl.Print()
	},
}

var listFirewallCmd = &cobra.Command{
	Use:     "list",
	Short:   "List firewall info",
	Example: "uthoctl firewall list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		firewalls, err := client.Firewall().List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "CreatedAt", "Rulecount", "Serverscount")
		for _, firewall := range firewalls {
			tbl.AddRow(firewall.ID, firewall.Name, firewall.CreatedAt, firewall.Rulecount, firewall.Serverscount)
		}
		tbl.Print()
	},
}

var deleteFirewallCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a firewall from your account.",
	Example: "uthoctl firewall delete <firewall-id>",
	Args:    cobra.ExactArgs(2),
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

		firewall, err := client.Firewall().Delete(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + firewall.Status)
	},
}

// Firewallrule
var firewallruleCmd = &cobra.Command{
	Use:   "firewallrule",
	Short: "Use this command to manage firewall rules.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createFirewallruleCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a firewall rule.",
	Example: "uthoctl firewall firewallrule create <firewall-id>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		firewallRuleType, _ := cmd.Flags().GetString("type")
		service, _ := cmd.Flags().GetString("service")
		protocol, _ := cmd.Flags().GetString("protocol")
		port, _ := cmd.Flags().GetString("port")
		addresses, _ := cmd.Flags().GetString("addresses")
		params := utho.CreateFirewallRuleParams{
			FirewallId: args[0],
			Type:       firewallRuleType,
			Service:    service,
			Protocol:   protocol,
			Port:       port,
			Addresses:  addresses,
		}
		firewallrule, err := client.Firewall().CreateFirewallRule(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Firewallrule Id", "Status")
		tbl.AddRow(firewallrule.ID, firewallrule.Status)
		tbl.Print()
	},
}

var getFirewallruleCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get firewall firewallrule info",
	Example: "uthoctl firewall firewallrule get <firewall-id> <firewallrule-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		firewallrule, err := client.Firewall().ReadFirewallRule(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Firewallid", "Type", "Service", "Protocol", "Port", "Addresses")
		tbl.AddRow(firewallrule.ID, firewallrule.Firewallid, firewallrule.Type, firewallrule.Service, firewallrule.Protocol, firewallrule.Port, firewallrule.Addresses)
		tbl.Print()
	},
}

var listFirewallruleCmd = &cobra.Command{
	Use:     "list",
	Short:   "List firewall Firewallrule",
	Example: "uthoctl firewall firewallrule list <firewall-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		firewallrules, err := client.Firewall().ListFirewallRules(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Firewallid", "Type", "Service", "Protocol", "Port", "Addresses")
		for _, firewallrule := range firewallrules {
			tbl.AddRow(firewallrule.ID, firewallrule.Firewallid, firewallrule.Type, firewallrule.Service, firewallrule.Protocol, firewallrule.Port, firewallrule.Addresses)
		}
		tbl.Print()
	},
}

var deleteFirewallruleCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete an firewall policy from your account.",
	Example: "uthoctl firewall firewallrule delete <firewall-id> <firewallrule-id>",
	Args:    cobra.ExactArgs(2),
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

		firewallrule, err := client.Firewall().DeleteFirewallRule(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + firewallrule.Status)
	},
}

func init() {
	rootCmd.AddCommand(firewallCmd)
	// Firewall
	firewallCmd.AddCommand(createFirewallCmd)
	firewallCmd.AddCommand(getFirewallCmd)
	firewallCmd.AddCommand(listFirewallCmd)
	firewallCmd.AddCommand(deleteFirewallCmd)

	// Firewall Rule
	firewallCmd.AddCommand(firewallruleCmd)
	firewallruleCmd.AddCommand(createFirewallruleCmd)
	createFirewallruleCmd.Flags().String("type", "", "Incoming or outgoing traffic eg: incoming, outgonig")
	createFirewallruleCmd.Flags().String("service", "", "")
	createFirewallruleCmd.Flags().String("protocol", "", "The type of traffic to be allowed. This may be one of 'tcp', 'udp', or 'icmp'")
	createFirewallruleCmd.Flags().String("port", "", "The ports on which traffic will be allowed specified as a string containing a single port, a range (e.g. '8000-9000'), or 'ALL' to open all ports for a protocol. ")
	createFirewallruleCmd.Flags().String("addresses", "", "An array of strings containing the IPv4 addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the Firewall will allow traffic")

	firewallruleCmd.AddCommand(getFirewallruleCmd)
	firewallruleCmd.AddCommand(listFirewallruleCmd)
	firewallruleCmd.AddCommand(deleteFirewallruleCmd)
}
