package cmd

import (
	"fmt"
	// "os"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var (
	// Used for flags.
	// cfgFile     string

	rootCmd = &cobra.Command{
		Use:   "flowtool",
		Short: "Template airflow vars",
		Long: `Tool for replacing Jinja2 templates for use with Apache Airflow.
				Output is returned to the clipboard as well stdout.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("template a SQL file with 'template' or 'templ'\nconfigure flowtool with 'config'")
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	//     cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Do not print formatted output to the terminal")
	rootCmd.PersistentFlags().String("env", "dev", "The desirev environment to template into config values")
}

//     rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "cinfig file (default")
//     rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
// 	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
// 	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
// 	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
// 	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
// 	viper.SetDefault("license", "apache")

//     rootCmd.AddCommand(addCmd)
//     rootCmd.AddCommand(initCmd)
// }

// func initConfig() {
//     if cfgFile != "" {
//         // Use config file from the flag
//         viper.SetConfigFile(cfgFile)
//     } else {
//         // Find home directory
//         home, err := os.UserHomeDir()
//         cobra.CheckErr(err)

//         // Search config in home dir with name .cobra (without extension)
//         viper.AddConfigPath(home)
//         viper.SetConfigType("yaml")
//         viper.SetConfigName(".cobra")
//     }

//     viper.AutomaticEnv()

//     if err := viper.ReadInConfig(); err == nil {
//         fmt.Println("Using config file:", viper.ConfigFileUsed())
//     }
// }
