package cmd

import (
	"achinour/cup/pkg/runner"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:     "up",
	Aliases: []string{"u"},
	Short:   "run docker compose up in detach mode on project",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runner.Run("up", args[0])
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
