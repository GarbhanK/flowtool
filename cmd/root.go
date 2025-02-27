package cmd

import (
	"fmt"
	"os"
	"log"

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

			err := viper.ReadInConfig();
			if err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					// config file not found, ignore error if desired
					log.Printf("Error: Config file not found: %s", err)
				} else {
					// config file was found but another error was produced	
					log.Printf("Error reading in config from viper: %s", err)
				}
			}

			verboseFlag, err := cmd.Flags().GetBool("verbose")
			if err != nil {
				log.Printf("Error retrieving verbose flag: %s", err.Error())
			} else if verboseFlag {
				fmt.Printf("Using config file %s:", viper.ConfigFileUsed())
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// cobra settings
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

		// Search config in home dir with name config.json (without extension)
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".flowtool")
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()
}
