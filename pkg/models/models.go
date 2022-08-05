package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Config struct {
	Port     string `env:"PORT" envDefault:"8080"`
	URL      string `env:"URL" envDefault:"http://localhost"`
	Redis    string `env:"REDIS" envDefault:"redis://localhost:6379"`
	Postgres string `env:"POSTGRES" envDefault:"postgres://postgres:postgres@localhost:5432/postgres"`

	PgUser     string `env:"PG_USER" envDefault:"postgres"`
	PgPassword string `env:"PG_PASSWORD" envDefault:"postgres"`
	PgHost     string `env:"PG_HOST" envDefault:"localhost"`
	PgPort     string `env:"PG_PORT" envDefault:"5432"`
	PgDatabase string `env:"PG_DB" envDefault:"postgres"`
}

func (c Config) PostgresURI() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", c.PgUser, c.PgPassword, c.PgHost, c.PgPort, c.PgDatabase)
}

type Film struct {
	LastUpdate      time.Time
	Rating          string
	Title           string
	Description     string
	SpecialFeatures string
	FullText        string
	FilmID          int
	Length          int
	ReplacementCost float64
	RentalDuration  int
	LanguageID      int
	ReleaseYear     int
	RentalRate      float64
}

func (f Film) IsEmpty() bool {
	return f.FilmID == 0
}

func (f Film) String() string {
	jsonBytes, err := json.Marshal(f)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
