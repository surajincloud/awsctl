/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"github.com/spf13/cobra"
	"github.com/surajincloud/awsctl/pkg"
)

// ec2Cmd represents the ec2 command
var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Print ec2 related information",
	Long: `For example,
		awsctl get ec2`,
	Run: getEC2command,
}

func getEC2command(cmd *cobra.Command, args []string) {

	var ec2Instance [] pkg.EC2Instance
	ec2Instance =pkg.GetEC2Instance()
	w:= tabwriter.NewWriter(os.Stdout,18,5,3,' ',tabwriter.TabIndent)
	defer w.Flush()
	fmt.Fprintln(w,"NAME", "\t" ,"INSTANCE_ID","\t" ,"INSTANCE_TYPE", "\t", "PRIVATE IP")
	for _,i:=range ec2Instance{
		fmt.Fprintln(
		w,i.InstanceName,"\t",
		i.InstanceID,"\t",
		i.InstanceType,"\t",
		i.InstancePrivateIP,)
	}
	
}

func init() {
	getCmd.AddCommand(ec2Cmd)
}
