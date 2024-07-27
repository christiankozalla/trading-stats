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
			"id": "ki0qjspe2a2iws5",
			"created": "2024-07-25 06:48:41.102Z",
			"updated": "2024-07-25 06:48:41.102Z",
			"name": "raw_trades_count",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "kjfoj6yy",
					"name": "userId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "qy8wv1o5",
					"name": "accountId",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "zml6ffhmc1kewy5",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "kztkpyt0",
					"name": "trade_count",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": false
					}
				}
			],
			"indexes": [],
			"listRule": "@request.auth.id != \"\" && @request.auth.id = userId && @request.query.accountId = accountId",
			"viewRule": "@request.auth.id != \"\" && @request.auth.id = userId && @request.query.accountId = accountId",
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT\n    id,\n    user as userId,\n    account as accountId,\n    COUNT(id) AS trade_count\nFROM \n    raw_trades"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ki0qjspe2a2iws5")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
