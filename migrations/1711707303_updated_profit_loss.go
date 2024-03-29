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

		collection, err := dao.FindCollectionByNameOrId("pteexcmbctdt6u8")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id = user.id && @request.query.accountId = account.id")

		collection.ViewRule = types.Pointer("@request.auth.id = user.id && @request.query.accountId = account.id")

		// remove
		collection.Schema.RemoveField("ys07mrz8")

		// remove
		collection.Schema.RemoveField("pnookmyu")

		// remove
		collection.Schema.RemoveField("wl1texyc")

		// remove
		collection.Schema.RemoveField("8z2myogt")

		// add
		new_user := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "y94n97lb",
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
			"id": "y6878adm",
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
			"id": "2kxmdhri",
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
			"id": "hu4370rp",
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

		collection.ListRule = types.Pointer("")

		collection.ViewRule = types.Pointer("")

		// add
		del_user := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ys07mrz8",
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
			"id": "pnookmyu",
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
			"id": "wl1texyc",
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
			"id": "8z2myogt",
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
		collection.Schema.RemoveField("y94n97lb")

		// remove
		collection.Schema.RemoveField("y6878adm")

		// remove
		collection.Schema.RemoveField("2kxmdhri")

		// remove
		collection.Schema.RemoveField("hu4370rp")

		return dao.SaveCollection(collection)
	})
}
