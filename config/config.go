package config

import (
	"github.com/ian-kent/gofigure"
)

var config *configuration

func Init() {
	config = &configuration{
		ImportAPIURL:  "http://localhost:21800"           ,
	}

	gofigure.Gofigure(config)
}

// Config values for the application.
type configuration struct {
	ImportAPIURL                 string `env:"IMPORT_API_URL" flag:"import-api-url" flagDesc:"The address of the import API"`
}

func ImportAPIURL() string {
	return config.ImportAPIURL
}
