package config

import "strings"

type Web struct {
	Url           string `yaml:"url"`
	ListeningAddr string `yaml:"listening_addr"`
	Port          string `yaml:"port"`
}

func (w *Web) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type innerWeb Web

	if err := unmarshal((*innerWeb)(w)); err != nil {
		return err
	}

	w.Url = strings.TrimSuffix(w.Url, "/")

	return nil
}
