package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/uthoplatforms/utho-cli/helper"
	"github.com/uthoplatforms/utho-go/utho"
)

var targetgroupCmd = &cobra.Command{
	Use:   "targetgroup",
	Short: "Use this command to manage object storages.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createTargetgroupCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a targetgroup.",
	Example: "uthoctl targetgroup create <targetgroup-name>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		protocol, _ := cmd.Flags().GetString("protocol")
		port, _ := cmd.Flags().GetString("port")
		healthCheckPath, _ := cmd.Flags().GetString("health_check_path")
		healthCheckProtocol, _ := cmd.Flags().GetString("health_check_protocol")
		healthCheckInterval, _ := cmd.Flags().GetString("health_check_interval")
		healthCheckTimeout, _ := cmd.Flags().GetString("health_check_timeout")
		healthyThreshold, _ := cmd.Flags().GetString("healthy_threshold")
		unhealthyThreshold, _ := cmd.Flags().GetString("unhealthy_threshold")

		params := utho.CreateTargetGroupParams{
			Name:                args[0],
			Protocol:            protocol,
			Port:                port,
			HealthCheckPath:     healthCheckPath,
			HealthCheckProtocol: healthCheckProtocol,
			HealthCheckInterval: healthCheckInterval,
			HealthCheckTimeout:  healthCheckTimeout,
			HealthyThreshold:    healthyThreshold,
			UnhealthyThreshold:  unhealthyThreshold,
		}
		targetgroup, err := client.TargetGroup().Create(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Targetgroup Name", "Targetgroup Id", "Status")
		tbl.AddRow(args[0], targetgroup.ID, targetgroup.Status)
		tbl.Print()
	},
}

var getTargetgroupCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get targetgroup info",
	Example: "uthoctl targetgroup get <targetgroup-id>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		targetgroup, err := client.TargetGroup().Read(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "Port", "Protocol", "HealthCheckPath")
		tbl.AddRow(targetgroup.ID, targetgroup.Name, targetgroup.Port, targetgroup.Protocol, targetgroup.HealthCheckPath)
		tbl.Print()
	},
}

var listTargetgroupCmd = &cobra.Command{
	Use:     "list",
	Short:   "List targetgroup info",
	Example: "uthoctl targetgroup list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		targetgroups, err := client.TargetGroup().List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "Port", "Protocol", "HealthCheckPath")
		for _, targetgroup := range targetgroups {
			tbl.AddRow(targetgroup.ID, targetgroup.Name, targetgroup.Port, targetgroup.Protocol, targetgroup.HealthCheckPath)
		}
		tbl.Print()
	},
}

var deleteTargetgroupCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a targetgroup from your account.",
	Example: "uthoctl targetgroup delete <targetgroup-id> <targetgroup-name>",
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

		targetgroup, err := client.TargetGroup().Delete(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + targetgroup.Status)
	},
}

// Target
var targetgroupTargetCmd = &cobra.Command{
	Use:   "target",
	Short: "Use this command to manage targetgroup Security Group.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createTargetgroupTargetCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a targetgroup target.",
	Example: "uthoctl targetgroup target create <targetgroup-id>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		backendProtocol, _ := cmd.Flags().GetString("backend_protocol")
		backendPort, _ := cmd.Flags().GetString("backend_port")
		ip, _ := cmd.Flags().GetString("ip")
		cloudid, _ := cmd.Flags().GetString("cloudid")

		params := utho.CreateTargetGroupTargetParams{
			TargetGroupId:   args[0],
			BackendProtocol: backendProtocol,
			BackendPort:     backendPort,
			IP:              ip,
			Cloudid:         cloudid,
		}
		target, err := client.TargetGroup().CreateTarget(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Target Name", "Target Id", "Status")
		tbl.AddRow(args[1], target.ID, target.Status)
		tbl.Print()
	},
}

var getTargetgroupTargetCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get targetgroup target info",
	Example: "uthoctl targetgroup target get <targetgroup-id> <target-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		target, err := client.TargetGroup().ReadTarget(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("IP", "Cloudid", "Status", "ID")
		tbl.AddRow(target.IP, target.Cloudid, target.Status, target.ID)
		tbl.Print()
	},
}

var listTargetgroupTargetCmd = &cobra.Command{
	Use:     "list",
	Short:   "List targetgroup policy",
	Example: "uthoctl targetgroup policy list <targetgroup-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		targets, err := client.TargetGroup().ListTargets(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("IP", "Cloudid", "Status", "ID")
		for _, target := range targets {
			tbl.AddRow(target.IP, target.Cloudid, target.Status, target.ID)
		}
		tbl.Print()
	},
}

var deleteTargetgroupTargetCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a targetgroup Target from your account.",
	Example: "uthoctl targetgroup target delete <targetgroup-id> <target-id>",
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

		target, err := client.TargetGroup().DeleteTarget(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + target.Status)
	},
}

func init() {
	rootCmd.AddCommand(targetgroupCmd)
	// Targetgroup
	targetgroupCmd.AddCommand(createTargetgroupCmd)
	createTargetgroupCmd.Flags().String("protocol", "", "")
	createTargetgroupCmd.Flags().String("port", "", "")
	createTargetgroupCmd.Flags().String("health_check_path", "", "")
	createTargetgroupCmd.Flags().String("health_check_protocol", "", "")
	createTargetgroupCmd.Flags().String("health_check_interval", "", "")
	createTargetgroupCmd.Flags().String("health_check_timeout", "", "")
	createTargetgroupCmd.Flags().String("healthy_threshold", "", "")
	createTargetgroupCmd.Flags().String("unhealthy_threshold", "", "")

	targetgroupCmd.AddCommand(getTargetgroupCmd)
	targetgroupCmd.AddCommand(listTargetgroupCmd)
	targetgroupCmd.AddCommand(deleteTargetgroupCmd)

	// TargetgroupTarget
	targetgroupCmd.AddCommand(targetgroupTargetCmd)
	targetgroupTargetCmd.AddCommand(createTargetgroupTargetCmd)
	createTargetgroupTargetCmd.Flags().String("backend_protocol", "", "")
	createTargetgroupTargetCmd.Flags().String("backend_port", "", "")
	createTargetgroupTargetCmd.Flags().String("ip", "", "")
	createTargetgroupTargetCmd.Flags().String("cloudid", "", "")

	targetgroupTargetCmd.AddCommand(getTargetgroupTargetCmd)
	targetgroupTargetCmd.AddCommand(listTargetgroupTargetCmd)
	targetgroupTargetCmd.AddCommand(deleteTargetgroupTargetCmd)
}
