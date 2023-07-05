package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/surajincloud/awsctl/pkg/ec2"
)

var sgCmd = &cobra.Command{

	Use:   "sg",
	Short: "Print Security Groups",
	Long: `
		For Example: $ awsctl get securitygroup
					 $ awsctl get securitygroups
					 $ awsctl get sg
	`,
	Run: getSG,
}

var securitygrp=&cobra.Command{
	Use: "securitygroup",
	Short: "Print Security Groups",
	Run: getSG,
}
var securitygrps=&cobra.Command{
	Use:"securitygroups",
	Short: "Print Security Groups",
	Run:getSG,
}

var group=[]*cobra.Command{
	securitygrp,
	securitygrps,
	sgCmd,
}

func getSG(cmd *cobra.Command, args []string) {

	var sgroup []ec2.SecurityGroup

	sgroup, err := ec2.DescribeSecurityGroup(cmd, args)
	if err != nil {
		log.Fatal("Unable to get Security Group")
	}

	w := tabwriter.NewWriter(os.Stdout, 18, 5, 3, ' ', tabwriter.TabIndent)
	defer w.Flush()

	fmt.Fprintln(w, "NAME", "\t", "GROUP ID", "\t", "DESCRIPTION")

	for _, i := range sgroup {

		fmt.Fprintln(w, i.SGName, "\t",
			i.SGId, "\t",
			i.SGDescription, "\t",
		)

	}

}



func init() {
	getCmd.AddCommand(group...)
}
