package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/garbhank/flowtool/pkg/utils"
)

func init() {
	rootCmd.AddCommand(templateCmd)
}

var templateCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"templ", "tmp", "t"},
	Short:   "Template airflow variables and add it to system clipboard",
	Long:    `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("First argument must be a sql file")
			os.Exit(1)
		}

		sqlFilename := args[0]

		var m map[string]string
		env, _ := cmd.Flags().GetString("env")

		m = utils.CreateMapping(env)
		m = utils.AddAirflowVars(m)

		// read in sql file
		templatedSQL := utils.TemplateSQLFile(sqlFilename, m)

		// Check if the "create or replace" was left in
		utils.ValidateSQL(templatedSQL)

		// Send the templated string to the clipboard (doesn't work on linux)
		utils.ExportToClipboard(templatedSQL)

		var beQuiet bool
		cfgFileQuiet := viper.GetBool("quiet")
		quietFlag, _ := cmd.Flags().GetBool("quiet")

		// if no quiet flag set (defaults to false), use config file setting
		if quietFlag {
			beQuiet = quietFlag
		} else {
			beQuiet = cfgFileQuiet
		}

		if !beQuiet {
			fmt.Print(utils.ClipboardToString())
		}
	},
}
