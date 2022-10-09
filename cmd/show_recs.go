package cmd

import (
	"fmt"
	"mikrodns/utils"
	"os"

	"github.com/spf13/cobra"
)

var showRecsCmd = &cobra.Command{
	Use:     "get_records",
	Aliases: []string{"get_all"},
	Short:   "Get all dns static records",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := utils.Dial()
		if err != nil {
			fmt.Println("there's err on connection", err)
			os.Exit(1)
		}
		res := utils.GetAllDnsRecords(client)
		for _, rec := range res {
			fmt.Printf("ID: %s\tAddress: %s\tIP: %s\tDisabled: %s\n", rec.Id[1:], rec.Address, rec.Host, rec.Disabled)
		}
	},
}

var showRecCmd = &cobra.Command{
	Use:     "get_record",
	Aliases: []string{"get"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "Get  dns static record with given id",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := utils.Dial()
		if err != nil {
			fmt.Println("there's err on connection", err)
			os.Exit(1)
		}
		res, err := utils.GetDnsRecord(client, args[0])
		if err != nil {
			fmt.Println(utils.Fata("There's no dns record with given id"))
			os.Exit(1)
		}
		fmt.Printf("ID: %s\tAddress: %s\tIP: %s\tDisabled: %s\n", res.Id[1:], res.Address, res.Host, res.Disabled)
	},
}

func init() {
	rootCmd.AddCommand(showRecsCmd)
	rootCmd.AddCommand(showRecCmd)
}
