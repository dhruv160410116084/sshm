/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"

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
var Version = "development"

var DATA_PATH string = ""
var SshList ConList

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sshm",
	Short: "sshm is a simple cli tool to manage your ssh connections",
	Long:  `sshm is a simple cli tool to manage your ssh connections. You can add,edit and view the connection strings.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	mydir, erri := os.Getwd()
	if erri != nil {
		fmt.Println(erri)
	}
	fmt.Println("current working dir is : ", mydir)
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	// fmt.Println(version)
	// fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	DATA_PATH = path.Dir(filename) + "/data.json"
	SshList = LoadData()
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
}
