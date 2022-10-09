package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mikrodns",
	Short: `mikrodns is a CLI client of dns of given mikrotik.`,
	Long: `mikrodns is a CLI client ofdns of given mikrotik. You can view, change and add dns static record. 
	 You need to add MIKROTIK_HOST, MIKROTIK_USER, MIKROTIK_PASS and MIKROTIK_TLS
	 to environment variables`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
