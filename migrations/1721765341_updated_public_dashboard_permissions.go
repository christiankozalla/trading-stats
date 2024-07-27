package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pn3cdlwnxlnrx87")
		if err != nil {
			return err
		}

		// update
		edit_is_trades_table_public := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "yw3yp4op",
			"name": "is_trades_table_public",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), edit_is_trades_table_public); err != nil {
			return err
		}
		collection.Schema.AddField(edit_is_trades_table_public)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pn3cdlwnxlnrx87")
		if err != nil {
			return err
		}

		// update
		edit_is_trades_table_public := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "yw3yp4op",
			"name": "is_trades_table_public",
			"type": "bool",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), edit_is_trades_table_public); err != nil {
			return err
		}
		collection.Schema.AddField(edit_is_trades_table_public)

		return dao.SaveCollection(collection)
	})
}
