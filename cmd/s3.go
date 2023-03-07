/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
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
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	out, err := client.ListBuckets(ctx,&s3.ListBucketsInput{})
	
	

	if err != nil {
		log.Fatal(err)
	}
	w := tabwriter.NewWriter(os.Stdout, 10, 5, 2, ' ', tabwriter.TabIndent)
	defer w.Flush()
	fmt.Fprintln(w, "NAME", "\t","CREATED_AT","\t","SIZE")
	
	for _, i := range out.Buckets {
		info,err:= client.ListObjectsV2(ctx,&s3.ListObjectsV2Input{
			Bucket: aws.String(*i.Name),
		})
		if err!=nil{
			log.Fatal("Error Occured",err)
		}
		for _,j:=range info.Contents{
			size:= j.Size/1024
			fmt.Fprintln(w, aws.ToString(i.Name), "\t", i.CreationDate,"\t",size,"KB")
		}
		
		
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
