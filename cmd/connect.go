/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "This will run ssh command with your selected connection",
	Long: `This will run ssh command with your selected connection.
			sshm connect index | shortcut
	Example:
			sshm connect 1
			sshm connect fb-stage`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("connect called")
		// command := exec.Command("ssh", "ubuntu@3.11.205.213")
		var command *exec.Cmd

		index, err := strconv.Atoi(args[0])
		if err == nil && index < len(SshList.List) {
			fmt.Println(SshList.List[index].Con)
			ssh_cmd := strings.Split(SshList.List[index].Con, " ")
			command = exec.Command(ssh_cmd[0], ssh_cmd[1:]...)

		} else if val, ok := SshShortMap[args[0]]; ok {
			ssh_cmd := strings.Split(val.Con, " ")
			command = exec.Command(ssh_cmd[0], ssh_cmd[1:]...)

		} else {
			fmt.Println("No match found")
			return
		}

		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		e := command.Run()
		if e != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
