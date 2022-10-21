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
	var default_message = fmt.Sprintf("Your dns record is:\n ID: %s\tIp address: %s\tHostname: %s\tDisabled: %s\n", record.Id, record.Address, record.Name, record.Disabled)
	var picker_message = fmt.Sprintf("Pick waht you need to change:\n1. Hostname\n2. IP Address\n3. End\n")
	fmt.Println(status)
	switch s := strconv.Itoa(status); s {
	case "88":
		fmt.Printf(default_message)
		t := 88
		fmt.Printf(picker_message)
		fmt.Scanf("%d\n", &t)
		menu_handler(record, t)
	case "1":
		fmt.Printf(default_message)
		fmt.Printf("You changing Hostname.\n")
		fmt.Scan(&record.Name)
		menu_handler(record, 88)
	case "2":
		fmt.Printf(default_message)
		fmt.Printf("You changing Ip address.\n")
		fmt.Scan(&record.Address)
		menu_handler(record, 88)
	case "3":
		t := 3
		fmt.Printf(default_message)
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
			fmt.Printf(picker_message)
			fmt.Scanf("%d\n", &t)
			menu_handler(record, t)
		}
	}
}
