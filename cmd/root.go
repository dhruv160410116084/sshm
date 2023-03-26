/*
Copyright © 2022 Dhruv <dhruv.160410116084@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type SshCon struct {
	Index    int
	Project  string
	Env      string
	Shortcut string
	Con      string
}

type ConList struct {
	List []SshCon
}

var SshIndexMap = make(map[int]SshCon)
var SshShortMap = make(map[string]SshCon)
var Version = "v1.1.7"

var DATA_PATH string = ""
var SshList ConList

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "sshm",
	Short:             "sshm is a simple cli tool to manage your ssh connections",
	Long:              `sshm is a simple cli tool to manage your ssh connections. You can add,edit and delete the connections.`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	dataPath := "./data.json"

	DATA_PATH = dataPath

	SshList = LoadData()
	// fmt.Println(SshIndexMap, SshShortMap)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sshm.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.CompletionOptions.DisableDefaultCmd = true

}
