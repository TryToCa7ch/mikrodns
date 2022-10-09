package cmd

import (
	"fmt"
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
			fmt.Println("there's err on connection", err)
			os.Exit(1)
		}
		var record utils.DnsRecord
		fmt.Print("Insert ip address of endpoint ")
		fmt.Scan(&record.Address)
		fmt.Print("Insert dns name of record ")
		fmt.Scan(&record.Host)
		record.Disabled = "false"
		status_code := utils.AddDnsRecord(client, record.Host, record.Address)
		fmt.Println(status_code)
	},
}

func init() {
	rootCmd.AddCommand(addRecsCmd)
}
