package module_spotify

import (
	"fmt"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import data from Spotify RGPD zip",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Import")
	},
}
