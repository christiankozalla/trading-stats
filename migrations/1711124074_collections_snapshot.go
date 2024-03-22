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
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2024-03-03 19:42:42.123Z",
				"updated": "2024-03-03 19:42:42.124Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
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
						"id": "users_avatar",
						"name": "avatar",
						"type": "file",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [
								"image/jpeg",
								"image/png",
								"image/svg+xml",
								"image/gif",
								"image/webp"
							],
							"thumbs": null,
							"maxSelect": 1,
							"maxSize": 5242880,
							"protected": false
						}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"onlyVerified": false,
					"requireEmail": false
				}
			},
			{
				"id": "zml6ffhmc1kewy5",
				"created": "2024-03-09 20:58:00.947Z",
				"updated": "2024-03-11 20:21:21.545Z",
				"name": "trading_accounts",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "u4i1ilzf",
						"name": "name",
						"type": "text",
						"required": true,
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
						"id": "qt7x31sl",
						"name": "user",
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
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id = user.id",
				"viewRule": "@request.auth.id = user.id",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id = user.id",
				"deleteRule": "@request.auth.id = user.id",
				"options": {}
			},
			{
				"id": "bwlrqzy67wf2auv",
				"created": "2024-03-11 20:08:08.650Z",
				"updated": "2024-03-13 13:22:56.585Z",
				"name": "trade_log_files",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "urhcwieb",
						"name": "user",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "dymkhxhs",
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
					},
					{
						"system": false,
						"id": "4ssenohb",
						"name": "file",
						"type": "file",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"mimeTypes": [],
							"thumbs": [],
							"maxSelect": 99,
							"maxSize": 5242880,
							"protected": false
						}
					},
					{
						"system": false,
						"id": "8f7grfm5",
						"name": "status",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "@request.auth.id = user.id",
				"viewRule": "@request.auth.id = user.id",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id = user.id",
				"deleteRule": "@request.auth.id = user.id",
				"options": {}
			},
			{
				"id": "ux1jly2k6sai33h",
				"created": "2024-03-12 07:28:09.858Z",
				"updated": "2024-03-21 06:52:43.665Z",
				"name": "raw_trades",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "nd3vuau1",
						"name": "user",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "cpsgpjhv",
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
					},
					{
						"system": false,
						"id": "k8avzpuo",
						"name": "ActivityType",
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
						"id": "ftew5ytq",
						"name": "DateTime",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "t3ybpl9i",
						"name": "TransDateTime",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "8e64t7tn",
						"name": "ShortSymbol",
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
						"id": "qjsyf0al",
						"name": "Symbol",
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
						"id": "2rmminxl",
						"name": "Multiplier",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "eafmjur0",
						"name": "OrderActionSource",
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
						"id": "bcybdnbr",
						"name": "InternalOrderID",
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
						"id": "d07kwn4b",
						"name": "ServiceOrderID",
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
						"id": "s5tsm28k",
						"name": "OrderType",
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
						"id": "vz6onohu",
						"name": "Quantity",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "5nkhwego",
						"name": "BuySell",
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
						"id": "wcroersb",
						"name": "FillPrice",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "dd0id8ai",
						"name": "Price",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "rzfdoqnh",
						"name": "Price2",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "rkvsyl24",
						"name": "OrderStatus",
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
						"id": "0gzxjpsc",
						"name": "FilledQuantity",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "bgnmiqqp",
						"name": "TradeAccount",
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
						"id": "cmudzpgf",
						"name": "OpenClose",
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
						"id": "2hrav7jd",
						"name": "ParentInternalOrderID",
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
						"id": "n1mhxcqz",
						"name": "PositionQuantity",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "wqcnclkj",
						"name": "FillExecutionServiceID",
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
						"id": "vx3luuvl",
						"name": "HighDuringPosition",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "1gkny0uk",
						"name": "LowDuringPosition",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "bzdnqfkh",
						"name": "Note",
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
						"id": "gfzx3zob",
						"name": "AccountBalance",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "jp1hthyz",
						"name": "ExchangeOrderID",
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
						"id": "ntuhuhjy",
						"name": "ClientOrderID",
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
						"id": "l9vrcy3a",
						"name": "TimeInForce",
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
						"id": "scarp9qx",
						"name": "Username",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_trades_InternalOrderID` + "`" + ` ON ` + "`" + `raw_trades` + "`" + ` (` + "`" + `InternalOrderID` + "`" + `)",
					"CREATE INDEX ` + "`" + `idx_trades_ParentInternalOrderID` + "`" + ` ON ` + "`" + `raw_trades` + "`" + ` (` + "`" + `ParentInternalOrderID` + "`" + `)",
					"CREATE INDEX ` + "`" + `idx_trades_OpenClose` + "`" + ` ON ` + "`" + `raw_trades` + "`" + ` (` + "`" + `OpenClose` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\" && @request.auth.id = user.id",
				"viewRule": "@request.auth.id != \"\" && @request.auth.id = user.id",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.id = user.id",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.id = user.id",
				"options": {}
			},
			{
				"id": "1q29zijc1iib3z8",
				"created": "2024-03-21 07:37:11.677Z",
				"updated": "2024-03-22 15:29:39.266Z",
				"name": "trades",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "3ch7yg9o",
						"name": "DateTime_open",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "3tjpesph",
						"name": "DateTime_close",
						"type": "date",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					},
					{
						"system": false,
						"id": "rl6z5yfl",
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
					},
					{
						"system": false,
						"id": "rrpqyids",
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
						"id": "95aaodnk",
						"name": "Symbol",
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
						"id": "0ql8zohj",
						"name": "ShortSymbol",
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
						"id": "lo445lyp",
						"name": "Multiplier",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "eiwlzswh",
						"name": "Quantity_open",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "r44tevss",
						"name": "Quantity_close",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "bzsejl6y",
						"name": "FillPrice_open",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "4ci1pj1j",
						"name": "FillPrice_close",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "cndw7ozf",
						"name": "BuySell_open",
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
						"id": "sywejqig",
						"name": "BuySell_close",
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
						"id": "p5ijo6wm",
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
				"listRule": "@request.auth.id != \"\" && @request.auth.id = user.id",
				"viewRule": "@request.auth.id != \"\" && @request.auth.id = user.id",
				"createRule": "@request.auth.id != \"\"",
				"updateRule": "@request.auth.id != \"\" && @request.auth.id = user.id",
				"deleteRule": "@request.auth.id != \"\" && @request.auth.id = user.id",
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
