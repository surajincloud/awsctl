package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

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

	fmt.Fprintln(w, "NAME", "\t", "State", "\t", "DEFAULT", "\t", "IPv4CIDR", "\t", "VPCID")
	for _, i := range vpclist {
		fmt.Fprintln(w,
			*i.Name, "\t",
			*i.State, "\t",
			i.Default, "\t",
			*i.Ipv4Cidr, "\t",
			*i.VpcID)
	}
	return nil
}

func init() {
	getCmd.AddCommand(vpcCmd)
}
