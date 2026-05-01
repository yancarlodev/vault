package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yancarlodev/vault/cmd/add"
	"github.com/yancarlodev/vault/cmd/edit"
	"github.com/yancarlodev/vault/cmd/list"
	"github.com/yancarlodev/vault/cmd/rm"
	"github.com/yancarlodev/vault/cmd/show"
	"github.com/yancarlodev/vault/infra"
	"os"
)

var (
	customConfigFilePath string
	configName           string = "config"
	configType           string = "yaml"
	version              string = "v0.1"
)

var rootCmd = &cobra.Command{
	Use:     "vlt",
	Short:   "Vault stores notes, secrets and passwords securely",
	Long:    "A secure and handy note taker that take care of your secrets for you.",
	Version: version,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(err)
	}
}

func init() {
	cobra.OnInitialize(initCLI)

	description := fmt.Sprintf("config file (default is %s/%s.%s)", infra.Dirs.ConfigHome(), configName, configType)

	rootCmd.Flags().StringVarP(&customConfigFilePath, "config", "c", "", description)

	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Yan Lepri yancarlodc@gmail.com")

	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(rm.RmCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(show.ShowCmd)
	rootCmd.AddCommand(edit.EditCmd)
}

func initCLI() {
	setupFolders()

	setConfigFile()

	setPrivateKey()

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func setupFolders() {
	_ = os.Mkdir(infra.Dirs.ConfigHome(), 0700)
	_ = os.Mkdir(infra.Dirs.DataHome(), 0700)
}

func setConfigFile() {
	if customConfigFilePath != "" {
		viper.SetConfigFile(customConfigFilePath)

		return
	}

	viper.AddConfigPath(infra.Dirs.ConfigHome())
	viper.SetConfigType("yaml")
	viper.SetConfigName(".vlt")
}

func setPrivateKey() {
	key, err := infra.GetPrivateKey()

	if err != nil || string(key) == "" {
		keyBytes, err := infra.GenerateCryptKey()

		if err != nil {
			cobra.CheckErr(err)
		}

		err = infra.SetPrivateKey(string(keyBytes))

		if err != nil {
			cobra.CheckErr(err)
		}
	}
}
