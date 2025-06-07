package cmd

import (
	"fmt"
	"github.com/apparentlymart/go-userdirs/userdirs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Dirs = userdirs.ForApp("Vault", "Lepri Developer", "com.yancarlodev.vlt")
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
	cobra.OnInitialize(initConfig)

	description := fmt.Sprintf("config file (default is %s/%s.%s)", Dirs.ConfigHome(), configName, configType)

	rootCmd.Flags().StringVarP(&customConfigFilePath, "config", "c", "", description)

	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Yan Lepri yancarlodc@gmail.com")

	rootCmd.AddCommand(addCmd)
}

func initConfig() {
	setConfigFile()

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func setConfigFile() {
	if customConfigFilePath != "" {
		viper.SetConfigFile(customConfigFilePath)

		return
	}

	viper.AddConfigPath(Dirs.ConfigHome())
	viper.SetConfigType("yaml")
	viper.SetConfigName(".vlt")
}
