package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/uthoplatforms/utho-cli/helper"
	"github.com/uthoplatforms/utho-go/utho"
)

var objectstorageCmd = &cobra.Command{
	Use:   "objectstorage",
	Short: "Use this command to manage object storages.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createObjectstorageCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a objectstorage.",
	Example: "uthoctl objectstorage create <objectstorage-name>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dcslug, _ := cmd.Flags().GetString("dcslug")
		billing, _ := cmd.Flags().GetString("billing")
		size, _ := cmd.Flags().GetString("size")
		price, _ := cmd.Flags().GetString("price")

		params := utho.CreateBucketParams{
			Name:    args[0],
			Dcslug:  dcslug,
			Billing: billing,
			Size:    size,
			Price:   price,
		}
		objectstorage, err := client.ObjectStorage().CreateBucket(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Bucket Name", "Bucket Id", "Status")
		tbl.AddRow(args[0], objectstorage.ID, objectstorage.Status)
		tbl.Print()
	},
}

var getObjectstorageCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get objectstorage info",
	Example: "uthoctl objectstorage get <location-slug> <bucket-name>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		bucket, err := client.ObjectStorage().ReadBucket(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Name", "Dcslug", "Size", "Status", "ObjectCount", "CurrentSize")
		tbl.AddRow(bucket.Name, bucket.Dcslug, bucket.Size, bucket.Status, bucket.ObjectCount, bucket.CurrentSize)
		tbl.Print()
	},
}

var listObjectstorageCmd = &cobra.Command{
	Use:     "list",
	Short:   "List objectstorage info",
	Example: "uthoctl objectstorage list <location-slug>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		buckets, err := client.ObjectStorage().ListBuckets(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Name", "Dcslug", "Size", "Status", "ObjectCount", "CurrentSize")
		for _, bucket := range buckets {
			tbl.AddRow(bucket.Name, bucket.Dcslug, bucket.Size, bucket.Status, bucket.ObjectCount, bucket.CurrentSize)
		}
		tbl.Print()
	},
}

var deleteObjectstorageCmd = &cobra.Command{
	Use:     "delete",
	Short:   "delete a objectstorage from your account.",
	Example: "uthoctl objectstorage delete <location-slug> <bucket-name>",
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

		objectstorage, err := client.ObjectStorage().DeleteBucket(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status: " + objectstorage.Status)
	},
}

// Accesskey
var accesskeyCmd = &cobra.Command{
	Use:   "accesskey",
	Short: "Use this command to manage objectstorage Accesskey.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createAccesskeyCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an objectstorage accesskey.",
	Example: "uthoctl objectstorage accesskey create <location-slug> <accesskey-name>",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		params := utho.CreateAccessKeyParams{
			Dcslug:        args[0],
			AccesskeyName: args[1],
		}
		accesskey, err := client.ObjectStorage().CreateAccessKey(params)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Accesskey Name", "Status")
		tbl.AddRow(args[1], accesskey.Status)
		tbl.Print()
	},
}

var getAccesskeyCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get objectstorage accesskey info",
	Example: "uthoctl objectstorage accesskey get <location-slug> <accesskey-id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		accesskey, err := client.ObjectStorage().ReadAccessKey(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Name", "Accesskey", "Dcslug", "Status", "CreatedAt")
		tbl.AddRow(accesskey.Name, accesskey.Accesskey, accesskey.Dcslug, accesskey.Status, accesskey.CreatedAt)
		tbl.Print()
	},
}

var listAccesskeyCmd = &cobra.Command{
	Use:     "list",
	Short:   "List objectstorage policy",
	Example: "uthoctl objectstorage policy list <objectstorage-id>",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := helper.NewUthoClient()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		accesskeys, err := client.ObjectStorage().ListAccessKeys(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tbl := table.New("Name", "Accesskey", "Dcslug", "Status", "CreatedAt")
		for _, accesskey := range accesskeys {
			tbl.AddRow(accesskey.Name, accesskey.Accesskey, accesskey.Dcslug, accesskey.Status, accesskey.CreatedAt)
		}
		tbl.Print()
	},
}

// var deleteAccesskeyCmd = &cobra.Command{
// 	Use:     "delete",
// 	Short:   "delete an objectstorage Accesskey from your account.",
// 	Example: "uthoctl objectstorage accesskey delete <objectstorage-id> <accesskey-id>",
// 	Args:    cobra.ExactArgs(2),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		confirm := helper.Ask()
// 		if !confirm {
// 			fmt.Println("Operation aborted.")
// 			os.Exit(1)
// 		}

// 		client, err := helper.NewUthoClient()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		accesskey, err := client.ObjectStorage().DeleteAccessKey(args[0], args[1])
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		fmt.Println("Status: " + accesskey.Status)
// 	},
// }

func init() {
	rootCmd.AddCommand(objectstorageCmd)
	// Objectstorage
	objectstorageCmd.AddCommand(createObjectstorageCmd)
	createObjectstorageCmd.Flags().String("dcslug", "", "Provide Zone dcslug eg: innoida")
	createObjectstorageCmd.Flags().String("Billing", "", "")
	createObjectstorageCmd.Flags().String("Size", "", "")
	createObjectstorageCmd.Flags().String("Price", "", "")

	objectstorageCmd.AddCommand(getObjectstorageCmd)
	objectstorageCmd.AddCommand(listObjectstorageCmd)
	objectstorageCmd.AddCommand(deleteObjectstorageCmd)

	// Accesskey
	objectstorageCmd.AddCommand(accesskeyCmd)
	accesskeyCmd.AddCommand(createAccesskeyCmd)
	accesskeyCmd.AddCommand(getAccesskeyCmd)
	accesskeyCmd.AddCommand(listAccesskeyCmd)
	// accesskeyCmd.AddCommand(deleteAccesskeyCmd)
}
