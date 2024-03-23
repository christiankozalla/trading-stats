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

		collection, err := dao.FindCollectionByNameOrId("pteexcmbctdt6u8")
		if err != nil {
			return err
		}

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"query": "SELECT\n  id,\n  user,\n  account,\n  DateTime_close,\n  (ProfitLoss * Multiplier) AS ProfitLoss_dollar\nFROM trades;"
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("m8ce08n4")

		// remove
		collection.Schema.RemoveField("loif3qm3")

		// remove
		collection.Schema.RemoveField("mnybsyzo")

		// remove
		collection.Schema.RemoveField("ybmmmjkb")

		// add
		new_user := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "lpsgj0uv",
			"name": "user",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_user); err != nil {
			return err
		}
		collection.Schema.AddField(new_user)

		// add
		new_account := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "fqyllcaj",
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
		}`), new_account); err != nil {
			return err
		}
		collection.Schema.AddField(new_account)

		// add
		new_DateTime_close := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "6r5mppig",
			"name": "DateTime_close",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_DateTime_close); err != nil {
			return err
		}
		collection.Schema.AddField(new_DateTime_close)

		// add
		new_ProfitLoss_dollar := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "4igmndob",
			"name": "ProfitLoss_dollar",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), new_ProfitLoss_dollar); err != nil {
			return err
		}
		collection.Schema.AddField(new_ProfitLoss_dollar)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pteexcmbctdt6u8")
		if err != nil {
			return err
		}

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"query": "SELECT\n  id,\n  user,\n  account,\n  DateTime_close,\n  (ProfitLoss * Multiplier) AS ProfitLoss_dollar\nFROM trades\nGROUP BY trades.` + "`" + `DateTime_close` + "`" + `"
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		// add
		del_user := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "m8ce08n4",
			"name": "user",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), del_user); err != nil {
			return err
		}
		collection.Schema.AddField(del_user)

		// add
		del_account := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "loif3qm3",
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
		}`), del_account); err != nil {
			return err
		}
		collection.Schema.AddField(del_account)

		// add
		del_DateTime_close := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "mnybsyzo",
			"name": "DateTime_close",
			"type": "date",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), del_DateTime_close); err != nil {
			return err
		}
		collection.Schema.AddField(del_DateTime_close)

		// add
		del_ProfitLoss_dollar := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ybmmmjkb",
			"name": "ProfitLoss_dollar",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), del_ProfitLoss_dollar); err != nil {
			return err
		}
		collection.Schema.AddField(del_ProfitLoss_dollar)

		// remove
		collection.Schema.RemoveField("lpsgj0uv")

		// remove
		collection.Schema.RemoveField("fqyllcaj")

		// remove
		collection.Schema.RemoveField("6r5mppig")

		// remove
		collection.Schema.RemoveField("4igmndob")

		return dao.SaveCollection(collection)
	})
}
