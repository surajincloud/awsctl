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
	"github.com/surajincloud/awsctl/pkg/ec2"
)

var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Print ec2 related information",
	Long: `For example,
		awsctl get ec2`,
	Run: getEC2,
}


func getEC2(cmd *cobra.Command, args []string) {

	var ec2Instance []ec2.EC2Instance
	Keys, _ := cmd.Flags().GetString("tags")
	if Keys!="" && strings.Contains(Keys, "=") {
		tags := strings.SplitN(Keys, "=", 2)
		value := strings.Split(tags[1], ",")
		ec2Instance = ec2.DescribeInstance(*&tags[0], value,cmd,args)
	} else if Keys!=""{
		ec2Instance = ec2.DescribeInstance(*&Keys, nil,cmd,args)
	}else{
		ec2Instance = ec2.DescribeInstance("", nil,cmd,args)
	}
	
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
