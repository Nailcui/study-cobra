package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of study-cobra",
	Long:  `All software has versions. This is study-cobra's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("study-cobra v0.9 -- HEAD")
	},
}
