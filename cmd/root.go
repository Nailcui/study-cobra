package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(
		versionCmd,
		echoCmd,
		)
}

var rootCmd = &cobra.Command{
	Use:   "study-cobra",
	Short: "StudyCobra is a project for study cobra library.",
	Long: `这里是一些长的描述：
                关于库的描述可以写到这里
                当输入-h的时候便可以看到这些信息`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", strings.Join(args, " "))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
