/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your ssh connections",
	Long: `This command will list your ssh connections list. For example:

sshm list`,
	// ValidArgs: []string{"fb", "google"},
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("list called")
		// getConnection("st")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Index", "Project", "Env", "Shrotcut", "Con"})
		table.SetAutoWrapText(false)

		for _, v := range SshList.List {
			_t := []string{strconv.Itoa(v.Index), v.Project, v.Env, v.Shortcut, v.Con}
			table.Append(_t)
		}
		table.Render()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
