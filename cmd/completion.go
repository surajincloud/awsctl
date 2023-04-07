/*
 * Copyright (c) 2023 NAME HERE <EMAIL ADDRESS>
 */

package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
)

var (
	// completionCmd represents the completion command
	completionCmd = &cobra.Command{
		Use:     "completion [bash|zsh|powershell|fish]",
		Short:   "Generates completion scripts for various shells",
		Example: "awsctl completion [bash|zsh|powershell|fish]",
		RunE:    completionCmdRun,
	}

	// completionBashCmd generates bash completion scripts
	completionBashCmd = &cobra.Command{
		Use:   "bash",
		Short: "Generates bash completion scripts",
		Run:   bashCompletion,
	}

	// completionFishCmd generates fish completion scripts
	completionFishCmd = &cobra.Command{
		Use:   "fish",
		Short: "Generates fish completion scripts",
		Run:   fishCompletion,
	}

	// completionPowerShellCmd generates PowerShell completion scripts
	completionPowerShellCmd = &cobra.Command{
		Use:   "powershell",
		Short: "Generates PowerShell completion scripts",
		Run:   powerShellCompletion,
	}

	// completionZshCmd generates Zsh completion scripts
	completionZshCmd = &cobra.Command{
		Use:   "zsh",
		Short: "Generates Zsh completion scripts",
		Run:   zshCompletion,
	}
)

// completionCmdRun runs the completion command
func completionCmdRun(cmd *cobra.Command, args []string) error {
	err := cmd.Help()

	// TODO : Write better error message
	if len(args) != 1 {
		return errors.New("invalid command")
	}
	if args[0] != "bash" && args[0] != "zsh" && args[0] != "fish" && args[0] != "powershell" {
		log.Fatalf("sorry, completion support is not yet implemented for %v", args)
	}

	if err != nil {
		log.Fatal(err)
	}
	return errors.New("a valid subcommand is required")
}

// bashCompletion generates bash completion scripts
func bashCompletion(cmd *cobra.Command, args []string) {
	err := rootCmd.GenBashCompletionFile("awsctl-completion.sh")
	if err != nil {
		log.Fatal(err)
	}
}

// fishCompletion generates fish completion scripts
func fishCompletion(cmd *cobra.Command, args []string) {
	err := rootCmd.GenFishCompletionFile("awsctl-completion.fish", true)
	if err != nil {
		log.Fatal(err)
	}
}

// powerShellCompletion generates PowerShell completion scripts
func powerShellCompletion(cmd *cobra.Command, args []string) {
	err := rootCmd.GenPowerShellCompletionFile("awsctl-completion.ps1")
	if err != nil {
		log.Fatal(err)
	}
}

// zshCompletion generates Zsh completion scripts
func zshCompletion(cmd *cobra.Command, args []string) {
	err := rootCmd.GenZshCompletionFile("awsctl-completion.zsh")
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	//adding subcommands to completion command
	completionCmd.AddCommand(completionBashCmd)
	completionCmd.AddCommand(completionZshCmd)
	completionCmd.AddCommand(completionFishCmd)
	completionCmd.AddCommand(completionPowerShellCmd)

	//adding completion command to root
	rootCmd.AddCommand(completionCmd)
}
