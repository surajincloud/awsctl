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

	Use: "sgrp",
	Short: "Print Security Groups",
	Long: `
		For Example: $ awsctl get sgrp
	`,
	Run: getSG,

}

func getSG(cmd *cobra.Command, args[]string){

	var sgroup []ec2.SECURITYGRP

	sgroup,err:= ec2.DescribeSecurityGroup(cmd,args)
	if err!=nil{
		log.Fatal("Unable to get Security Group")
	}

	w:= tabwriter.NewWriter(os.Stdout,18,5,3,' ',tabwriter.TabIndent)
	defer w.Flush()

	fmt.Fprintln(w,"NAME","\t","GROUP ID","\t","DESCRIPTION")

	for _,i:= range sgroup{

		fmt.Fprintln(w,i.SGName,"\t",
	                 i.SGId,"\t",
					 i.SGDescription,"\t",
	)	

	}



}

func init(){
	getCmd.AddCommand(sgCmd)
}
