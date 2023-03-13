/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/surajincloud/awsctl/pkg/EC2"
)

var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Print ec2 related information",
	Long: `For example,
		awsctl get ec2`,
	Run: func(cmd *cobra.Command, args []string) {
		tag, _ := cmd.Flags().GetStringSlice("tags")
		if len(tag) != 0 {
			getEc2instanceTag(cmd, args)
		} else {
			getEC2command(cmd, args)
		}

	},
}
var keyPair = [2]string{"", ""}

func getEc2instanceTag(cmd *cobra.Command, args []string) {
	var ec2Instance []EC2.EC2Instance
	tags, _ := cmd.Flags().GetStringSlice("tags")

	ec2Instance = EC2.GetEC2InstanceTag(&tags[0], tags[1:])

	w := tabwriter.NewWriter(os.Stdout, 18, 5, 3, ' ', tabwriter.TabIndent)
	defer w.Flush()

	fmt.Fprintln(w, "TagValue", "\t", "INSTANCE_ID", "\t", "INSTANCE_TYPE", "\t", "PRIVATE IP")
	for _, i := range ec2Instance {
		fmt.Fprintln(
			w, i.InstanceName, "\t",
			i.InstanceID, "\t",
			i.InstanceType, "\t",
			i.InstancePrivateIP)
	}
}

func getEC2command(cmd *cobra.Command, args []string) {

	var ec2Instance []EC2.EC2Instance
	ec2Instance = EC2.GetEC2Instance()
	w := tabwriter.NewWriter(os.Stdout, 18, 5, 3, ' ', tabwriter.TabIndent)
	defer w.Flush()

	fmt.Fprintln(w, "NAME", "\t", "INSTANCE_ID", "\t", "INSTANCE_TYPE", "\t", "PRIVATE IP")
	for _, i := range ec2Instance {
		fmt.Fprintln(
			w, i.InstanceName, "\t",
			i.InstanceID, "\t",
			i.InstanceType, "\t",
			i.InstancePrivateIP)
	}

}

func init() {
	getCmd.AddCommand(ec2Cmd)
	ec2Cmd.Flags().StringSlice("tags", []string{}, "get instance using key value Example: --tags=key,value1,value2")
}
