package eventhandlers

import (
	"encoding/csv"
	"io"
	"log"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/hook"
)

func CreateTradeRecordsFromLogFiles(app *pocketbase.PocketBase) hook.Handler[*core.RecordCreateEvent] {
	return func(e *core.RecordCreateEvent) error {
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
					if strings.Contains(globexCode, ".") {
						parts := strings.Split(globexCode, ".")
						globexCode = parts[0][:len(parts[0])-2]
					} else if len(globexCode) >= 2 {
						globexCode = globexCode[:len(globexCode)-2]
					}

					multiplier := globexCodeToMultiplierMap[globexCode]
					trade["Multiplier"] = multiplier
					trade["ShortSymbol"] = globexCode
				}

				trade["user"] = userId
				trade["account"] = accountId

				err = createTradeRecord(app, trade)
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
}

func createTradeRecord(app *pocketbase.PocketBase, data map[string]any) error {
	tradesCollection, err := app.Dao().FindCollectionByNameOrId("trades")
	if err != nil {
		return err
	}

	record := models.NewRecord(tradesCollection)

	// WITH data validation - which works, but only creates empty records
	form := forms.NewRecordUpsert(app, record)

	form.LoadData(data)

	if err := form.Submit(); err != nil {
		return err
	}

	return nil
}
