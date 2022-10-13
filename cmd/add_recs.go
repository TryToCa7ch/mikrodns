package cmd

import (
	"fmt"
	"mikrodns/color_print"
	"mikrodns/utils"
	"os"

	"github.com/spf13/cobra"
)

var addRecsCmd = &cobra.Command{
	Use:     "add_records",
	Aliases: []string{"add"},
	Short:   "Add dns static record",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := utils.Dial()
		if err != nil {
			fmt.Println(color_print.Fata("there's err on connection", err))
			os.Exit(1)
		}
		var record utils.DnsRecord
		fmt.Print("Insert ip address of endpoint\n")
		fmt.Scan(&record.Address)
		fmt.Print("Insert dns name of record\n")
		fmt.Scan(&record.Host)
		record.Disabled = "false"
		added_record, err := utils.AddDnsRecord(client, record.Host, record.Address)
		if err != nil {
			fmt.Println(color_print.Fata("There's error when trying to add new record: \n", err))
		} else {
			fmt.Println(added_record.Id, added_record.Address, added_record.Host, added_record.Disabled)
		}
	},
}

func init() {
	rootCmd.AddCommand(addRecsCmd)
}
