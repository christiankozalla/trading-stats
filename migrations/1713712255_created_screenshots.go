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
			"id": "il1hgvyunyj2rse",
			"created": "2024-04-21 15:10:55.003Z",
			"updated": "2024-04-21 15:10:55.003Z",
			"name": "screenshots",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "bsn1sajc",
					"name": "image",
					"type": "file",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"mimeTypes": [
							"image/png",
							"image/jpeg"
						],
						"thumbs": [
							"100x100", "0x40"
						],
						"maxSelect": 1,
						"maxSize": 2097152,
						"protected": false
					}
				},
				{
					"system": false,
					"id": "d99getyj",
					"name": "date",
					"type": "date",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "978g8nyy",
					"name": "comment",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "3dlovchr",
					"name": "account",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "zml6ffhmc1kewy5",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				}
			],
			"indexes": [],
			"listRule": "@request.auth.id != \"\" && @request.auth.id = account.user.id",
			"viewRule": "@request.auth.id != \"\" && @request.auth.id = account.user.id",
			"createRule": "@request.auth.id != \"\"",
			"updateRule": "@request.auth.id != \"\" && @request.auth.id = account.user.id",
			"deleteRule": "@request.auth.id != \"\" && @request.auth.id = account.user.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("il1hgvyunyj2rse")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
