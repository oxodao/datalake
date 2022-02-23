package module_spotify

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/oxodao/datalake/services"
	"github.com/spf13/cobra"
)

/** @TODO read directly the zip and import other data too **/

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import data from Spotify GDPR zip",
	Run: func(cmd *cobra.Command, args []string) {

		//#region Ugly, find a proper way to do this
		found := false
		var spot DataProvider
		for _, d := range services.Get().DataProviders {
			if d.GetName() == "spotify" {
				spot = d.(DataProvider)
				found = true
				break
			}
		}

		if !found {
			fmt.Println("No spotify provider")
			os.Exit(1)
		}
		//#endregion

		filePath, err := cmd.Flags().GetString("file")
		if err != nil || len(filePath) == 0 {
			fmt.Println("Please specify the file path")
			os.Exit(1)
		}

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Println("File not found")
			os.Exit(1)
		}

		fmt.Println("Importing...")
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		arr := []SpotifyStream{}
		err = json.Unmarshal(data, &arr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, elt := range arr {
			idArtist, err := spot.ORM.UpsertArtist(elt.ArtistName)
			if err != nil {
				fmt.Println(err)
				continue
			}

			idTrack, err := spot.ORM.UpsertTrack(idArtist, elt.TrackName)
			if err != nil {
				fmt.Println(err)
				continue
			}

			_, err = spot.ORM.UpsertPlayedTrack(idTrack, elt.EndTime, elt.DurationPlayed)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	},
}

func registerImportCmd(cmd *cobra.Command) {
	importCmd.Flags().StringP("file", "f", "", "The file you get from Spotify when you ask for your data")
	cmd.AddCommand(importCmd)
}
