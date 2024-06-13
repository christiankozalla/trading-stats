package eventhandlers

import (
	"image"
	_ "image/jpeg"
	_ "image/png"

	"github.com/pocketbase/pocketbase/core"
)

func AddImageMeta(e *core.RecordCreateEvent) error {
	f, _, err := e.HttpContext.Request().FormFile("image")
	if err != nil {
		return err
	}
	defer f.Close()

	imgConfig, _, err := image.DecodeConfig(f)

	if err != nil {
		return err
	}

	e.Record.Set("imageWidth", imgConfig.Width)
	e.Record.Set("imageHeight", imgConfig.Height)

	return nil
}

func UpdateImageMeta(e *core.RecordUpdateEvent) error {
	f, _, err := e.HttpContext.Request().FormFile("image")
	if err != nil {
		return err
	}
	defer f.Close()

	imgConfig, _, err := image.DecodeConfig(f)

	if err != nil {
		return err
	}

	e.Record.Set("imageWidth", imgConfig.Width)
	e.Record.Set("imageHeight", imgConfig.Height)

	return nil
}
