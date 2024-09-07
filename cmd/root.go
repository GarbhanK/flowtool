package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "flowtool",
		Short: "Template airflow vars",
		Long: `Tool for replacing Jinja2 templates for use with Apache Airflow.
				Output is returned to the clipboard as well stdout.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("give me a .sql file or change flowtool config")

			verboseFlag, _ := cmd.Flags().GetBool("verbose")
			err := viper.ReadInConfig()
			if verboseFlag && err == nil {
				fmt.Println("Using config file:", viper.ConfigFileUsed())
			} else if verboseFlag && err != nil {
				fmt.Println("config file not found in $HOME/.flowtool.yaml or ./flowtool.yaml")
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Do not print formatted output to the terminal")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Print additional information to the terminal")
	rootCmd.PersistentFlags().String("env", "dev", "The desired environment to template into config values")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home dir with name .flowtool (without extension)
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".flowtool")
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()
}
