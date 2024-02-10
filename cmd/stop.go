package cmd

import (
	"achinour/cup/pkg/runner"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"s"},
	Short:   "run docker compose stop on project",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runner.Run("stop", args[0])
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
