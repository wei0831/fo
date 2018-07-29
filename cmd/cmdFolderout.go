package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wei0831/fo/utli"
)

var cmdFolderout = &cobra.Command{
	Use:   "folderout",
	Short: "Move files out of folders",
	Long: `Move files inside child folders to target directory 
	and then delete those empty child folders.`,
	Run: func(cmd *cobra.Command, args []string) {
		utli.Folderout(dir, to, wet)
	},
}

func init() {
	rootCmd.AddCommand(cmdFolderout)
	cmdFolderout.PersistentFlags().BoolVarP(&wet, "wet", "w", false, "wet run, commit changes")
	cmdFolderout.PersistentFlags().StringVarP(&dir, "dir", "d", ".", "search target dir")
	cmdFolderout.PersistentFlags().StringVarP(&to, "to", "t", "", "move to target dir")
}
