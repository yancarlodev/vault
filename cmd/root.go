package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	configFilePath string

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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(
		&configFilePath,
		"config",
		"c",
		"",
		fmt.Sprintf("config file (default is %s/.vlt.yaml)", getConfigPath()),
	)
}

func initConfig() {
	if configFilePath != "" {
		viper.SetConfigFile(configFilePath)
	} else {
		configPath := getConfigPath()

		viper.AddConfigPath(fmt.Sprintf("%s/%s", configPath, "vlt"))
		viper.SetConfigType("yaml")
		viper.SetConfigName(".vlt")
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
