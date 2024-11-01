package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"io/ioutil"
	"net/http"
)

type Image struct {
	Filename    string
	ContentType string
	Data        []byte
	Size        int
}

func okContentType(contentType string) bool {
	return contentType == "image/png" || contentType == "image/jpeg" || contentType == "image/gif"
}

// Process uploaded file into an image
func Process(r *http.Request, field string) (*Image, error) {
	file, info, err := r.FormFile(field)

	if err != nil {
		return nil, err
	}

	contentType := info.Header.Get("Content-Type")

	if !okContentType(contentType) {
		return nil, errors.New(fmt.Sprintf("Wrong content type: %s", contentType))
	}

	bs, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	_, _, err = image.Decode(bytes.NewReader(bs))

	if err != nil {
		return nil, err
	}

	i := &Image{
		Filename:    info.Filename,
		ContentType: contentType,
		Data:        bs,
		Size:        len(bs),
	}

	return i, nil
}

// Create PNG thumbnail from image
func (i *Image) ThumbnailPNG(width int, height int) (*Image, error) {
	return ThumbnailPNG(i, width, height)
}

// Create PNG thumbnail
func ThumbnailPNG(i *Image, width int, height int) (*Image, error) {
	img, _, err := image.Decode(bytes.NewReader(i.Data))

	thumbnail := resize.Thumbnail(uint(width), uint(height), img, resize.Lanczos3)

	data := new(bytes.Buffer)
	err = png.Encode(data, thumbnail)

	if err != nil {
		return nil, err
	}

	bs := data.Bytes()

	t := &Image{
		Filename:    "thumbnail.png",
		ContentType: "image/png",
		Data:        bs,
		Size:        len(bs),
	}

	return t, nil
}
