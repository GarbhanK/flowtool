package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/garbhank/flowtool/pkg/utils"
)

func init() {
    configCmd.AddCommand(listCmd)
    configCmd.AddCommand(addCmd) 
    rootCmd.AddCommand(configCmd)
}

// I want to take from this
// https://dev.to/divrhino/building-an-interactive-cli-app-with-go-cobra-promptui-346n

var configCmd = &cobra.Command{
    Use: "config",
    Short: "Add a key/value entry to the config file",
    Long: `All software has versions. This is Hugo's`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Config base. subcommands 'add', 'remove' or 'list'")
    },
}

var listCmd = &cobra.Command{
    Use: "list",
    Short: "List current config",
    Long: `ay the beat go off`,
    Run: func(cmd *cobra.Command, args []string) {
        utils.ListConfig()
    },
}

var addCmd = &cobra.Command{
    Use: "add",
    Short: "Add a key/value entry to the config file",
    Long: `1234`,
    Run: func(cmd *cobra.Command, args []string) {
        var key string
        var val string

        config := utils.ReadConfig()
        fmt.Println(config)

        // Take key/val from user input
        fmt.Println("Enter new key: ") 
        fmt.Scanln(&key)
        fmt.Println("Enter new val: ") 
        fmt.Scanln(&val)
        
        utils.AddConfig(config, key, val)
        fmt.Println("Entry added!")
    },
}
