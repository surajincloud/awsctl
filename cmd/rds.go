/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/surajincloud/awsctl/pkg/rds"
)

// rdsCmd represent the rds command
var rdsCmd = &cobra.Command{
	Use:   "rds",
	Short: "Retrieve information about AWS RDS instances",
	Long: `This command retrieves information about AWS RDS instances, including their instance ID, status, endpoint, engine and version.
Example:
	awsctl get rds`,
	RunE: getRdsCommand,
}

func getRdsCommand(cmd *cobra.Command, args []string) error {

	rdsInstances, err := rds.GetRDSInstances()

	if err != nil {
		return err
	}

	if len(rdsInstances) == 0 {
		fmt.Println("No RDS instances found")
		return nil
	}
	w := tabwriter.NewWriter(os.Stdout, 5, 2, 3, ' ', tabwriter.TabIndent)
	defer w.Flush()

	fmt.Fprintln(w, "DB_INSTANCE_ID", "\t", "STATUS", "\t", "ENDPOINT", "\t", "ENGINE", "\t", "VERSION")

	for _, i := range rdsInstances {
		fmt.Fprintln(
			w,
			i.DBInstanceID, "\t",
			i.Status, "\t",
			i.Endpoint, "\t",
			i.Engine, "\t",
			i.Version, "\t",
		)
	}
	return nil
}

func init() {
	getCmd.AddCommand(rdsCmd)
}
