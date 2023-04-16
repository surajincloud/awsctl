package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/cobra"
	"github.com/surajincloud/awsctl/pkg/vpc"
)

var vpcCmd = &cobra.Command{
	Use:   "vpc",
	Short: "Get VPCs",
	Long:  "Get Vpcs with Name and Cidr",
	RunE:  getVPC,
}

func getVPC(cmd *cobra.Command, args []string) error {

	vpclist, err := vpc.GetVPC()
	if err != nil {
		log.Fatal(err)
	}
	w := tabwriter.NewWriter(os.Stdout, 18, 5, 2, ' ', tabwriter.TabIndent)
	defer w.Flush()

	fmt.Fprintln(w, "NAME", "\t", "STATE", "\t", "DEFAULT", "\t", "IPv4CIDR", "\t", "VPCID")
	for _, i := range vpclist {
		fmt.Fprintln(w,
			aws.String(*i.Name), "\t",
			aws.String(*i.State), "\t",
			aws.Bool(i.Default), "\t",
			aws.String(*i.Ipv4Cidr), "\t",
			aws.String(*i.VpcID),
		)
	}
	return nil
}

func init() {
	getCmd.AddCommand(vpcCmd)
}
