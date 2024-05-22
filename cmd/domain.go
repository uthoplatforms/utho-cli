package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/uthoplatforms/utho-cli/helper"
	"github.com/uthoplatforms/utho-go/utho"
)

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Use this command to manage domains you have purchased from a domain name registrar that you are managing through the Utho DNS interface.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createDomainCmd = &cobra.Command{
	Use:   "create",
	Short: "Adds a domain to your account.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		domain, err := client.Domain().CreateDomain(utho.CreateDomainParams{Domain: args[0]})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Domain", "Status")
		tbl.AddRow(args[0], domain.Status)
		tbl.Print()
	},
}

var getDomainCmd = &cobra.Command{
	Use:   "get",
	Short: "Get domain info",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		domain, err := client.Domain().ReadDomain(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Domain", "Dns Record Count", "Created At")
		tbl.AddRow(domain.Domain, domain.DnsrecordCount, domain.CreatedAt)
		tbl.Print()
	},
}

var listDomainCmd = &cobra.Command{
	Use:   "list",
	Short: "List domain info",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		domains, err := client.Domain().ListDomains()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Domain", "Dns Record Count", "Created At")
		for _, domain := range domains {
			tbl.AddRow(domain.Domain, domain.DnsrecordCount, domain.CreatedAt)
		}
		tbl.Print()
	},
}

var deleteDomainCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a domain from your account.",
	Args:  cobra.ExactArgs(1),
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

		domain, err := client.Domain().DeleteDomain(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Status")
		tbl.AddRow(domain.Status)
		tbl.Print()
	},
}

var dnsCmd = &cobra.Command{
	Use:   "records",
	Short: "Use this command to to manage the DNS records for your domains.",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createDomainRecordCmd = &cobra.Command{
	Use:   "create",
	Short: "Adds a record to your domain.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		recordType, _ := cmd.Flags().GetString("type")
		hostname, _ := cmd.Flags().GetString("hostname")
		value, _ := cmd.Flags().GetString("value")
		ttl, _ := cmd.Flags().GetString("ttl")
		porttype, _ := cmd.Flags().GetString("port-type")
		port, _ := cmd.Flags().GetString("port")
		priority, _ := cmd.Flags().GetString("priority")
		wight, _ := cmd.Flags().GetString("wight")

		params := utho.CreateDnsRecordParams{
			Domain:   args[0],
			Type:     recordType,
			Hostname: hostname,
			Value:    value,
			TTL:      ttl,
			Porttype: porttype,
			Port:     port,
			Priority: priority,
			Wight:    wight,
		}
		record, err := client.Domain().CreateDnsRecord(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Domain", "Record Id", "Record Name", "Status")
		tbl.AddRow(args[0], record.ID, hostname, record.Status)
		tbl.Print()
	},
}

var listDomainRecordCmd = &cobra.Command{
	Use:     "list",
	Short:   "List domain Record",
	Example: "uthoctl domain records list <domain>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dnsRecords, err := client.Domain().ListDnsRecords(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("id", "hostname", "type", "value", "ttl", "priority")
		for _, record := range dnsRecords {
			tbl.AddRow(record.ID, record.Hostname, record.Type, record.Value, record.TTL, record.Priority)
		}
		tbl.Print()
	},
}

var deleteDomainRecordCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a domain record from your account.",
	Example: "uthoctl domain records delete <domain> <record-id>",
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

		record, err := client.Domain().DeleteDnsRecord(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + record.Status)
	},
}

func init() {
	rootCmd.AddCommand(domainCmd)
	domainCmd.AddCommand(createDomainCmd)
	domainCmd.AddCommand(getDomainCmd)
	domainCmd.AddCommand(listDomainCmd)
	domainCmd.AddCommand(deleteDomainCmd)
	domainCmd.AddCommand(dnsCmd)

	dnsCmd.AddCommand(createDomainRecordCmd)
	createDomainRecordCmd.Flags().String("type", "", "The Record Type (A, AAAA, CAA, CNAME, MX, TXT, SRV, NS)")
	createDomainRecordCmd.Flags().String("hostname", "", "Name (Hostname) The host name, alias, or service being defined by the record.")
	createDomainRecordCmd.Flags().String("value", "", "Variable data depending on record type. For example, the value for an A record would be the IPv4 address to which the domain will be mapped. For a CAA record, it would contain the domain name of the CA being granted permission to issue certificates.")
	createDomainRecordCmd.Flags().String("ttl", "", "This value is the time to live for the record, in seconds. This defines the time frame that clients can cache queried information before a refresh should be requested. If not set, the default value is 1800")
	createDomainRecordCmd.Flags().String("port-type", "", "The port that the service is accessible on (for SRV records only. null otherwise).")
	createDomainRecordCmd.Flags().String("port", "", "Port")
	createDomainRecordCmd.Flags().String("priority", "", "The priority of the host (for SRV and MX records. null otherwise). ")
	createDomainRecordCmd.Flags().String("wight", "", "The weight of records with the same priority (for SRV records only. null otherwise).")

	dnsCmd.AddCommand(listDomainRecordCmd)
	dnsCmd.AddCommand(deleteDomainRecordCmd)
}
