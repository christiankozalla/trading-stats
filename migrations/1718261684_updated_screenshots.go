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
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("il1hgvyunyj2rse")
		if err != nil {
			return err
		}

		// add
		new_image_height := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "bia4eiuu",
			"name": "imageHeight",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), new_image_height); err != nil {
			return err
		}
		collection.Schema.AddField(new_image_height)

		// add
		new_image_width := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "znqukfa3",
			"name": "imageWidth",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), new_image_width); err != nil {
			return err
		}
		collection.Schema.AddField(new_image_width)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("il1hgvyunyj2rse")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("bia4eiuu")

		// remove
		collection.Schema.RemoveField("znqukfa3")

		return dao.SaveCollection(collection)
	})
}
