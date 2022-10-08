package main

import (
	// "fmt"
	"log"

	utils "mikrotik_helper/utils"

	"github.com/spf13/cobra"
	// "reflect"
	// "format"
)

var rootCmd = &cobra.Command{
	Use:   "mikrodns",
	Short: `mikrodns is a CLI client of dns of given mikrotik.`,
	Long: `mikrodns is a CLI client ofdns of given mikrotik. You can view, change and add dns static record. 
	 You need to add MIKROTIK_HOST, MIKROTIK_USER, MIKROTIK_PASS and MIKROTIK_TLS
	 to environment variables`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set command values based on preference
	},
}

func main() {
	log.Printf("starting, %v %v %v \n", utils.Address, utils.Username, utils.Password)

	client, err := utils.Dial()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// log.Print(utils.AddDnsRecord(client, "test.test", "192.168.66.187"))
	records := utils.GetAllDnsRecords(client)
	for _, rec := range records {
		log.Print(rec)
	}
}
