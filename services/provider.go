package services

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/datalake/config"
	"github.com/oxodao/datalake/orm"
	"github.com/spf13/cobra"
)

var provider *Provider = nil

type Provider struct {
	Config        *config.Config
	DB            *sqlx.DB
	ORM           *orm.ORM
	Password      *Password
	DataProviders []DataProvider
}

func NewProvider() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	db, err := sqlx.Open("postgres", cfg.Database.GetDSN())
	if err != nil {
		return err
	}

	provider = &Provider{
		Config:        cfg,
		DB:            db,
		ORM:           orm.New(db),
		Password:      &Password{},
		DataProviders: []DataProvider{},
	}

	return nil
}

func SetProviders(providers []DataProvider) {
	provider.DataProviders = providers
}

func Get() *Provider {
	if provider == nil {
		provider = &Provider{}
	}
	return provider
}

type DataProvider interface {
	IsEnabled() bool
	GetName() string
	Process()
	RegisterRoutes(r *mux.Router)
	RegisterCustomCommands(cmd *cobra.Command)
}
