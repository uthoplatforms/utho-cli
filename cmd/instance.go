package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/uthoplatforms/utho-cli/helper"
	"github.com/uthoplatforms/utho-go/utho"
)

var instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Use this command to manage compute instances.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createCloudInstanceCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an compute instance.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dcslug, _ := cmd.Flags().GetString("dcslug")
		image, _ := cmd.Flags().GetString("image")
		planid, _ := cmd.Flags().GetString("planid")
		auth, _ := cmd.Flags().GetString("auth")
		rootPassword, _ := cmd.Flags().GetString("root_password")
		firewall, _ := cmd.Flags().GetString("firewall")
		enablebackup, _ := cmd.Flags().GetString("enablebackup")
		support, _ := cmd.Flags().GetString("support")
		management, _ := cmd.Flags().GetString("management")
		billingcycle, _ := cmd.Flags().GetString("billingcycle")
		backupid, _ := cmd.Flags().GetString("backupid")
		snapshotid, _ := cmd.Flags().GetString("snapshotid")
		sshkeys, _ := cmd.Flags().GetString("sshkeys")
		params := utho.CreateCloudInstanceParams{
			Dcslug:       dcslug,
			Image:        image,
			Planid:       planid,
			Auth:         auth,
			RootPassword: rootPassword,
			Firewall:     firewall,
			Enablebackup: enablebackup,
			Support:      support,
			Management:   management,
			Billingcycle: billingcycle,
			Backupid:     backupid,
			Snapshotid:   snapshotid,
			Sshkeys:      sshkeys,
			Cloud:        []utho.CloudHostname{{Hostname: args[0]}},
		}
		instance, err := client.CloudInstances().Create(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Instance Name", "ID", "Password", "Ipv4", "Status")
		tbl.AddRow(args[0], instance.ID, instance.Password, instance.Ipv4, instance.Status)
		tbl.Print()
	},
}

var getCloudInstanceCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get instance info",
	Example: "uthoctl instance get <instance-id>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		instance, err := client.CloudInstances().Read(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Hostname", "CPU", "RAM", "Disksize", "IP", "Billingcycle", "Image")
		tbl.AddRow(instance.ID, instance.Hostname, instance.CPU, instance.RAM, instance.Disksize, instance.IP, instance.Billingcycle, instance.Image.Image)
		tbl.Print()
	},
}

var listCloudInstanceCmd = &cobra.Command{
	Use:   "list",
	Short: "List instance info",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		instances, err := client.CloudInstances().List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("ID", "Hostname", "CPU", "RAM", "Disksize", "IP", "Billingcycle", "Image")
		for _, instance := range instances {
			tbl.AddRow(instance.ID, instance.Hostname, instance.CPU, instance.RAM, instance.Disksize, instance.IP, instance.Billingcycle, instance.Image.Image)
		}
		tbl.Print()
	},
}

var deleteCloudInstanceCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a instance from your account.",
	Example: "uthoctl instance delete <instance-id>",
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

		instance, err := client.CloudInstances().Delete(args[0],
			utho.DeleteCloudInstanceParams{Confirm: "I am aware this action will delete data and server permanently"},
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Status")
		tbl.AddRow(instance.Status)
		tbl.Print()
	},
}

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Use this command to to manage snapshot for your instances.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createSnapshotCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a snapshot for compute instance.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		instance, err := client.CloudInstances().CreateSnapshot(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + instance.Status)
	},
}

var deleteSnapshotCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete an instance snapshot.",
	Example: "uthoctl instance snapshot delete <instance-id> <snapshot-id>",
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

		instance, err := client.CloudInstances().DeleteSnapshot(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Status")
		tbl.AddRow(instance.Status)
		tbl.Print()
	},
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Use this command to to manage backup for your instances.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var enableBackupCmd = &cobra.Command{
	Use:     "enable",
	Short:   "enable backup for compute instance.",
	Example: "uthoctl instance backup enable <instance-id>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		instance, err := client.CloudInstances().EnableBackup(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + instance.Status)
	},
}

var disableBackupCmd = &cobra.Command{
	Use:     "disable",
	Short:   "disable an instance backup.",
	Example: "uthoctl instance backup disable <instance-id>",
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

		instance, err := client.CloudInstances().DisableBackup(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Status")
		tbl.AddRow(instance.Status)
		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(instanceCmd)

	instanceCmd.AddCommand(createCloudInstanceCmd)
	createCloudInstanceCmd.Flags().String("dcslug", "", "Provide Zone dcslug eg: innoida")
	createCloudInstanceCmd.Flags().String("image", "", "Image name eg: centos-7.4-x86_64")
	createCloudInstanceCmd.Flags().String("planid", "", "Cloud Plan ID")
	createCloudInstanceCmd.Flags().String("auth", "", "")
	createCloudInstanceCmd.Flags().String("root_password", "", "")
	createCloudInstanceCmd.Flags().String("firewall", "", "")
	createCloudInstanceCmd.Flags().String("enablebackup", "", "Please pass value 'on' to enable weekly backups")
	createCloudInstanceCmd.Flags().String("support", "", "")
	createCloudInstanceCmd.Flags().String("management", "", "")
	createCloudInstanceCmd.Flags().String("billingcycle", "", "If you required billing cycle other then hourly billing you can pass value as eg: monthly, 3month, 6month, 12month. by default its selected as hourly")
	createCloudInstanceCmd.Flags().String("backupid", "", "Provide a backupid if you have a backup in same datacenter location")
	createCloudInstanceCmd.Flags().String("snapshotid", "", "Provide a snapshot id if you have a snapshot in same datacenter location")
	createCloudInstanceCmd.Flags().String("sshkeys", "", "Privide SSH Key ids or pass multiple SSH Key ids with commans (eg: 432,331)")

	instanceCmd.AddCommand(getCloudInstanceCmd)
	instanceCmd.AddCommand(listCloudInstanceCmd)
	instanceCmd.AddCommand(deleteCloudInstanceCmd)

	// Snapshot
	instanceCmd.AddCommand(snapshotCmd)
	snapshotCmd.AddCommand(createSnapshotCmd)
	snapshotCmd.AddCommand(deleteSnapshotCmd)

	// Backup
	instanceCmd.AddCommand(backupCmd)
	backupCmd.AddCommand(enableBackupCmd)
	backupCmd.AddCommand(disableBackupCmd)

}
