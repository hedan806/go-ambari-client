package commands

import (
	"github.com/hedan806/go-ambari-client/internal/cmd/generate/utils"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate Go APIs, tests and examples for documentation",
	// Long:  "TODO",
}

// Execute launches the CLI application.
//
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.PrintErr(err)
		os.Exit(1)
	}
}

// RegisterCmd adds a command to rootCmd.
//
func RegisterCmd(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
