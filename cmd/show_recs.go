package cmd

import (
	"fmt"
	"mikrotik_helper/utils"
	"os"

	"github.com/spf13/cobra"
)

var showRecsCmd = &cobra.Command{
	Use:     "get_records",
	Aliases: []string{"get"},
	Short:   "Get all dns static records",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := utils.Dial()
		if err != nil {
			fmt.Println("there's err on connection", err)
			os.Exit(1)
		}
		res := utils.GetAllDnsRecords(client)
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(showRecsCmd)
}
