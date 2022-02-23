package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/oxodao/datalake/cmd"
	"github.com/oxodao/datalake/module_spotify"
	"github.com/oxodao/datalake/services"
	"github.com/spf13/cobra"
)

func main() {
	err := services.NewProvider()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	services.SetProviders([]services.DataProvider{
		module_spotify.New(),
	})

	for _, dp := range services.Get().DataProviders {
		customCmd := &cobra.Command{
			Use:   dp.GetName(),
			Short: strings.Title(dp.GetName()) + " module for Datalake",
		}

		dp.RegisterCustomCommands(customCmd)

		cmd.RootCmd.AddCommand(customCmd)
	}

	cmd.Execute()
}
