/*
 * Copyright (c) 2023 NAME HERE <EMAIL ADDRESS>
 */

package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	// completionCmd represents the completion command
	completionCmd = &cobra.Command{
		Use:     "completion [bash|zsh|powershell|fish]",
		Short:   "Generates completion scripts for various shells",
		Example: "awsctl completion bash",
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
	if len(args) != 1 {
		return errors.New("invalid command")
	}
	if args[0] != "bash" && args[0] != "zsh" && args[0] != "fish" && args[0] != "powershell" {
		return fmt.Errorf("sorry, completion support is not yet implemented for %v", args[0])
	}
	return nil
}

// bashCompletion generates bash completion scripts
func bashCompletion(cmd *cobra.Command, args []string) {
	err := rootCmd.GenBashCompletion(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

// fishCompletion generates fish completion scripts
func fishCompletion(cmd *cobra.Command, args []string) {
	err := rootCmd.GenFishCompletion(os.Stdout, true)
	if err != nil {
		log.Fatal(err)
	}
}

// powerShellCompletion generates PowerShell completion scripts
func powerShellCompletion(cmd *cobra.Command, args []string) {
	err := rootCmd.GenPowerShellCompletion(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

// zshCompletion generates Zsh completion scripts
func zshCompletion(cmd *cobra.Command, args []string) {
	err := rootCmd.GenZshCompletion(os.Stdout)
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
