package cmd

import (
	"achinour/cup/pkg/runner"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:     "down",
	Short:   "run docker compose down on project",
	Aliases: []string{"d"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runner.Run("down", args[0])
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
