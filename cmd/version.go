package cmd

import (
	"fmt"

	"github.com/oxodao/datalake/utils"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Datalake version " + utils.CURRENT_VERSION + " (Commit " + utils.CURRENT_COMMIT + ")")
	},
}
