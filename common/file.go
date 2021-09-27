package common

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func DecodeBase64toFile(data string) (string, error) {
	name := uuid.New().String()
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", xerrors.New("bad base64")
	}
	ImageType := data[11:idx]

	unbased, err := base64.StdEncoding.DecodeString(data[idx+8:])
	if err != nil {
		return "", xerrors.New("bad base64")
	}
	r := bytes.NewReader(unbased)
	switch ImageType {
	case "png":
		filename := name + ".png"
		im, err := png.Decode(r)
		if err != nil {
			return "", xerrors.New("bad png")
		}

		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return "", xerrors.New("cant open png")
		}
		err = png.Encode(f, im)
		if err != nil {
			return "", xerrors.New("cant encode")
		}
		return filename, nil
	case "jpeg":
		filename := name + ".jpeg"
		im, err := png.Decode(r)
		if err != nil {
			return "", xerrors.New("bad jpeg")
		}

		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return "", xerrors.New("cant open jpeg")
		}
		err = jpeg.Encode(f, im, nil)
		if err != nil {
			return "", xerrors.New("cant encode")
		}
		return filename, nil

	}
	return "", xerrors.New("bad type image")
}
