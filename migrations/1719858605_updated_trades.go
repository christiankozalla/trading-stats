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

		collection, err := dao.FindCollectionByNameOrId("1q29zijc1iib3z8")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = user.id")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = user.id")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = user.id")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = user.id")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("1q29zijc1iib3z8")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = user.id && @request.query.accountId = account.id")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = user.id && @request.query.accountId = account.id")

		collection.UpdateRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = user.id && @request.query.accountId = account.id")

		collection.DeleteRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = user.id && @request.query.accountId = account.id")

		return dao.SaveCollection(collection)
	})
}
