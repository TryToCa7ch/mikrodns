package cmd

import (
	"fmt"
	"log"
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
			log.Println(color_print.Fata("Provide static record id"))
			os.Exit(1)
		}
		client, err := utils.Dial()
		if err != nil {
			log.Println(color_print.Fata("there's err on connection", err))
			os.Exit(1)
		}
		record, err := utils.GetDnsRecord(client, args[0])
		menu_status := 88
		menu_handler(record, menu_status)
		if err != nil {
			log.Println(color_print.Fata("There's error when trying to change record: \n", err))
		}
	},
}

func init() {
	log.SetFlags(0)

	rootCmd.AddCommand(changeRecsCmd)
}

func menu_handler(record utils.DnsRecord, status int) {
	var default_message = fmt.Sprintf("Your dns record is:\n ID: %s\tIp address: %s\tHostname: %s\tDisabled: %s\n", record.Id, record.Address, record.Name, record.Disabled)
	var picker_message = "Pick waht you need to change:\n1. Hostname\n2. IP Address\n3. End\n"

	switch s := strconv.Itoa(status); s {
	case "88":
		log.Println(default_message)
		t := 88
		log.Println(picker_message)
		fmt.Scanf("%d\n", &t)
		menu_handler(record, t)
	case "1":
		log.Println(default_message)
		log.Printf("You changing Hostname.\n")
		fmt.Scan(&record.Name)
		menu_handler(record, 88)
	case "2":
		log.Println(default_message)
		log.Printf("You changing Ip address.\n")
		fmt.Scan(&record.Address)
		menu_handler(record, 88)
	case "3":
		t := 3
		log.Println(default_message)
		log.Printf("Confirm changing[Y/n]\n")
		var confirmation string
		fmt.Scanf("%s\n", &confirmation)
		if confirmation == "Y" || confirmation == "y" {
			rec, err := utils.ChangeDnsRecord(record)
			if err != nil {
				log.Printf("There's error %s", err)
			} else {
				log.Printf("You added %s", rec.Address)
				break
			}
		} else {
			log.Println(picker_message)
			fmt.Scanf("%d\n", &t)
			menu_handler(record, t)
		}
	}
}
