package cmd

import (
	"github.com/spf13/cobra"

	cmd_user "github.com/oxodao/datalake/cmd/user"
)

var RootCmd = &cobra.Command{
	Use:   "datalake",
	Short: "Datalake is a tool to aggregate your personal data from everywhere",
	Long:  `Datalake is a tool to aggregate your personal data from everywhere. Build statistics and visualize your life online.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(versionCmd)

	cmd_user.RegisterUserCommand(RootCmd)
}

func Register(r *cobra.Command) {
	RootCmd.AddCommand(r)
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
