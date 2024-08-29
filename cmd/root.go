package cmd

import (
	"fmt"
	// "os"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "flowtool",
		Short: "Template airflow vars",
		Long: `A fast and Flexible satic site generator built with
                    love by spf13 adn friends in Go.
                    Complete diocs available at link`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from root")
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
//     cobra.OnInitialize(initConfig)

// rootCmd.PersistentFlags().StringVarP(&Verbose, "verbose", "v", false, "verbose output")
rootCmd.PersistentFlags().String("env", "dev", "The desirev environment to template into config values")
}
//     rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "cinfig file (default")
//     rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
// 	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
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
