package cmd

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
	"github.com/nanernunes/scaffold"
	"github.com/nanernunes/scaffold/pkg/helpers"
	"github.com/spf13/cobra"
)

var types = map[string]string{
	"string": "string",
	"int":    "int",
	"float":  "float64",
	"bool":   "bool",

	"references": "references",
}

func init() {
	rootCmd.AddCommand(newGenerate)
}

var newGenerate = &cobra.Command{
	Use:     "generate",
	Short:   "Generate a MVC Resource",
	Aliases: []string{"g"},
	Long: `You must specify the name of the resource in CamelCase.
	Use "scaffold generate" for a complete help description.`,
	Args: cobra.MinimumNArgs(1),
	Run:  Generate,
}

func Generate(cmd *cobra.Command, args []string) {
	if len(args) >= 1 {
		here, _ := os.Getwd()
		project := path.Base(here)

		resource := helpers.NewName(args[0]).Camel().ToString()
		template := scaffold.NewTemplate(project)

		fields := make(map[*helpers.Name]string)

		for _, tuple := range args[1:] {
			field := strings.Split(tuple, ":")

			if _, ok := types[field[1]]; !ok {
				keys := make([]string, 0, len(types))
				for key := range types {
					keys = append(keys, key)
				}
				log.Fatalf(
					"type %s is not an allowed option, try one these: %s.",
					color.RedString(field[1]),
					color.GreenString(strings.Join(keys, ", ")),
				)
			}

			if field[1] == "references" {
				if _, err := os.Stat(path.Join("app", "models", strings.ToLower(field[0])+".go")); os.IsNotExist(err) {
					log.Fatalf(
						"there is no %s model in your project, try creating it before referencing it.",
						color.RedString(helpers.NewName(field[0]).Lower().Camel().ToString()),
					)
				}
			}

			fields[helpers.NewName(field[0])] = types[field[1]]
		}

		template.AddMVC(resource, fields)
	}
}
