package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/uthoplatforms/utho-cli/helper"
	"github.com/uthoplatforms/utho-go/utho"
)

var loadbalancerCmd = &cobra.Command{
	Use:   "loadbalancer",
	Short: "Use this command to manage loadbalancers.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createLoadbalancerCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a loadbalancer.",
	Example: "uthoctl loadbalancer create <loadbalancer-name>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dcslug, _ := cmd.Flags().GetString("dcslug")
		loadbalancerType, _ := cmd.Flags().GetString("type")

		params := utho.CreateLoadbalancerParams{
			Name:   args[0],
			Dcslug: dcslug,
			Type:   loadbalancerType,
		}
		loadbalancer, err := client.Loadbalancers().Create(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Loadbalancer Name", "Loadbalancer Id", "Status")
		tbl.AddRow(args[0], loadbalancer.ID, loadbalancer.Status)
		tbl.Print()
	},
}

var getLoadbalancerCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get loadbalancer info",
	Example: "uthoctl loadbalancer get <loadbalancer-id>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		loadbalancer, err := client.Loadbalancers().Read(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "IP", "Algorithm", "Type", "Status")
		tbl.AddRow(loadbalancer.ID, loadbalancer.IP, loadbalancer.Algorithm, loadbalancer.Type, loadbalancer.Status)
		tbl.Print()
	},
}

var listLoadbalancerCmd = &cobra.Command{
	Use:     "list",
	Short:   "List loadbalancer info",
	Example: "uthoctl loadbalancer list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		loadbalancers, err := client.Loadbalancers().List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "IP", "Algorithm", "Type", "Status")
		for _, loadbalancer := range loadbalancers {
			tbl.AddRow(loadbalancer.ID, loadbalancer.IP, loadbalancer.Algorithm, loadbalancer.Type, loadbalancer.Status)
		}
		tbl.Print()
	},
}

var deleteLoadbalancerCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a loadbalancer from your account.",
	Example: "uthoctl loadbalancer delete <loadbalancer-id>",
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

		loadbalancer, err := client.Loadbalancers().Delete(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + loadbalancer.Status)
	},
}

// LoadbalancerAcl
var loadbalancerAclCmd = &cobra.Command{
	Use:   "acl",
	Short: "Use this command to manage Loadbalancer ACL.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()

	},
}

var createLoadbalancerAclCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a Loadbalancer acl.",
	Example: "uthoctl loadbalancer acl create <loadbalancer-id> <acl-name>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		conditionType, _ := cmd.Flags().GetString("condition_type")
		frontendId, _ := cmd.Flags().GetString("frontend_id")
		value, _ := cmd.Flags().GetString("value")
		params := utho.CreateLoadbalancerACLParams{
			LoadbalancerId: args[0],
			Name:           args[1],
			ConditionType:  conditionType,
			FrontendID:     frontendId,
			Value:          value,
		}
		acl, err := client.Loadbalancers().CreateACL(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Acl Name", "Acl Id", "Status")
		tbl.AddRow(args[1], acl.ID, acl.Status)
		tbl.Print()
	},
}

var getLoadbalancerAclCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get Loadbalancer acl info",
	Example: "uthoctl loadbalancer acl get <loadbalancer-id> <acl-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		acl, err := client.Loadbalancers().ReadACL(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "ACLCondition", "Value")
		tbl.AddRow(acl.ID, acl.Name, acl.ACLCondition, acl.Value)
		tbl.Print()
	},
}

var listLoadbalancerAclCmd = &cobra.Command{
	Use:     "list",
	Short:   "List Loadbalancer acl",
	Example: "uthoctl loadbalancer acl list <loadbalancer-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		acls, err := client.Loadbalancers().ListACLs(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "ACLCondition", "Value")
		for _, acl := range acls {
			tbl.AddRow(acl.ID, acl.Name, acl.ACLCondition, acl.Value)
		}
		tbl.Print()
	},
}

var deleteLoadbalancerAclCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a Loadbalancer ACL from your account.",
	Example: "uthoctl loadbalancer acl delete <loadbalancer-id> <acl-id>",
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

		acl, err := client.Loadbalancers().DeleteACL(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + acl.Status)
	},
}

// Loadbalancer Frontend
var loadbalancerFrontendCmd = &cobra.Command{
	Use:   "frontend",
	Short: "Use this command to manage Loadbalancer Frontend.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()

	},
}

var createLoadbalancerFrontendCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a Loadbalancer frontend.",
	Example: "uthoctl loadbalancer frontend create <loadbalancer-id> <frontend-name>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		proto, _ := cmd.Flags().GetString("proto")
		port, _ := cmd.Flags().GetString("port")
		certificateId, _ := cmd.Flags().GetString("certificate_id")
		algorithm, _ := cmd.Flags().GetString("algorithm")
		redirecthttps, _ := cmd.Flags().GetString("redirecthttps")
		cookie, _ := cmd.Flags().GetString("cookie")
		params := utho.CreateLoadbalancerFrontendParams{
			LoadbalancerId: args[0],
			Name:           args[1],
			Proto:          proto,
			Port:           port,
			CertificateID:  certificateId,
			Algorithm:      algorithm,
			Redirecthttps:  redirecthttps,
			Cookie:         cookie,
		}
		frontend, err := client.Loadbalancers().CreateFrontend(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Frontend Name", "Frontend Id", "Status")
		tbl.AddRow(args[1], frontend.ID, frontend.Status)
		tbl.Print()
	},
}

var getLoadbalancerFrontendCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get Loadbalancer frontend info",
	Example: "uthoctl loadbalancer frontend get <loadbalancer-id> <frontend-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		frontend, err := client.Loadbalancers().ReadFrontend(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "Algorithm", "CertificateID", "Port")
		tbl.AddRow(frontend.ID, frontend.Name, frontend.Algorithm, frontend.CertificateID, frontend.Port)
		tbl.Print()
	},
}

var listLoadbalancerFrontendCmd = &cobra.Command{
	Use:     "list",
	Short:   "List Loadbalancer frontend",
	Example: "uthoctl loadbalancer frontend list <loadbalancer-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		frontends, err := client.Loadbalancers().ListFrontends(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "Algorithm", "CertificateID", "Port")
		for _, frontend := range frontends {
			tbl.AddRow(frontend.ID, frontend.Name, frontend.Algorithm, frontend.CertificateID, frontend.Port)
		}
		tbl.Print()
	},
}

var deleteLoadbalancerFrontendCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a Loadbalancer Frontend from your account.",
	Example: "uthoctl loadbalancer frontend delete <loadbalancer-id> <frontend-id>",
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

		frontend, err := client.Loadbalancers().DeleteFrontend(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + frontend.Status)
	},
}

// Loadbalancer Backend
var loadbalancerBackendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Use this command to manage Loadbalancer Backend.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createLoadbalancerBackendCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a Loadbalancer backend.",
	Example: "uthoctl loadbalancer backend create <loadbalancer-id> <frontend-id> <cloud-id>",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		port, _ := cmd.Flags().GetString("port")
		params := utho.CreateLoadbalancerBackendParams{
			LoadbalancerId: args[0],
			FrontendID:     args[1],
			BackendPort:    port,
			Cloudid:        args[2],
		}
		backend, err := client.Loadbalancers().CreateBackend(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Backend Name", "Backend Id", "Status")
		tbl.AddRow(args[1], backend.ID, backend.Status)
		tbl.Print()
	},
}

var getLoadbalancerBackendCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get Loadbalancer backend info",
	Example: "uthoctl loadbalancer backend get <loadbalancer-id> <backend-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		backend, err := client.Loadbalancers().ReadBackend(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "IP", "Cloudid", "Name", "RAM", "CPU", "Disk")
		tbl.AddRow(backend.ID, backend.IP, backend.Cloudid, backend.Name, backend.RAM, backend.CPU, backend.Disk)
		tbl.Print()
	},
}

var listLoadbalancerBackendCmd = &cobra.Command{
	Use:     "list",
	Short:   "List Loadbalancer backend",
	Example: "uthoctl loadbalancer backend list <loadbalancer-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		backends, err := client.Loadbalancers().ListBackends(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "IP", "Cloudid", "Name", "RAM", "CPU", "Disk")
		for _, backend := range backends {
			tbl.AddRow(backend.ID, backend.IP, backend.Cloudid, backend.Name, backend.RAM, backend.CPU, backend.Disk)
		}
		tbl.Print()
	},
}

var deleteLoadbalancerBackendCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a Loadbalancer Backend from your account.",
	Example: "uthoctl loadbalancer backend delete <loadbalancer-id> <backend-id>",
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

		backend, err := client.Loadbalancers().DeleteBackend(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + backend.Status)
	},
}

// Loadbalancer Route
var loadbalancerRouteCmd = &cobra.Command{
	Use:   "route",
	Short: "Use this command to manage Loadbalancer Route.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createLoadbalancerRouteCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a Loadbalancer route.",
	Example: "uthoctl loadbalancer route create <loadbalancer-id> <frontend-id> <acl-id>",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		routeCondition, _ := cmd.Flags().GetString("route_condition")
		targetGroups, _ := cmd.Flags().GetString("target_groups")
		params := utho.CreateLoadbalancerRouteParams{
			LoadbalancerId: args[0],
			FrontendID:     args[1],
			ACLID:          args[2],
			RouteCondition: routeCondition,
			TargetGroups:   targetGroups,
		}
		route, err := client.Loadbalancers().CreateRoute(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Route Id", "Status")
		tbl.AddRow(route.ID, route.Status)
		tbl.Print()
	},
}

var getLoadbalancerRouteCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get Loadbalancer route info",
	Example: "uthoctl loadbalancer route get <loadbalancer-id> <route-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		route, err := client.Loadbalancers().ReadRoute(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "ACLID", "ACLName", "RoutingCondition", "BackendID")
		tbl.AddRow(route.ID, route.ACLID, route.ACLName, route.RoutingCondition, route.BackendID)
		tbl.Print()
	},
}

var listLoadbalancerRouteCmd = &cobra.Command{
	Use:     "list",
	Short:   "List Loadbalancer route",
	Example: "uthoctl loadbalancer route list <loadbalancer-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		routes, err := client.Loadbalancers().ListRoutes(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "ACLID", "ACLName", "RoutingCondition", "BackendID")
		for _, route := range routes {
			tbl.AddRow(route.ID, route.ACLID, route.ACLName, route.RoutingCondition, route.BackendID)
		}
		tbl.Print()
	},
}

var deleteLoadbalancerRouteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a Loadbalancer Route from your account.",
	Example: "uthoctl loadbalancer route delete <loadbalancer-id> <route-id>",
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

		route, err := client.Loadbalancers().DeleteRoute(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + route.Status)
	},
}

func init() {
	rootCmd.AddCommand(loadbalancerCmd)
	// Loadbalancer
	loadbalancerCmd.AddCommand(createLoadbalancerCmd)
	createLoadbalancerCmd.Flags().String("dcslug", "", "Provide Zone dcslug eg: innoida")
	createLoadbalancerCmd.Flags().String("type", "", "Load-Balancer type must be either application or network. The default value is application")

	loadbalancerCmd.AddCommand(getLoadbalancerCmd)
	loadbalancerCmd.AddCommand(listLoadbalancerCmd)
	loadbalancerCmd.AddCommand(deleteLoadbalancerCmd)

	// acl
	loadbalancerCmd.AddCommand(loadbalancerAclCmd)
	loadbalancerAclCmd.AddCommand(createLoadbalancerAclCmd)
	createLoadbalancerAclCmd.Flags().String("condition_type", "", "")
	createLoadbalancerAclCmd.Flags().String("frontend_id", "", "")
	createLoadbalancerAclCmd.Flags().String("value", "", "")

	loadbalancerAclCmd.AddCommand(getLoadbalancerAclCmd)
	loadbalancerAclCmd.AddCommand(listLoadbalancerAclCmd)
	loadbalancerAclCmd.AddCommand(deleteLoadbalancerAclCmd)

	// Frontend
	loadbalancerCmd.AddCommand(loadbalancerFrontendCmd)
	loadbalancerFrontendCmd.AddCommand(createLoadbalancerFrontendCmd)
	createLoadbalancerFrontendCmd.Flags().String("proto", "", "")
	createLoadbalancerFrontendCmd.Flags().String("port", "", "")
	createLoadbalancerFrontendCmd.Flags().String("certificate_id", "", "")
	createLoadbalancerFrontendCmd.Flags().String("algorithm", "", "")
	createLoadbalancerFrontendCmd.Flags().String("redirecthttps", "", "")
	createLoadbalancerFrontendCmd.Flags().String("cookie", "", "")

	loadbalancerFrontendCmd.AddCommand(getLoadbalancerFrontendCmd)
	loadbalancerFrontendCmd.AddCommand(listLoadbalancerFrontendCmd)
	loadbalancerFrontendCmd.AddCommand(deleteLoadbalancerFrontendCmd)

	// Backend
	loadbalancerCmd.AddCommand(loadbalancerBackendCmd)
	loadbalancerBackendCmd.AddCommand(createLoadbalancerBackendCmd)
	createLoadbalancerBackendCmd.Flags().String("port", "", "")

	loadbalancerBackendCmd.AddCommand(getLoadbalancerBackendCmd)
	loadbalancerBackendCmd.AddCommand(listLoadbalancerBackendCmd)
	loadbalancerBackendCmd.AddCommand(deleteLoadbalancerBackendCmd)

	// Route
	loadbalancerCmd.AddCommand(loadbalancerRouteCmd)
	loadbalancerRouteCmd.AddCommand(createLoadbalancerRouteCmd)
	createLoadbalancerCmd.Flags().String("route_condition", "", "")
	createLoadbalancerCmd.Flags().String("target_groups", "", "")

	loadbalancerRouteCmd.AddCommand(getLoadbalancerRouteCmd)
	loadbalancerRouteCmd.AddCommand(listLoadbalancerRouteCmd)
	loadbalancerRouteCmd.AddCommand(deleteLoadbalancerRouteCmd)

}
