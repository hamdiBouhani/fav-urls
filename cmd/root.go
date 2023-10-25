package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "favlinks",
	Short: "favlinks",
}

func registerCommands() {
	rootCmd.AddCommand(NewMigrateCmd())
	rootCmd.AddCommand(NewAddUrlCmd())
	rootCmd.AddCommand(NewListUrlCmd())
}

func Execute() {
	registerCommands()
	cobra.CheckErr(rootCmd.Execute())
}
