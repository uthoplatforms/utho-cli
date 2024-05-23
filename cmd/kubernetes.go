package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/uthoplatforms/utho-cli/helper"
	"github.com/uthoplatforms/utho-go/utho"
)

var kubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "Use this command to manage kubernetes cluster.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createKubernetesCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a kubernetes.",
	Example: "uthoctl kubernetes create <kubernetes-name>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dcslug, _ := cmd.Flags().GetString("dcslug")
		clusterLabel, _ := cmd.Flags().GetString("cluster_label")
		clusterVersion, _ := cmd.Flags().GetString("cluster_version")
		auth, _ := cmd.Flags().GetString("auth")
		vpc, _ := cmd.Flags().GetString("vpc")
		securityGroups, _ := cmd.Flags().GetString("security_groups")

		params := utho.CreateKubernetesParams{
			Dcslug:         dcslug,
			ClusterLabel:   clusterLabel,
			ClusterVersion: clusterVersion,
			Nodepools:      []utho.CreateNodepoolsParams{},
			Auth:           auth,
			Vpc:            vpc,
			SecurityGroups: securityGroups,
		}
		kubernetes, err := client.Kubernetes().Create(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Kubernetes Name", "Kubernetes Id", "Status")
		tbl.AddRow(args[0], kubernetes.ID, kubernetes.Status)
		tbl.Print()
	},
}

var getKubernetesCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get kubernetes info",
	Example: "uthoctl kubernetes get <kubernetes-id>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		kubernetes, err := client.Kubernetes().Read(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Dcslug", "Name", "RAM", "CPU", "Disksize", "IP", "Status", "WorkerCount")
		tbl.AddRow(kubernetes.ID, kubernetes.Dcslug, kubernetes.Hostname, kubernetes.RAM, kubernetes.CPU, kubernetes.Disksize, kubernetes.IP, kubernetes.Status, kubernetes.WorkerCount)
		tbl.Print()
	},
}

var listKubernetesCmd = &cobra.Command{
	Use:     "list",
	Short:   "List kubernetes info",
	Example: "uthoctl kubernetes list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		kubernetess, err := client.Kubernetes().List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Dcslug", "Name", "RAM", "CPU", "Disksize", "IP", "Status", "WorkerCount")
		for _, kubernetes := range kubernetess {
			tbl.AddRow(kubernetes.ID, kubernetes.Dcslug, kubernetes.Hostname, kubernetes.RAM, kubernetes.CPU, kubernetes.Disksize, kubernetes.IP, kubernetes.Status, kubernetes.WorkerCount)
		}
		tbl.Print()
	},
}

var deleteKubernetesCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a kubernetes from your account.",
	Example: "uthoctl kubernetes delete <kubernetes-id>",
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

		kubernetes, err := client.Kubernetes().Delete(utho.DeleteKubernetesParams{
			ClusterId: args[0],
			Confirm:   "I am aware this action will delete data and cluster permanently",
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + kubernetes.Status)
	},
}

// Loadbalancer
var kubernetesLoadbalancerCmd = &cobra.Command{
	Use:   "loadbalancer",
	Short: "Use this command to manage kubernetes Load Balancer.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createKubernetesLoadbalancerCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a kubernetes loadbalancer.",
	Example: "uthoctl kubernetes loadbalancer create <kubernetes-id> <loadbalancer-id>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		params := utho.CreateKubernetesLoadbalancerParams{
			KubernetesId:   args[0],
			LoadbalancerId: args[1],
		}
		loadbalancer, err := client.Kubernetes().CreateLoadbalancer(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Kubernetes Loadbalancer Id", "Status")
		tbl.AddRow(loadbalancer.ID, loadbalancer.Status)
		tbl.Print()
	},
}

var getKubernetesLoadbalancerCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get kubernetes loadbalancer info",
	Example: "uthoctl kubernetes loadbalancer get <kubernetes-id> <loadbalancer-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		loadbalancer, err := client.Kubernetes().ReadLoadbalancer(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "IP")
		tbl.AddRow(loadbalancer.ID, loadbalancer.Name, loadbalancer.IP)
		tbl.Print()
	},
}

var listKubernetesLoadbalancerCmd = &cobra.Command{
	Use:     "list",
	Short:   "List kubernetes Loadbalancer",
	Example: "uthoctl kubernetes loadbalancer list <kubernetes-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		loadbalancers, err := client.Kubernetes().ListLoadbalancers(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "IP")
		for _, loadbalancer := range loadbalancers {
			tbl.AddRow(loadbalancer.ID, loadbalancer.Name, loadbalancer.IP)
		}
		tbl.Print()
	},
}

var deleteKubernetesLoadbalancerCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a kubernetes policy from your account.",
	Example: "uthoctl kubernetes loadbalancer delete <kubernetes-id> <loadbalancer-id>",
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

		loadbalancer, err := client.Kubernetes().DeleteLoadbalancer(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + loadbalancer.Status)
	},
}

// Securitygroup
var kubernetesecuritygroupCmd = &cobra.Command{
	Use:   "securitygroup",
	Short: "Use this command to manage kubernetes Security Group.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createKubernetesSecuritygroupCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a kubernetes securitygroup.",
	Example: "uthoctl kubernetes securitygroup create <kubernetes-id> <securitygroup-id>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		params := utho.CreateKubernetesSecurityGroupParams{
			KubernetesId:              args[0],
			KubernetesSecurityGroupId: args[1],
		}
		securitygroup, err := client.Kubernetes().CreateSecurityGroup(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Securitygroup id", "Status")
		tbl.AddRow(args[1], securitygroup.Status)
		tbl.Print()
	},
}

var getKubernetesSecuritygroupCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get kubernetes securitygroup info",
	Example: "uthoctl kubernetes securitygroup get <kubernetes-id> <securitygroup-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		securitygroup, err := client.Kubernetes().ReadSecurityGroup(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name")
		tbl.AddRow(securitygroup.ID, securitygroup.Name)
		tbl.Print()
	},
}

var listKubernetesSecuritygroupCmd = &cobra.Command{
	Use:     "list",
	Short:   "List kubernetes policy",
	Example: "uthoctl kubernetes securitygroup list <kubernetes-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		securitygroups, err := client.Kubernetes().ListSecurityGroups(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name")
		for _, securitygroup := range securitygroups {
			tbl.AddRow(securitygroup.ID, securitygroup.Name)
		}
		tbl.Print()
	},
}

var deleteKubernetesSecuritygroupCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a kubernetes Securitygroup from your account.",
	Example: "uthoctl kubernetes securitygroup delete <kubernetes-id> <securitygroup-id>",
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

		securitygroup, err := client.Kubernetes().DeleteSecurityGroup(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + securitygroup.Status)
	},
}

// Targetgroup
var kubernetesTargetgroupCmd = &cobra.Command{
	Use:   "targetgroup",
	Short: "Use this command to manage kubernetes Target Group.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createKubernetesTargetgroupCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a kubernetes targetgroup.",
	Example: "uthoctl kubernetes targetgroup create <kubernetes-id> <targetgroup-id>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		params := utho.CreateKubernetesTargetgroupParams{
			KubernetesId:            args[0],
			KubernetesTargetgroupId: args[1],
		}
		targetgroup, err := client.Kubernetes().CreateTargetgroup(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Targetgroup Name", "Targetgroup Id", "Status")
		tbl.AddRow(args[1], targetgroup.ID, targetgroup.Status)
		tbl.Print()
	},
}

var getKubernetesTargetgroupCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get kubernetes targetgroup info",
	Example: "uthoctl kubernetes targetgroup get <kubernetes-id> <targetgroup-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		targetgroup, err := client.Kubernetes().ReadTargetgroup(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "Protocol", "Port")
		tbl.AddRow(targetgroup.ID, targetgroup.Name, targetgroup.Protocol, targetgroup.Port)
		tbl.Print()
	},
}

var listKubernetesTargetgroupCmd = &cobra.Command{
	Use:     "list",
	Short:   "List kubernetes policy",
	Example: "uthoctl kubernetes policy list <kubernetes-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		targetgroups, err := client.Kubernetes().ListTargetgroups(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "Protocol", "Port")
		for _, targetgroup := range targetgroups {
			tbl.AddRow(targetgroup.ID, targetgroup.Name, targetgroup.Protocol, targetgroup.Port)
		}
		tbl.Print()
	},
}

var deleteKubernetesTargetgroupCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a kubernetes Targetgroup from your account.",
	Example: "uthoctl kubernetes targetgroup delete <kubernetes-id> <targetgroup-id>",
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

		targetgroup, err := client.Kubernetes().DeleteTargetgroup(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + targetgroup.Status)
	},
}

func init() {
	rootCmd.AddCommand(kubernetesCmd)
	// Kubernetes
	kubernetesCmd.AddCommand(createKubernetesCmd)
	createKubernetesCmd.Flags().String("dcslug", "", "")
	createKubernetesCmd.Flags().String("cluster_label", "", "")
	createKubernetesCmd.Flags().String("cluster_version", "", "")
	createKubernetesCmd.Flags().String("auth", "", "")
	createKubernetesCmd.Flags().String("vpc", "", "")
	createKubernetesCmd.Flags().String("security_groups", "", "")

	kubernetesCmd.AddCommand(getKubernetesCmd)
	kubernetesCmd.AddCommand(listKubernetesCmd)
	kubernetesCmd.AddCommand(deleteKubernetesCmd)

	// Loadbalancer
	kubernetesCmd.AddCommand(kubernetesLoadbalancerCmd)
	kubernetesLoadbalancerCmd.AddCommand(createKubernetesLoadbalancerCmd)
	kubernetesLoadbalancerCmd.AddCommand(getKubernetesLoadbalancerCmd)
	kubernetesLoadbalancerCmd.AddCommand(listKubernetesLoadbalancerCmd)
	kubernetesLoadbalancerCmd.AddCommand(deleteKubernetesLoadbalancerCmd)

	// Securitygroup
	kubernetesCmd.AddCommand(kubernetesecuritygroupCmd)
	kubernetesecuritygroupCmd.AddCommand(createKubernetesSecuritygroupCmd)
	kubernetesecuritygroupCmd.AddCommand(getKubernetesSecuritygroupCmd)
	kubernetesecuritygroupCmd.AddCommand(listKubernetesSecuritygroupCmd)
	kubernetesecuritygroupCmd.AddCommand(deleteKubernetesSecuritygroupCmd)

	// Targetgroup
	kubernetesCmd.AddCommand(kubernetesTargetgroupCmd)
	kubernetesTargetgroupCmd.AddCommand(createKubernetesTargetgroupCmd)
	kubernetesTargetgroupCmd.AddCommand(getKubernetesTargetgroupCmd)
	kubernetesTargetgroupCmd.AddCommand(listKubernetesTargetgroupCmd)
	kubernetesTargetgroupCmd.AddCommand(deleteKubernetesTargetgroupCmd)
}
