/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete ssh connection from sshm list",
	Long: `Delete ssh connection from sshm list.
			sshm delete index | shrotcut
Example:
		sshm delete 2
		sshm delete google-s`,
	Run: func(cmd *cobra.Command, args []string) {

		index, err := strconv.Atoi(args[0])
		if err == nil && index < len(SshList.List) {
			SshList.List[index] = SshList.List[len(SshList.List)-1]
			fmt.Println(SshList.List[index].Project, SshList.List[index].Env, "deleted")
			SshList.List = SshList.List[:len(SshList.List)-1]
			SaveData(SshList)

		} else {
			for _, v := range SshList.List {
				if v.Shortcut == args[0] {
					index := v.Index
					SshList.List[index] = SshList.List[len(SshList.List)-1]
					fmt.Println(SshList.List[index].Project, SshList.List[index].Env, "deleted")
					SshList.List = SshList.List[:len(SshList.List)-1]
					SaveData(SshList)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
