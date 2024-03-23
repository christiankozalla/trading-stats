package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "pteexcmbctdt6u8",
			"created": "2024-03-22 18:18:38.923Z",
			"updated": "2024-03-22 18:18:38.923Z",
			"name": "profit_loss",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "kq4yvzxl",
					"name": "account",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "zml6ffhmc1kewy5",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "wltcosli",
					"name": "ProfitLoss",
					"type": "json",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSize": 2000000
					}
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT\n  id,\n  account,\n  ProfitLoss\nFROM trades\nGROUP BY trades.` + "`" + `DateTime_close` + "`" + `"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pteexcmbctdt6u8")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
