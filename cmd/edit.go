/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit ssh connection.",
	Long: `Edit ssh connection
		   sshm edit [index]
		Example:
		   sshm edit 0	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			return
		}
		index, err := strconv.Atoi(args[0])

		if err == nil && index < len(SshList.List) {
			_strCon := SshList.List[index]
			promptProj := promptui.Prompt{
				Label: "Enter Project Name: ",
				Validate: func(s string) error {
					if len(s) <= 3 {
						fmt.Println("Project Name is too short")

					}
					return nil
				},
				Default: _strCon.Project,
			}
			project, errProj := promptProj.Run()
			if errProj != nil {
				fmt.Println(err)
			}

			promptEnv := promptui.Prompt{
				Label:   "Enter Env  (ex. stage,production): ",
				Default: _strCon.Env,
				Validate: func(s string) error {
					if len(s) <= 0 {
						fmt.Println("Invalid Env")
					}
					return nil
				},
			}
			env, errEnv := promptEnv.Run()

			if errEnv != nil {
				fmt.Println(errEnv)
			}

			promptCon := promptui.Prompt{
				Label:   "Enter ssh connection (No spaces): ",
				Default: _strCon.Con,
				Validate: func(s string) error {
					if len(s) <= 4 {
						fmt.Println("Invalid ssh connnection")
					}
					return nil
				},
			}
			con, errCon := promptCon.Run()

			if errCon != nil {
				fmt.Println(errEnv)
			}

			promptShortcut := promptui.Prompt{
				Label: "Enter Shortcut or alias for connection: ",
				// Default: project[0:3] + env[0:1] + " ",
				Default: _strCon.Shortcut,
				Validate: func(s string) error {
					if len(s) <= 0 {
						fmt.Println("Invalid Env")
					}
					return nil
				},
			}
			shortcut, errShort := promptShortcut.Run()

			if errShort != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			promptConfirm := promptui.Select{
				Label: "Are you sure you want to add ?",
				Items: []string{"Yes", "No"},
			}
			_, confirm, err := promptConfirm.Run()
			if err != nil {
				fmt.Println(err)
			}

			// fmt.Println(project, con, confirm)
			if confirm == "Yes" {
				_con := SshCon{Index: index, Env: env, Project: project, Con: con, Shortcut: shortcut}
				SshList.List[index] = _con
				SaveData(SshList)
				fmt.Println("Ssh connect Updated!")
			} else {
				fmt.Println("Cancelled")
			}

			fmt.Println(project, env, con, shortcut)
		}

	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
