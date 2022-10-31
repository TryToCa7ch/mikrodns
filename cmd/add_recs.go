package cmd

import (
	"fmt"
	"log"
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
			log.Println(color_print.Fata("there's err on connection", err))
			os.Exit(1)
		}
		var record utils.DnsRecord
		log.Print("Insert ip address of endpoint\n")
		fmt.Scan(&record.Address)
		log.Print("Insert dns name of record\n")
		fmt.Scan(&record.Name)
		record.Disabled = "false"
		added_record, err := utils.AddDnsRecord(client, record.Name, record.Address)
		if err != nil {
			log.Println(color_print.Fata("There's error when trying to add new record: \n", err))
		} else {
			log.Println(added_record.Id, added_record.Address, added_record.Name, added_record.Disabled)
		}
	},
}

func init() {
	log.SetFlags(0)

	rootCmd.AddCommand(addRecsCmd)
}
