package cmd

import (
	"strings"

	"github.com/nanernunes/scaffold"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a Scaffold Project",
	Long:  `You must specify the name of the project.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   New,
}

func New(cmd *cobra.Command, args []string) {
	if len(args) == 1 {
		project := strings.ToLower(args[0])
		scaffold.NewTemplate(project).Skel()
	}
}
