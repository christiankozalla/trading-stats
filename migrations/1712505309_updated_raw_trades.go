package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ux1jly2k6sai33h")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE UNIQUE INDEX ` + "`" + `idx_trades_InternalOrderID` + "`" + ` ON ` + "`" + `raw_trades` + "`" + ` (\n  ` + "`" + `InternalOrderID` + "`" + `,\n  ` + "`" + `account` + "`" + `\n)",
			"CREATE INDEX ` + "`" + `idx_trades_ParentInternalOrderID` + "`" + ` ON ` + "`" + `raw_trades` + "`" + ` (\n  ` + "`" + `ParentInternalOrderID` + "`" + `,\n  ` + "`" + `account` + "`" + `\n)",
			"CREATE INDEX ` + "`" + `idx_trades_OpenClose` + "`" + ` ON ` + "`" + `raw_trades` + "`" + ` (` + "`" + `OpenClose` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("ux1jly2k6sai33h")
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(`[
			"CREATE UNIQUE INDEX ` + "`" + `idx_trades_InternalOrderID` + "`" + ` ON ` + "`" + `raw_trades` + "`" + ` (` + "`" + `InternalOrderID` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_trades_ParentInternalOrderID` + "`" + ` ON ` + "`" + `raw_trades` + "`" + ` (` + "`" + `ParentInternalOrderID` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_trades_OpenClose` + "`" + ` ON ` + "`" + `raw_trades` + "`" + ` (` + "`" + `OpenClose` + "`" + `)"
		]`), &collection.Indexes); err != nil {
			return err
		}

		return dao.SaveCollection(collection)
	})
}
