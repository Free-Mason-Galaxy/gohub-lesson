// Package cmd
// descr
// author fm
// date 2022/11/22 14:27
package cmd

import (
	"github.com/spf13/cobra"
	"gohub-lesson/pkg/console"
	"gohub-lesson/pkg/helpers"
)

var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "Generate App Key, will print the generated Key",
	Run:   runGenerateKey,
	Args:  cobra.NoArgs,
}

// runGenerateKey 生成 key
func runGenerateKey(cmd *cobra.Command, args []string) {
	console.Success("---")
	console.Success("App Key:")
	console.Success(helpers.RandomString(32))
	console.Success("---")
	console.Warning("please go to .env file to change the APP_KEY option")
}
