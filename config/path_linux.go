//go:build linux

package config

import (
	"errors"
	"os"
	"strings"
)

func FindConfigPath() (string, error) {
	searchedPath := []string{}

	if _, err := os.Stat("datalake.yml"); !os.IsNotExist(err) {
		return "datalake.yml", nil
	}
	searchedPath = append(searchedPath, "./datalake.yml")

	home, err := os.UserHomeDir()
	if err == nil {
		if _, err := os.Stat(home + "/.config/datalake/datalake.yml"); !os.IsNotExist(err) {
			return home + "/.config/datalake/datalake.yml", nil
		}

		searchedPath = append(searchedPath, home+"/.config/datalake/datalake.yml")
	} else {
		searchedPath = append(searchedPath, "Could not access home dir: "+err.Error())
	}

	if _, err := os.Stat("/etc/datalake/datalake.yml"); !os.IsNotExist(err) {
		return "/etc/datalake/datalake.yml", nil
	}

	searchedPath = append(searchedPath, "/etc/datalake/datalake.yml")

	return "", errors.New("Could not find config file, searched in \n\t- " + strings.Join(searchedPath, "\n\t- "))
}
