package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mikrodns",
	Short: `mikrodns is a CLI client of dns of given mikrotik.`,
	Long: `mikrodns is a CLI client ofdns of given mikrotik. You can view, change and add dns static record. 
	 You need to add MIKROTIK_HOST, MIKROTIK_USER, MIKROTIK_PASS and MIKROTIK_TLS
	 to environment variables or to ~/.mikrodns.yml`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Whoops. There was an error while executing your CLI ", err)
	}
}
