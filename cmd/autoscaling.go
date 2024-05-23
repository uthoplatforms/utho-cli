package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/uthoplatforms/utho-cli/helper"
	"github.com/uthoplatforms/utho-go/utho"
)

var autoscalingCmd = &cobra.Command{
	Use:   "autoscaling",
	Short: "Use this command to manage autoscalings.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createAutoscalingCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an autoscaling Policy.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		osDiskSize, _ := cmd.Flags().GetInt("os_disk_size")
		dcslug, _ := cmd.Flags().GetString("dcslug")
		minsize, _ := cmd.Flags().GetString("minsize")
		maxsize, _ := cmd.Flags().GetString("maxsize")
		desiredsize, _ := cmd.Flags().GetString("desiredsize")
		planid, _ := cmd.Flags().GetString("planid")
		planname, _ := cmd.Flags().GetString("planname")
		instanceTemplateid, _ := cmd.Flags().GetString("instance_templateid")
		publicIpEnabledStr, _ := cmd.Flags().GetString("public_ip_enabled")
		vpc, _ := cmd.Flags().GetString("vpc")
		loadBalancers, _ := cmd.Flags().GetString("load_balancers")
		securityGroups, _ := cmd.Flags().GetString("security_groups")
		stackid, _ := cmd.Flags().GetString("stackid")
		stackimage, _ := cmd.Flags().GetString("stackimage")
		targetGroups, _ := cmd.Flags().GetString("target_groups")

		publicIpEnabled, err := helper.StringToBool(publicIpEnabledStr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		params := utho.CreateAutoScalingParams{
			Name:               args[0],
			OsDiskSize:         osDiskSize,
			Dcslug:             dcslug,
			Minsize:            minsize,
			Maxsize:            maxsize,
			Desiredsize:        desiredsize,
			Planid:             planid,
			Planname:           planname,
			InstanceTemplateid: instanceTemplateid,
			PublicIPEnabled:    publicIpEnabled,
			Vpc:                vpc,
			LoadBalancers:      loadBalancers,
			SecurityGroups:     securityGroups,
			Policies:           []utho.CreatePoliciesParams{},
			Schedules:          []utho.CreateSchedulesParams{},
			Stackid:            stackid,
			Stackimage:         stackimage,
			TargetGroups:       targetGroups,
		}
		autoscaling, err := client.AutoScaling().Create(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Autoscaling Name", "Autoscaling Id", "Status")
		tbl.AddRow(args[0], autoscaling.ID, autoscaling.Status)
		tbl.Print()
	},
}

var getAutoscalingCmd = &cobra.Command{
	Use:   "get",
	Short: "Get autoscaling info",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		autoscaling, err := client.AutoScaling().Read(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "Dcslug", "Minsize", "Maxsize", "Image", "Status")
		tbl.AddRow(autoscaling.ID, autoscaling.Name, autoscaling.Dcslug, autoscaling.Minsize, autoscaling.Maxsize, autoscaling.Image, autoscaling.Status)
		tbl.Print()
	},
}

var listAutoscalingCmd = &cobra.Command{
	Use:   "list",
	Short: "List autoscaling info",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		autoscalings, err := client.AutoScaling().List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "Dcslug", "Minsize", "Maxsize", "Image", "Status")
		for _, autoscaling := range autoscalings {
			tbl.AddRow(autoscaling.ID, autoscaling.Name, autoscaling.Dcslug, autoscaling.Minsize, autoscaling.Maxsize, autoscaling.Image, autoscaling.Status)
		}
		tbl.Print()
	},
}

var deleteAutoscalingCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete an autoscaling from your account.",
	Example: "uthoctl autoscaling delete <autoscaling-id> <autoscaling-name>",
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

		autoscaling, err := client.AutoScaling().Delete(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + autoscaling.Status)
	},
}

// Policy
var policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Use this command to manage autoscalings policy.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createPolicyCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an autoscaling Policy.",
	Example: "uthoctl autoscaling policy create <autoscaling-id> <policy-name>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		policyType, _ := cmd.Flags().GetString("type")
		compare, _ := cmd.Flags().GetString("compare")
		value, _ := cmd.Flags().GetString("value")
		adjust, _ := cmd.Flags().GetString("adjust")
		period, _ := cmd.Flags().GetString("period")
		cooldown, _ := cmd.Flags().GetString("cooldown")
		product, _ := cmd.Flags().GetString("product")

		params := utho.CreateAutoScalingPolicyParams{
			Name:      args[1],
			Type:      policyType,
			Compare:   compare,
			Value:     value,
			Adjust:    adjust,
			Period:    period,
			Cooldown:  cooldown,
			Product:   product,
			Productid: args[0],
		}
		policy, err := client.AutoScaling().CreatePolicy(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Policy Name", "Autoscaling Id", "Status")
		tbl.AddRow(args[1], policy.ID, policy.Status)
		tbl.Print()
	},
}

var getPolicyCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get autoscaling policy info",
	Example: "uthoctl autoscaling policy get <autoscaling-id> <policy-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		policy, err := client.AutoScaling().ReadPolicy(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Productid", "Name", "Type", "Value", "Status", "Cloudid", "Maxsize", "Minsize")
		tbl.AddRow(policy.ID, policy.Productid, policy.Name, policy.Type, policy.Value, policy.Status, policy.Cloudid, policy.Maxsize, policy.Minsize)
		tbl.Print()
	},
}

var listPolicyCmd = &cobra.Command{
	Use:     "list",
	Short:   "List autoscaling policy",
	Example: "uthoctl autoscaling policy list <autoscaling-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		policies, err := client.AutoScaling().ListPolicies(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Autoscaling id", "Name", "Type", "Value", "Status", "Cloudid", "Maxsize", "Minsize")
		for _, policy := range policies {
			tbl.AddRow(policy.ID, policy.Productid, policy.Name, policy.Type, policy.Value, policy.Status, policy.Cloudid, policy.Maxsize, policy.Minsize)
		}
		tbl.Print()
	},
}

var deletePolicyCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete an autoscaling policy from your account.",
	Example: "uthoctl autoscaling policy delete <policy-id>",
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

		autoscaling, err := client.AutoScaling().DeletePolicy(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + autoscaling.Status)
	},
}

// Schedule Policy
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Use this command to manage autoscalings Schedule Policy.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createScheduleCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an autoscaling schedule schedule.",
	Example: "uthoctl autoscaling schedule create <autoscaling-id> <schedule-name>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		desiredsize, _ := cmd.Flags().GetString("desiredsize")
		recurrence, _ := cmd.Flags().GetString("recurrence")
		startDate, _ := cmd.Flags().GetString("start_date")

		params := utho.CreateAutoScalingScheduleParams{
			AutoScalingId: args[0],
			Name:          args[1],
			Desiredsize:   desiredsize,
			Recurrence:    recurrence,
			StartDate:     startDate,
		}
		schedule, err := client.AutoScaling().CreateSchedule(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Schedule Name", "Schedule Id", "Status")
		tbl.AddRow(args[1], schedule.ID, schedule.Status)
		tbl.Print()
	},
}

var getScheduleCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get autoscaling schedule info",
	Example: "uthoctl autoscaling schedule get <autoscaling-id> <schedule-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		schedule, err := client.AutoScaling().ReadSchedule(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Group id", "Name", "Desiredsize", "Recurrence", "StartDate", "Status", "Timezone")
		tbl.AddRow(schedule.ID, schedule.Groupid, schedule.Name, schedule.Desiredsize, schedule.Recurrence, schedule.StartDate, schedule.Status, schedule.Timezone)
		tbl.Print()
	},
}

var listScheduleCmd = &cobra.Command{
	Use:     "list",
	Short:   "List autoscaling policy",
	Example: "uthoctl autoscaling policy list <autoscaling-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		schedules, err := client.AutoScaling().ListSchedules(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Group id", "Name", "Desiredsize", "Recurrence", "StartDate", "Timezone")
		for _, schedule := range schedules {
			tbl.AddRow(schedule.ID, schedule.Groupid, schedule.Name, schedule.Desiredsize, schedule.Recurrence, schedule.StartDate, schedule.Timezone)
		}
		tbl.Print()
	},
}

var deleteScheduleCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete an autoscaling policy from your account.",
	Example: "uthoctl autoscaling schedule delete <autoscaling-id> <schedule-id>",
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

		autoscaling, err := client.AutoScaling().DeleteSchedule(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + autoscaling.Status)
	},
}

// Loadbalancer
var autoscalingLoadbalancerCmd = &cobra.Command{
	Use:   "loadbalancer",
	Short: "Use this command to manage autoscalings Load Balancer.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createAutoscalingLoadbalancerCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an autoscaling loadbalancer.",
	Example: "uthoctl autoscaling loadbalancer create <autoscaling-id> <loadbalancer-id>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		params := utho.CreateAutoScalingLoadbalancerParams{
			AutoScalingId:  args[0],
			LoadbalancerId: args[1],
		}
		loadbalancer, err := client.AutoScaling().CreateLoadbalancer(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("AutoScaling Loadbalancer Id", "Status")
		tbl.AddRow(loadbalancer.ID, loadbalancer.Status)
		tbl.Print()
	},
}

var getAutoscalingLoadbalancerCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get autoscaling loadbalancer info",
	Example: "uthoctl autoscaling loadbalancer get <autoscaling-id> <loadbalancer-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		loadbalancer, err := client.AutoScaling().ReadLoadbalancer(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "IP")
		tbl.AddRow(loadbalancer.ID, loadbalancer.Name, loadbalancer.IP)
		tbl.Print()
	},
}

var listAutoscalingLoadbalancerCmd = &cobra.Command{
	Use:     "list",
	Short:   "List autoscaling Loadbalancer",
	Example: "uthoctl autoscaling loadbalancer list <autoscaling-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		loadbalancers, err := client.AutoScaling().ListLoadbalancers(args[0])
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

var deleteAutoscalingLoadbalancerCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete an autoscaling policy from your account.",
	Example: "uthoctl autoscaling loadbalancer delete <autoscaling-id> <loadbalancer-id>",
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

		loadbalancer, err := client.AutoScaling().DeleteLoadbalancer(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + loadbalancer.Status)
	},
}

// Securitygroup
var securitygroupCmd = &cobra.Command{
	Use:   "securitygroup",
	Short: "Use this command to manage autoscalings Security Group.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createSecuritygroupCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an autoscaling securitygroup.",
	Example: "uthoctl autoscaling securitygroup create <autoscaling-id> <securitygroup-id>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		params := utho.CreateAutoScalingSecurityGroupParams{
			AutoScalingId:              args[0],
			AutoScalingSecurityGroupId: args[1],
		}
		securitygroup, err := client.AutoScaling().CreateSecurityGroup(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Securitygroup Name", "Securitygroup Id", "Status")
		tbl.AddRow(args[1], securitygroup.ID, securitygroup.Status)
		tbl.Print()
	},
}

var getSecuritygroupCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get autoscaling securitygroup info",
	Example: "uthoctl autoscaling securitygroup get <autoscaling-id> <securitygroup-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		securitygroup, err := client.AutoScaling().ReadSecurityGroup(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name")
		tbl.AddRow(securitygroup.ID, securitygroup.Name)
		tbl.Print()
	},
}

var listSecuritygroupCmd = &cobra.Command{
	Use:     "list",
	Short:   "List autoscaling securitygroup",
	Example: "uthoctl autoscaling securitygroup list <autoscaling-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		securitygroups, err := client.AutoScaling().ListSecurityGroups(args[0])
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

var deleteSecuritygroupCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete an autoscaling سecuritygroup from your account.",
	Example: "uthoctl autoscaling securitygroup delete <autoscaling-id> <securitygroup-id>",
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

		securitygroup, err := client.AutoScaling().DeleteSecurityGroup(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + securitygroup.Status)
	},
}

// Targetgroup
var autoscalingtargetgroupCmd = &cobra.Command{
	Use:   "targetgroup",
	Short: "Use this command to manage autoscalings Target Group.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createAutoscalingTargetgroupCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an autoscaling targetgroup.",
	Example: "uthoctl autoscaling targetgroup create <autoscaling-id> <targetgroup-id>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		params := utho.CreateAutoScalingTargetgroupParams{
			AutoScalingId:            args[0],
			AutoScalingTargetgroupId: args[1],
		}
		targetgroup, err := client.AutoScaling().CreateTargetgroup(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Targetgroup Name", "Targetgroup Id", "Status")
		tbl.AddRow(args[1], targetgroup.ID, targetgroup.Status)
		tbl.Print()
	},
}

var getAutoscalingTargetgroupCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get autoscaling targetgroup info",
	Example: "uthoctl autoscaling targetgroup get <autoscaling-id> <targetgroup-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		targetgroup, err := client.AutoScaling().ReadTargetgroup(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Name", "Protocol", "Port")
		tbl.AddRow(targetgroup.ID, targetgroup.Name, targetgroup.Protocol, targetgroup.Port)
		tbl.Print()
	},
}

var listAutoscalingTargetgroupCmd = &cobra.Command{
	Use:     "list",
	Short:   "List autoscaling policy",
	Example: "uthoctl autoscaling policy list <autoscaling-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		targetgroups, err := client.AutoScaling().ListTargetgroups(args[0])
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

var deleteAutoscalingTargetgroupCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete an autoscaling Targetgroup from your account.",
	Example: "uthoctl autoscaling targetgroup delete <autoscaling-id> <targetgroup-id>",
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

		targetgroup, err := client.AutoScaling().DeleteTargetgroup(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + targetgroup.Status)
	},
}

func init() {
	rootCmd.AddCommand(autoscalingCmd)

	autoscalingCmd.AddCommand(createAutoscalingCmd)
	createAutoscalingCmd.Flags().Int("os_disk_size", 0, "")
	createAutoscalingCmd.Flags().String("dcslug", "", "")
	createAutoscalingCmd.Flags().String("minsize", "", "")
	createAutoscalingCmd.Flags().String("maxsize", "", "")
	createAutoscalingCmd.Flags().String("desiredsize", "", "")
	createAutoscalingCmd.Flags().String("planid", "", "")
	createAutoscalingCmd.Flags().String("planname", "", "")
	createAutoscalingCmd.Flags().String("instance_templateid", "", "")
	createAutoscalingCmd.Flags().String("public_ip_enabled", "", "")
	createAutoscalingCmd.Flags().String("vpc", "", "")
	createAutoscalingCmd.Flags().String("load_balancers", "", "")
	createAutoscalingCmd.Flags().String("security_groups", "", "")
	createAutoscalingCmd.Flags().String("stackid", "", "")
	createAutoscalingCmd.Flags().String("stackimage", "", "")
	createAutoscalingCmd.Flags().String("target_groups", "", "")

	autoscalingCmd.AddCommand(getAutoscalingCmd)
	autoscalingCmd.AddCommand(listAutoscalingCmd)
	autoscalingCmd.AddCommand(deleteAutoscalingCmd)

	// Policy
	autoscalingCmd.AddCommand(policyCmd)
	policyCmd.AddCommand(createPolicyCmd)
	createPolicyCmd.Flags().String("dcslug", "", "Provide Zone dcslug eg: innoida")
	createPolicyCmd.Flags().String("type", "", "")
	createPolicyCmd.Flags().String("compare", "", "")
	createPolicyCmd.Flags().String("value", "", "")
	createPolicyCmd.Flags().String("adjust", "", "")
	createPolicyCmd.Flags().String("period", "", "")
	createPolicyCmd.Flags().String("cooldown", "", "")
	createPolicyCmd.Flags().String("product", "", "")

	policyCmd.AddCommand(getPolicyCmd)
	policyCmd.AddCommand(listPolicyCmd)
	policyCmd.AddCommand(deletePolicyCmd)

	// Schedule
	autoscalingCmd.AddCommand(scheduleCmd)
	scheduleCmd.AddCommand(createScheduleCmd)
	createScheduleCmd.Flags().String("desiredsize", "", "")
	createScheduleCmd.Flags().String("recurrence", "", "")
	createScheduleCmd.Flags().String("start_date", "", "")

	scheduleCmd.AddCommand(getScheduleCmd)
	scheduleCmd.AddCommand(listScheduleCmd)
	scheduleCmd.AddCommand(deleteScheduleCmd)

	// Loadbalancer
	autoscalingCmd.AddCommand(autoscalingLoadbalancerCmd)
	autoscalingLoadbalancerCmd.AddCommand(createAutoscalingLoadbalancerCmd)
	autoscalingLoadbalancerCmd.AddCommand(getAutoscalingLoadbalancerCmd)
	autoscalingLoadbalancerCmd.AddCommand(listAutoscalingLoadbalancerCmd)
	autoscalingLoadbalancerCmd.AddCommand(deleteAutoscalingLoadbalancerCmd)

	// Securitygroup
	autoscalingCmd.AddCommand(securitygroupCmd)
	securitygroupCmd.AddCommand(createSecuritygroupCmd)
	securitygroupCmd.AddCommand(getSecuritygroupCmd)
	securitygroupCmd.AddCommand(listSecuritygroupCmd)
	securitygroupCmd.AddCommand(deleteSecuritygroupCmd)

	// Targetgroup
	autoscalingCmd.AddCommand(autoscalingtargetgroupCmd)
	autoscalingtargetgroupCmd.AddCommand(createAutoscalingTargetgroupCmd)
	autoscalingtargetgroupCmd.AddCommand(getAutoscalingTargetgroupCmd)
	autoscalingtargetgroupCmd.AddCommand(listAutoscalingTargetgroupCmd)
	autoscalingtargetgroupCmd.AddCommand(deleteAutoscalingTargetgroupCmd)
}
