package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wei0831/fo/utli"
)

var cmdRevert = &cobra.Command{
	Use:   "revert logPath",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utli.Revert(args[0], wet)
	},
}

func init() {
	rootCmd.AddCommand(cmdRevert)
	cmdRevert.PersistentFlags().BoolVarP(&wet, "wet", "w", false, "wet run, commit changes")
}
