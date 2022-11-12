/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "This will get the version detail",
	Long:  `This will get the version detail`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("sshm %s ", Version)

		color.Cyan("                    - Author:Dhruv (BugLightYear)")

		var logo string = `
                                            &&&&&&&        &&&&&&&
                                         &&&&&&&&&&        &&&&&&&&&&
                                       &&&&&&&&&&&&        &&&&&&&&&&&&
                                      &&&&&&&&&                &&&&&&&&&
                                      &&&&&&&&&                &&&&&&&&&	
                                      &&&&&&&&&                &&&&&&&&&
                                      &&&&&&&&&                &&&&&&&&&	
                                      &&&&&&&&&                &&&&&&&&&	
                                      &&&&&&&&&   &&&    &&&   &&&&&&&&&
                                    &&&&&&&&&     &&&    &&&    &&&&&&&&&
                                  &&&&&&&&&                       &&&&&&&&&
                                &&&&&&&&&                           &&&&&&&&&	
                                  &&&&&&&&&         **  **        &&&&&&&&&		      
                                    &&&&&&&&&         &&         &&&&&&&&&
                                       &&&&&&&&&               &&&&&&&&&  
                                       &&&&&&&&&               &&&&&&&&& 
                                       &&&&&&&&&   &&     &&   &&&&&&&&&
                                       &&&&&&&&&   &&&&&&&&&   &&&&&&&&&	
                                       &&&&&&&&&               &&&&&&&&&
                                       &&&&&&&&&               &&&&&&&&&	
                                       &&&&&&&&&&&&         &&&&&&&&&&&&	
                                         &&&&&&&&&&         &&&&&&&&&&
                                            &&&&&&&         &&&&&&&
                                
    `
		color.Red(logo)

	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
