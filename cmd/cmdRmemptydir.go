package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wei0831/fo/utli"
)

var cmdRmemptydir = &cobra.Command{
	Use:   "rmemptydir",
	Short: "remove empty folders.",
	Run: func(cmd *cobra.Command, args []string) {
		utli.RmEmptyDir(dir, wet)
	},
}

func init() {
	rootCmd.AddCommand(cmdRmemptydir)
	cmdRmemptydir.PersistentFlags().BoolVarP(&wet, "wet", "w", false, "wet run, commit changes")
	cmdRmemptydir.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "search target dir")
}
