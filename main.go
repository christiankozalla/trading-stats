package main

import (
	"log"
	"os"

	"trading-statistics/eventhandlers"
	_ "trading-statistics/migrations"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	var indexFallback bool
	app.RootCmd.PersistentFlags().BoolVar(
		&indexFallback,
		"indexFallback",
		true,
		"fallback the request to index.html on missing static path (eg. when pretty urls are used with SPA)",
	)

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		Automigrate: os.Getenv("AUTOMIGRATE") == "true",
	})

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), indexFallback))
		return nil
	})

	app.OnRecordAfterCreateRequest("trade_log_files").Add(eventhandlers.CreateTradeRecordsFromLogFiles(app))

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
