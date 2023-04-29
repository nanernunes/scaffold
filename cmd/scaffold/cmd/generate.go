package cmd

import (
	"os"
	"path"
	"strings"

	"github.com/nanernunes/scaffold"
	"github.com/nanernunes/scaffold/pkg/helpers"
	"github.com/spf13/cobra"
)

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
			fields[helpers.NewName(field[0])] = field[1]
		}

		template.AddMVC(resource, fields)
	}
}
