package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flu",
	Short: "Fly Utilities",
	Long:  `Flu is a utility package for managing local Fly apps`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(lsCmd)
}

func initConfig() {
	// Placeholder for init func
}
