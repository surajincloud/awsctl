/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "awsctl",
	Short: "AWS CLI written in Go",
	Long: `AWS CLI written in
	       Golang`,
	PersistentPreRun: func(cmd *cobra.Command, args []string){
		region := cmd.Root().PersistentFlags().Lookup("region").Value.String()
		if region!=""{
			os.Setenv("AWS_DEFAULT_REGION",region)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().String("region", "", "Search resources in a specific region")
}
