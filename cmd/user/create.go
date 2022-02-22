package cmd_user

import (
	"fmt"
	"os"

	"github.com/oxodao/datalake/models"
	"github.com/oxodao/datalake/services"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		if err != nil || len(username) == 0 {
			fmt.Println("Username required")
			os.Exit(1)
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil || len(password) == 0 {
			fmt.Println("Password required")
			os.Exit(1)
		}

		p, err := services.Get().Password.Hash(password)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		u := &models.User{
			Name:     username,
			Password: p,
		}

		err = services.Get().ORM.Auth.CreateUser(u)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func registerCreateCmd() {
	createCmd.Flags().StringP("username", "u", "", "Username")
	createCmd.Flags().StringP("password", "p", "", "Password")
	rootCmd.AddCommand(createCmd)
}
