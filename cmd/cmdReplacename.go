package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wei0831/fo/utli"
)

var cmdReplaceName = &cobra.Command{
	Use:   "replacename find replace",
	Short: "replacename given the search pattern and replace pattern.",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		utli.ReplaceName(dir, to, args[0], args[1], exclude, mode, wet)
	},
}

func init() {
	rootCmd.AddCommand(cmdReplaceName)
	cmdReplaceName.PersistentFlags().BoolVarP(&wet, "wet", "w", false, "wet run, commit changes")
	cmdReplaceName.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "search target dir")
	cmdReplaceName.PersistentFlags().StringVarP(&to, "to", "t", "", "move to target dir")
	cmdReplaceName.PersistentFlags().StringVarP(&exclude, "exclude", "e", "", "exclude the pattern")
	cmdReplaceName.PersistentFlags().IntVarP(&mode, "mode", "m", 0, "0: File Only 1: Folder Only 2: Both")
}
