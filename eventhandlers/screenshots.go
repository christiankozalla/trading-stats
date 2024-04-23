package eventhandlers

import (
	"bytes"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/filesystem"
)

// This is a pre-processor for user-provided images
// the screenshots collection has an image column with a max-size of 2MB
// 2MB applies to the initial file provided by the user, before pre-processing
// Actually, this pre-processing might be more overhead that use!
// Alternatively reduce the max-size limit even further and skip pre-processing alltogether!
func ResizeImages(e *core.RecordCreateEvent) error {
	files, ok := e.UploadedFiles["image"]
	if !ok {
		return errors.New("no screenshot file has been sent")
	}

	file := files[0]
	src, err := file.Reader.Open()
	log.Println("reading file", file.Name, file, file.Size)

	if err != nil {
		log.Printf("Error reading image data: %v\n", err)
		return errors.New("failed to read image data")
	}
	defer src.Close()

	// Decode the image
	img, format, err := image.Decode(src)

	if err != nil {
		log.Printf("Error decoding image: %v\n", err)
		return errors.New("invalid image format")
	}

	log.Println("decoded image", format)

	// Resize the image
	resizedImg := resizeImage(img, 800)

	// Encode the resized image to webp format
	var resizedImgWebpBuf bytes.Buffer
	// opts, err := encoder.NewLosslessEncoderOptions(encoder.PresetDefault, 6)
	opts, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 90)

	if err != nil {
		log.Println("error creating webp encoding options")
		return errors.New("error creating webp encoding options")
	}

	err = webp.Encode(&resizedImgWebpBuf, resizedImg, opts)

	if err != nil {
		return err
	}
	newFileName := strings.TrimSuffix(file.Name, filepath.Ext(file.Name)) + ".webp"
	webpFile, err := filesystem.NewFileFromBytes(resizedImgWebpBuf.Bytes(), newFileName)

	log.Println("created webp file", webpFile.Name, webpFile.OriginalName, webpFile.Size)

	if err != nil {
		return err
	}

	e.UploadedFiles["image"][0] = webpFile
	e.Record.Set("image", webpFile.Name)
	file = nil

	return nil
}

func resizeImage(img image.Image, maxWidth int) image.Image {
	return imaging.Resize(img, maxWidth, 0, imaging.Lanczos) // preserving the aspect-ratio
}
