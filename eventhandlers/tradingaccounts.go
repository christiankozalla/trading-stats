package eventhandlers

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/hook"
)

func CreateDefaultPublicDashboardRecord(app *pocketbase.PocketBase) hook.Handler[*core.RecordCreateEvent] {
	return func(e *core.RecordCreateEvent) error {
		accountId := e.Record.GetString("id")

		pdpCollection, err := app.Dao().FindCollectionByNameOrId("public_dashboard_permissions")
		if err != nil {
			return err
		}

		newRecord := models.NewRecord(pdpCollection)
		form := forms.NewRecordUpsert(app, newRecord)
		form.LoadData(map[string]any{
			"account":                accountId,
			"is_trades_table_public": false, // not public by default
		})

		if err := form.Submit(); err != nil {
			return err
		}

		return nil
	}
}
