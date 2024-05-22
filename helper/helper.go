package helper

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"github.com/uthoplatforms/utho-go/utho"
)

func SaveToken(token string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home directory:", err)
		os.Exit(1)
	}

	configDir := filepath.Join(home, ".config")
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		fmt.Println("Error creating config directory:", err)
		os.Exit(1)
	}

	viper.SetConfigType("yaml")
	viper.Set("token", token)

	configFile := filepath.Join(configDir, "uthoctl.yaml")
	if err := viper.WriteConfigAs(configFile); err != nil {
		fmt.Println("Error writing config file:", err)
		os.Exit(1)
	}

	fmt.Println("Token saved successfully at", configFile)
}

func NewUthoClient() (utho.Client, error) {
	token := viper.GetString("token")
	if token == "" {
		return nil, errors.New("no token found. please login first")
	}

	clinet, err := utho.NewClient(token)
	if err != nil {
		return nil, err
	}
	return clinet, err
}

func Ask() bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Are you sure you want to proceed? (y/n): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	input = strings.TrimSpace(input)
	if strings.ToLower(input) == "y" {
		return true
	} else {
		return false
	}
}
