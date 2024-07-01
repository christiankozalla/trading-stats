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

		collection.ListRule = types.Pointer("account.public_dashboard_permissions_via_account.is_trades_table_public = true || @request.auth.id = user.id")

		collection.ViewRule = types.Pointer("account.public_dashboard_permissions_via_account.is_trades_table_public = true || @request.auth.id = user.id")

		// remove
		collection.Schema.RemoveField("i9dawps2")

		// remove
		collection.Schema.RemoveField("wlym6h9w")

		// remove
		collection.Schema.RemoveField("yr6chann")

		// remove
		collection.Schema.RemoveField("ssxboqff")

		// add
		new_user := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "rpesp17t",
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
			"id": "uwegbhub",
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
			"id": "f3pip0wj",
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
			"id": "q4xdiwlq",
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

		collection.ListRule = types.Pointer("@request.query.accountId = account.id &&  (account.public_dashboard_permissions_via_account.is_trades_table_public = true || @request.auth.id = user.id)")

		collection.ViewRule = types.Pointer("@request.query.accountId = account.id &&  (account.public_dashboard_permissions_via_account.is_trades_table_public = true || @request.auth.id = user.id)")

		// add
		del_user := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "i9dawps2",
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
			"id": "wlym6h9w",
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
			"id": "yr6chann",
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
			"id": "ssxboqff",
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
		collection.Schema.RemoveField("rpesp17t")

		// remove
		collection.Schema.RemoveField("uwegbhub")

		// remove
		collection.Schema.RemoveField("f3pip0wj")

		// remove
		collection.Schema.RemoveField("q4xdiwlq")

		return dao.SaveCollection(collection)
	})
}
