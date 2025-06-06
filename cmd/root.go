package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	customConfigFilePath string
	configPath           string
	configName           string = "config"
	configType           string = "yaml"

	rootCmd = &cobra.Command{
		Use:   "vlt",
		Short: "Vault stores notes, secrets and passwords securely",
		Long:  "A secure and handy note taker that take care of your secrets for you.",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(err)
	}
}

func init() {
	configPath = getConfigPath()

	cobra.OnInitialize(initConfig)

	description := fmt.Sprintf("config file (default is %s/vlt/%s.%s)", configPath, configName, configType)

	rootCmd.PersistentFlags().StringVarP(
		&customConfigFilePath,
		"config",
		"c",
		"",
		description,
	)
}

func initConfig() {
	if customConfigFilePath != "" {
		viper.SetConfigFile(customConfigFilePath)
	} else {
		vltConfigPath := fmt.Sprintf("%s/%s", configPath, "vlt")

		viper.AddConfigPath(vltConfigPath)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func getConfigPath() string {
	configPath, err := os.UserConfigDir()
	cobra.CheckErr(err)

	return configPath
}
