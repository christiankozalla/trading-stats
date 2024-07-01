package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pn3cdlwnxlnrx87")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.id != \"\"")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = account.user.id")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = account.user.id")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pn3cdlwnxlnrx87")
		if err != nil {
			return err
		}

		collection.CreateRule = types.Pointer("@request.auth.id != \"\" && @request.query.accountId != \"\"")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = account.user.id && @request.query.accountId = account.id")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = account.user.id && @request.query.accountId = account.id")

		return dao.SaveCollection(collection)
	})
}
