package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		taList := []models.Record{}
		err := dao.DB().NewQuery("SELECT ta.* FROM trading_accounts ta WHERE NOT EXISTS (SELECT 1 FROM public_dashboard_permissions pdp WHERE pdp.account = ta.id)").All(&taList)
		if err != nil {
			return err
		}

		pdpCollection, err := dao.FindCollectionByNameOrId("public_dashboard_permissions")
		if err != nil {
			return err
		}

		for _, record := range taList {
			newRecord := models.NewRecord(pdpCollection)
			newRecord.Set("account", record.Get("id"))
			newRecord.Set("is_trades_table_public", false)

			if err := dao.Save(newRecord); err != nil {
				return err
			}
		}

		return nil
	}, func(db dbx.Builder) error {
		// add down queries...

		return nil
	})
}
