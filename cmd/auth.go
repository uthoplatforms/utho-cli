package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/uthoplatforms/utho-cli/helper"
	"golang.org/x/term"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Please authenticate uthoctl for use with your Utho account. You can generate a token in the control panel at https://console.utho.com/api",
	Run: func(cmd *cobra.Command, args []string) {
		var token string
		for {
			fmt.Fprint(os.Stderr, "Enter your api token: ")
			b, _ := term.ReadPassword(int(syscall.Stdin))
			token = string(b)
			if token != "" {
				break
			}
		}
		helper.SaveToken(token)
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
