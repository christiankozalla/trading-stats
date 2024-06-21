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
			"id": "pn3cdlwnxlnrx87",
			"created": "2024-06-19 07:18:43.274Z",
			"updated": "2024-06-19 07:18:43.274Z",
			"name": "public_dashboard_permissions",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "qyo3cfbv",
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
					"id": "yw3yp4op",
					"name": "is_trades_table_public",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				}
			],
			"indexes": [],
			"listRule": "",
			"viewRule": "",
			"createRule": "@request.auth.id != \"\" && @request.query.accountId != \"\"",
			"updateRule": "@request.auth.id != \"\" && @request.auth.id = account.user.id && @request.query.accountId = account.id",
			"deleteRule": "@request.auth.id != \"\" && @request.auth.id = account.user.id && @request.query.accountId = account.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pn3cdlwnxlnrx87")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
