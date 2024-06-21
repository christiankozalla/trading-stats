package eventhandlers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
	"unicode"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/hook"
)

func CreateTradeRecordsFromLogFiles(app *pocketbase.PocketBase) hook.Handler[*core.RecordCreateEvent] {
	return func(e *core.RecordCreateEvent) error {
		err := createRawTradeRecords(app, e)
		if err != nil {
			log.Fatal(err)
			return err
		}

		err = createTradeRecords(app, e)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return nil
	}
}

func createRawTradeRecords(app *pocketbase.PocketBase, e *core.RecordCreateEvent) error {
	files := e.UploadedFiles["file"]
	userId := e.Record.GetString("user")
	accountId := e.Record.GetString("account")

	defer app.Dao().SaveRecord(e.Record) // setting the status after operation

	for _, f := range files {
		reader, err := f.Reader.Open()
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer reader.Close()

		csvReader := csv.NewReader(reader)
		csvReader.Comma = '\t'

		header, err := csvReader.Read()

		if err != nil {
			log.Fatal(err)
			return err
		}

		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
				e.Record.Set("status", "error")
				return err
			}

			trade := make(map[string]any)
			for i, value := range record {
				trade[header[i]] = value
			}

			// adds multiplier to calculate cash ($) from ticks
			// because ProfitLoss is given in ticks, each tick has different value for each future
			// hence the multiplier
			globexCode, ok := trade["Symbol"].(string)
			if ok {
				// todo: figure out whether the Symbol refers to a Future, Stock or Option
				var extractedGlobexCode string
				if parts := strings.Split(globexCode, "."); len(parts) > 1 {
					extractedGlobexCode = extractGlobexCode(parts[0])
				} else if parts := strings.Split(globexCode, "_"); len(parts) > 1 {
					extractedGlobexCode = extractGlobexCode(parts[0])
				} else {
					extractedGlobexCode = extractGlobexCode(globexCode)
				}

				multiplier := globexCodeToMultiplierMap[extractedGlobexCode]
				if multiplier == 0 {
					app.Logger().Error(fmt.Sprintf("Matching Globex Code Not Found for Symbol '%s'", globexCode))
					multiplier = 1
				}
				trade["Multiplier"] = multiplier
				trade["ShortSymbol"] = extractedGlobexCode
			}

			trade["user"] = userId
			trade["account"] = accountId

			err = createRecord(app, "raw_trades", trade)
			if err != nil {
				log.Println(err)
				e.Record.Set("status", "error")
				return err
			}
		}
	}

	e.Record.Set("status", "processed")

	return nil
}

func createRecord(app *pocketbase.PocketBase, nameOrId string, data map[string]any) error {
	collection, err := app.Dao().FindCollectionByNameOrId(nameOrId)
	if err != nil {
		return err
	}

	record := models.NewRecord(collection)

	// WITH data validation - which works, but only creates empty records
	form := forms.NewRecordUpsert(app, record)

	form.LoadData(data)

	if err := form.Submit(); err != nil {
		return err
	}

	return nil
}

type Trade struct {
	ID             string  `db:"id" json:"id"`
	DateTimeOpen   string  `db:"DateTime_open" json:"DateTime_open"`
	DateTimeClose  string  `db:"DateTime_close" json:"DateTime_close"`
	User           string  `db:"user" json:"user"`
	Account        string  `db:"account" json:"account"`
	Symbol         string  `db:"Symbol" json:"Symbol"`
	ShortSymbol    string  `db:"ShortSymbol" json:"ShortSymbol"`
	Multiplier     float64 `db:"Multiplier" json:"Multiplier"`
	QuantityOpen   float64 `db:"Quantity_open" json:"Quantity_open"`
	BuySellOpen    string  `db:"BuySell_open" json:"BuySell_open"`
	FillPriceOpen  float64 `db:"FillPrice_open" json:"FillPrice_open"`
	FillPriceClose float64 `db:"FillPrice_close" json:"FillPrice_close"`
	BuySellClose   string  `db:"BuySell_close" json:"BuySell_close"`
	QuantityClose  float64 `db:"Quantity_close" json:"Quantity_close"`
	ProfitLoss     float64 `db:"ProfitLoss" json:"ProfitLoss"`
}

func createTradeRecords(app *pocketbase.PocketBase, e *core.RecordCreateEvent) error {
	accountId := e.Record.GetString("account")

	trades := []Trade{}

	err := app.Dao().DB().NewQuery(`
		SELECT 
			CONCAT(openTrades.id, '_', closeTrades.id) AS id,
			openTrades.DateTime AS DateTime_open,
			closeTrades.DateTime AS DateTime_close,
			openTrades.user AS user,
			openTrades.account AS account,
			openTrades.Symbol AS Symbol,
			openTrades.ShortSymbol AS ShortSymbol,
			openTrades.Multiplier AS Multiplier,
			openTrades.Quantity AS Quantity_open,
			openTrades.BuySell AS BuySell_open,
			openTrades.FillPrice AS FillPrice_open,
			closeTrades.FillPrice AS FillPrice_close,
			closeTrades.BuySell AS BuySell_close,
			closeTrades.Quantity AS Quantity_close,
			(closeTrades.FillPrice - openTrades.FillPrice) AS ProfitLoss
		FROM raw_trades AS openTrades
		INNER JOIN raw_trades AS closeTrades ON openTrades.InternalOrderID = closeTrades.ParentInternalOrderID
		WHERE
			openTrades.account = {:accountId}
		AND
			closeTrades.account = {:accountId}
		AND
			openTrades.OpenClose = 'Open'
		AND
			closeTrades.OpenClose = 'Close';
	`).Bind(dbx.Params{
		"accountId": accountId,
	}).All(&trades)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, trade := range trades {
		// this conversion should be possible with PocketBase funcs reading the db struct tag
		tradeMap := tradeStructToMap(trade)

		err = createRecord(app, "trades", tradeMap)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

func tradeStructToMap(trade Trade) map[string]any {
	return map[string]any{
		"ID":              trade.ID,
		"DateTime_open":   trade.DateTimeOpen,
		"DateTime_close":  trade.DateTimeClose,
		"user":            trade.User,
		"account":         trade.Account,
		"Symbol":          trade.Symbol,
		"ShortSymbol":     trade.ShortSymbol,
		"Multiplier":      trade.Multiplier,
		"Quantity_open":   trade.QuantityOpen,
		"BuySell_open":    trade.BuySellOpen,
		"FillPrice_open":  trade.FillPriceOpen,
		"FillPrice_close": trade.FillPriceClose,
		"BuySell_close":   trade.BuySellClose,
		"Quantity_close":  trade.QuantityClose,
		"ProfitLoss":      trade.ProfitLoss,
	}
}

func extractGlobexCode(s string) string {
	for i, char := range s {
		if unicode.IsDigit(char) {
			if i > 0 {
				return s[0 : i-1]
			}
			break
		}
	}
	return s
}
