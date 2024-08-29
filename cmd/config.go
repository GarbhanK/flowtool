package cmd

import (
	"fmt"

	"github.com/garbhank/flowtool/pkg/utils"
	"github.com/spf13/cobra"
)

func init() {
	// top level command
	rootCmd.AddCommand(configCmd)

	// subcommands
	configCmd.AddCommand(listCmd)
	configCmd.AddCommand(addCmd)
	configCmd.AddCommand(removeCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Used to 'add', 'remove' or 'list' current config",
	Long: `Top level command for interacting with the config.json file.
            Subcommands are 'add', 'remove' and 'list'.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Config base. subcommands 'add', 'remove' or 'list'")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List current config",
	Long:  `Display the current key/value pairs in the current config.json.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ListConfig()
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a key/value entry to the config file",
	Long:  `Take user input to set a new key/value pair and write it to config.json.`,
	Run: func(cmd *cobra.Command, args []string) {
		var key string
		var val string

		config := utils.ReadConfig()
		utils.ListConfig()

		// Take key/val from user input
		fmt.Println("\nEnter new key: ")
		fmt.Scanln(&key)
		fmt.Println("Enter new val: ")
		fmt.Scanln(&val)

		utils.AddConfig(config, key, val)
		fmt.Println("\nEntry added!")
	},
}

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "remove key/value pairs from the config.json",
	Run: func(cmd *cobra.Command, args []string) {
		var key string

		config := utils.ReadConfig()
		utils.ListConfig()

		fmt.Println("\nEnter key you want to remove: ")
		fmt.Scanln(&key)

		utils.RemoveConfig(config, key)
		fmt.Printf("Successfully removed %s from config!\n", key)
	},
}
