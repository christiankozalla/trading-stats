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

		collection, err := dao.FindCollectionByNameOrId("1q29zijc1iib3z8")
		if err != nil {
			return err
		}

		// add
		new_Logfile := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "jq93zmtm",
			"name": "Logfile",
			"type": "relation",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "bwlrqzy67wf2auv",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_Logfile); err != nil {
			return err
		}
		collection.Schema.AddField(new_Logfile)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("1q29zijc1iib3z8")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("jq93zmtm")

		return dao.SaveCollection(collection)
	})
}
