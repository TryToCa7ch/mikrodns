package cmd

import (
	"log"
	"mikrodns/color_print"
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
			log.Println(color_print.Fata("there's err on connection: ", err))
			os.Exit(1)
		}
		res, err := utils.GetAllDnsRecords(client)
		if err == nil {
			for _, rec := range res {
				log.Printf("ID: %s\tAddress: %s\tIP: %s\tDisabled: %s\n", rec.Id[1:], rec.Name, rec.Address, rec.Disabled)
			}
		} else {
			log.Printf(color_print.Fata("There's errors: %s", err))
			os.Exit(1)
		}
	},
}

var showRecCmd = &cobra.Command{
	Use:     "get_record",
	Aliases: []string{"get"},
	Args:    cobra.MaximumNArgs(1),
	Short:   "Get  dns static record with given id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println(color_print.Fata("Provide static record id"))
			os.Exit(1)
		}
		client, err := utils.Dial()
		if err != nil {
			log.Println(color_print.Fata("there's err on connection: ", err))
			os.Exit(1)
		}
		res, err := utils.GetDnsRecord(client, args[0])
		if err != nil {
			log.Println(color_print.Fata("There's no dns record with given id"))
			os.Exit(1)
		}
		log.Printf("ID: %s\tAddress: %s\tIP: %s\tDisabled: %s\n", res.Id[1:], res.Name, res.Address, res.Disabled)
	},
}

func init() {
	log.SetFlags(0)

	rootCmd.AddCommand(showRecsCmd)
	rootCmd.AddCommand(showRecCmd)
}
