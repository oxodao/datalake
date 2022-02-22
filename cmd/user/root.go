package cmd_user

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "user",
	Short: "User commands",
}

func RegisterUserCommand(cmd *cobra.Command) {
	registerCreateCmd()

	cmd.AddCommand(rootCmd)
}
