package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}

var rootCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "scaffold-cli",
	Long:  "scaffold-cli controls the Scaffold.",
}

func Execute() error {
	return rootCmd.Execute()
}
