package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// add up migration
		dao := daos.New(db)

		settings, err := dao.FindSettings()
		if err != nil {
			return err
		}
		settings.Meta.AppName = "inloopo trading stats"
		settings.Logs.MaxDays = 2

		return dao.SaveSettings(settings)
	}, nil)
}
