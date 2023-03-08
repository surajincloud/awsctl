/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/cobra"
	"github.com/surajincloud/awsctl/pkg/s3API"
)

// s3Cmd represents the s3 command
var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Print s3 related information",
	Long: `For example,
	  awsctl get s3`,
	RunE: getS3,
}

func getS3(cmd *cobra.Command, args []string) error {

	w := tabwriter.NewWriter(os.Stdout, 10, 5, 2, ' ', tabwriter.TabIndent)
	defer w.Flush()
	fmt.Fprintln(w, "NAME", "\t", "CREATED_AT", "\t", "SIZE")

	bucketList, err := s3API.GetBucket()
	if err != nil {
		log.Fatal("Unable to get Data From Aws")
	}

	for _, i := range bucketList {
		fmt.Fprintln(w, aws.String(i.Name), "\t", i.CreatedAt, "\t", i.Size, "KB")
	}

	return nil
}

func init() {
	getCmd.AddCommand(s3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// s3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// s3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
