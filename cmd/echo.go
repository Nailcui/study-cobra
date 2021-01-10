package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var isUp = false
var isLower = false
var separator = " "

func init() {
	echoCmd.PersistentFlags().BoolVarP(&isUp, "up", "u", false, "转大写输出")
	echoCmd.PersistentFlags().BoolVarP(&isLower, "lower", "l", false, "转小写输出")
	echoCmd.PersistentFlags().StringVarP(&separator, "separator", "s",  " ", "自定义分隔符")
}

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "鹦鹉学舌",
	Long:  `鹦鹉学舌:
                你说什么我就回复什么`,
	Run: func(cmd *cobra.Command, args []string) {
		output := strings.Join(args, separator)
		if isUp {
			output = strings.ToUpper(output)
		}
		if isLower {
			output = strings.ToLower(output)
		}
		fmt.Println(output)
	},
}
