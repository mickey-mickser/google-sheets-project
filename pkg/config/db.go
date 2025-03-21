package config

import (
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"github.com/mickey-mickser/telegram-project/pkg/repos"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var allowedDrivers = map[string]struct{}{
	"postgres": {},
}

type db struct {
	dbParams `json:"db"`
	db       *gorm.DB
}

type dbParams struct {
	URL    string `json:"url"`
	Driver string `json:"driver"`
}

func (d *db) populate(cfgFile []byte) error {
	if err := json.Unmarshal(cfgFile, d); err != nil {
		return err
	}

	return nil
}

func (d *db) validate() error {
	return errors.Wrap(d.check(), "failed to validate db")
}

func (d *db) check() error {
	if _, err := pq.ParseURL(d.URL); err != nil {
		return err
	}
	if _, ok := allowedDrivers[d.Driver]; !ok {
		return fmt.Errorf("not allowed driver %s (postgres)", d.Driver)
	}

	return nil
}

func (d *db) DB() *gorm.DB {
	if d.db == nil {
		var err error
		d.db, err = repos.NewDB(d.URL)
		if err != nil {
			panic(err)
		}
	}

	return d.db
}
