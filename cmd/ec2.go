/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
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
		tag, _ := cmd.Flags().GetString("tags")
		if tag != "" {
			getEc2instanceTag(cmd, args)
		} else {
			getEC2command(cmd, args)
		}

	},
}

func getEc2instanceTag(cmd *cobra.Command, args []string) {
	var ec2Instance []EC2.EC2Instance
	Keys, _ := cmd.Flags().GetString("tags")

	if strings.Contains(Keys, "=") { // Both a Key and Values are passed
		tags := strings.SplitN(Keys, "=", 2)
		value := strings.Split(tags[1], ",")
		ec2Instance = EC2.GetInstanceWithKeyValue(&tags[0], value)
	} else { // Only if a Key is passed
		ec2Instance = EC2.GetInstanceWithKeyOnly(&Keys)
	}

	w := tabwriter.NewWriter(os.Stdout, 18, 5, 3, ' ', tabwriter.TabIndent)
	defer w.Flush()

	fmt.Fprintln(w, "KEYS", "\t", "INSTANCE_ID", "\t", "INSTANCE_TYPE", "\t", "PRIVATE IP")
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

var keyValue string

func init() {
	getCmd.AddCommand(ec2Cmd)
	ec2Cmd.Flags().StringVarP(&keyValue, "tags", "t", keyValue, "get instance using key value Example:--tags key=value1,value2,value3")
}
