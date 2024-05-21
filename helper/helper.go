package helper

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
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
