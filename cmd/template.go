package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/garbhank/flowtool/pkg/template"
	"github.com/garbhank/flowtool/pkg/utils"
)

func init() {
	rootCmd.AddCommand(templateCmd)
}

var templateCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"templ", "tmp", "t"},
	Short:   "Template airflow variables and add it to system clipboard",
	Long:    `Template airflow variables and add it to system clipboard`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("First argument must be a sql file")
			os.Exit(1)
		}
		env, _ := cmd.Flags().GetString("env")

		t := template.NewTemplater(args[0], env)

		// add airflow specific template variables to the current mapping config
		t.AddAirflowVars()

		// find and replace template variables in the file
		t.TemplateSQLFile()

		// Warn if sql statement has create/insert/update/delete/drop
		t.ValidateSQL()

		// Send the templated string to the clipboard (doesn't work on linux)
		utils.ExportToClipboard(t.FileTemplated)

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
