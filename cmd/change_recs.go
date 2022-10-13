package cmd

import (
	"fmt"
	"mikrodns/color_print"
	"mikrodns/utils"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var changeRecsCmd = &cobra.Command{
	Use:     "change_records",
	Aliases: []string{"change"},
	Short:   "Change dns static record of given id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(color_print.Fata("Provide static record id"))
			os.Exit(1)
		}
		client, err := utils.Dial()
		if err != nil {
			fmt.Println(color_print.Fata("there's err on connection", err))
			os.Exit(1)
		}
		record, err := utils.GetDnsRecord(client, args[0])
		menu_status := 88
		menu_handler(record, menu_status)
		if err != nil {
			fmt.Println(color_print.Fata("There's error when trying to change record: \n", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(changeRecsCmd)
}

func menu_handler(record utils.DnsRecord, status int) {
	fmt.Println(status)
	switch s := strconv.Itoa(status); s {
	case "88":
		fmt.Printf("Your dns record is:\n ID: %s\tIp address: %s\tHostname: %s\tDisabled: %s\n", record.Id, record.Host, record.Address, record.Disabled)
		t := 88
		fmt.Printf("Pick waht you need to change:\n1. Ip address\n2. Hostname\n3. Disabled status\n4. End\n")
		fmt.Scanf("%d\n", &t)
		menu_handler(record, t)
	case "1":
		fmt.Printf("Your dns record is:\n ID: %s\tIp address: %s\tHostname: %s\tDisabled: %s\n", record.Id, record.Host, record.Address, record.Disabled)
		t := 1
		fmt.Printf("You changing Hostname.\n")
		fmt.Scan(&record.Address)
		fmt.Printf("Pick waht you need to change:\n1. Ip address\n2. Hostname\n3. Disabled status\n4. End\n")
		fmt.Scanf("%d\n", &t)
		menu_handler(record, t)
	case "2":
		fmt.Printf("Your dns record is:\n ID: %s\tIp address: %s\tHostname: %s\tDisabled: %s\n", record.Id, record.Host, record.Address, record.Disabled)
		t := 2
		fmt.Printf("You changing Ip address.\n")
		fmt.Scan(&record.Host)
		fmt.Printf("Pick waht you need to change:\n1. Ip address\n2. Hostname\n3. Disabled status\n4. End\n")
		fmt.Scanf("%d\n", &t)
		menu_handler(record, t)
	case "3":
		fmt.Printf("Your dns record is:\n ID: %s\tIp address: %s\tHostname: %s\tDisabled: %s\n", record.Id, record.Host, record.Address, record.Disabled)
		t := 3
		fmt.Printf("You changing Disabled status.\n")
		fmt.Scan(&record.Disabled)
		fmt.Printf("Pick waht you need to change:\n1. Ip address\n2. Hostname\n3. Disabled status\n4. End\n")
		fmt.Scanf("%d\n", &t)
		menu_handler(record, t)
	case "4":
		t := 4
		fmt.Printf("Your dns record is:\n ID: %s\tIp address: %s\tHostname: %s\tDisabled: %s\n", record.Id, record.Host, record.Address, record.Disabled)
		fmt.Printf("Confirm changing[Y/n]\n")
		var confirmation string
		fmt.Scanf("%s\n", &confirmation)
		if confirmation == "Y" || confirmation == "y" {
			rec, err := utils.ChangeDnsRecord(record)
			if err != nil {
				fmt.Printf("There's error %s", err)
			} else {
				fmt.Printf("You added %s", rec.Address)
				break
			}
		} else {
			fmt.Printf("Pick waht you need to change:\n1. Ip address\n2. Hostname\n3. Disabled status\n4. End\n")
			fmt.Scanf("%d\n", &t)
			menu_handler(record, t)
		}
	}
}
