package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ki0qjspe2a2iws5")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = userId && @request.query.logfileId = id")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = userId && @request.query.logfileId = id")

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"query": "SELECT\n    id,\n    user as userId,\n    COUNT(id) AS trade_count\nFROM \n    raw_trades"
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		// remove
		collection.Schema.RemoveField("kjfoj6yy")

		// remove
		collection.Schema.RemoveField("qy8wv1o5")

		// remove
		collection.Schema.RemoveField("kztkpyt0")

		// add
		new_userId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qswxgsb6",
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
		}`), new_userId); err != nil {
			return err
		}
		collection.Schema.AddField(new_userId)

		// add
		new_trade_count := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ylx8kjk0",
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
		}`), new_trade_count); err != nil {
			return err
		}
		collection.Schema.AddField(new_trade_count)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ki0qjspe2a2iws5")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = userId && @request.query.accountId = accountId")

		collection.ViewRule = types.Pointer("@request.auth.id != \"\" && @request.auth.id = userId && @request.query.accountId = accountId")

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"query": "SELECT\n    id,\n    user as userId,\n    account as accountId,\n    COUNT(id) AS trade_count\nFROM \n    raw_trades"
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		// add
		del_userId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), del_userId); err != nil {
			return err
		}
		collection.Schema.AddField(del_userId)

		// add
		del_accountId := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), del_accountId); err != nil {
			return err
		}
		collection.Schema.AddField(del_accountId)

		// add
		del_trade_count := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
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
		}`), del_trade_count); err != nil {
			return err
		}
		collection.Schema.AddField(del_trade_count)

		// remove
		collection.Schema.RemoveField("qswxgsb6")

		// remove
		collection.Schema.RemoveField("ylx8kjk0")

		return dao.SaveCollection(collection)
	})
}
