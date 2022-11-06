/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This will add your ssh command into sshm list",
	Long: `This will add your ssh command into sshm list. Example. 
			sshm add ssh ubuntu@3.10.136.168`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("add called")
		// fmt.Println(args)

		promptProj := promptui.Prompt{
			Label: "Enter Project Name: ",
			Validate: func(s string) error {
				if len(s) <= 3 {
					return errors.New("Project Name is too short")
				}
				return nil
			},
		}
		project, err := promptProj.Run()

		promptEnv := promptui.Prompt{
			Label: "Enter Env  (ex. stage,production): ",
			Validate: func(s string) error {
				if len(s) <= 0 {
					return errors.New("Invalid Env")
				}
				return nil
			},
		}
		env, err := promptEnv.Run()

		promptCon := promptui.Prompt{
			Label: "Enter ssh connection: ",
			Validate: func(s string) error {
				if len(s) <= 4 {
					return errors.New("Invalid ssh connnection")
				}
				return nil
			},
		}
		con, err := promptCon.Run()

		promptShortcut := promptui.Prompt{
			Label: "Enter Shortcut or alias for connection: ",
			// Default: project[0:3] + env[0:1] + " ",
			Validate: func(s string) error {
				if len(s) <= 0 {
					return errors.New("Invalid Env")
				}
				return nil
			},
		}
		shortcut, err := promptShortcut.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		promptConfirm := promptui.Select{
			Label: "Are you sure you want to add ?",
			Items: []string{"Yes", "No"},
		}
		_, confirm, err := promptConfirm.Run()

		// fmt.Println(project, con, confirm)
		if confirm == "Yes" {
			_con := SshCon{Index: len(SshList.List), Env: env, Project: project, Con: con, Shortcut: shortcut}
			SshList.List = append(SshList.List, _con)
			SaveData(SshList)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
