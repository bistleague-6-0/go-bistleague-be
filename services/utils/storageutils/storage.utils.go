package storageutils

import (
	"encoding/base64"
	"errors"
	"strings"
)

var signatures = map[string]string{
	"JVBERi0":     "application/pdf",
	"iVBORw0KGgo": "image/png",
	"/9j/":        "image/jpeg",
}

type Base64File struct {
	Name     string
	Ext      string
	Contents []byte
}

func NewBase64FromString(base64File string, filename string) (*Base64File, error) {
	data, ext, err := DecodeBase64WithFormat(base64File)
	if err != nil {
		return nil, err
	}
	return &Base64File{
		Name:     filename,
		Ext:      ext,
		Contents: data,
	}, nil
}

func NewBase64File(name string, contents []byte) *Base64File {
	return &Base64File{
		Name:     name,
		Contents: contents,
	}
}

func detectMimeType(b64 string) (string, error) {
	for s, mimeType := range signatures {
		if strings.HasPrefix(b64, s) {
			return mimeType, nil
		}
	}
	return "", errors.New("unsupported base64 format")
}

func DecodeBase64WithFormat(base64Data string) ([]byte, string, error) {
	mimeType, err := detectMimeType(base64Data)
	if err != nil {
		return nil, "", err
	}

	decodedData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, "", err
	}

	var ext string
	switch mimeType {
	case "application/pdf":
		ext = ".pdf"
	case "image/png":
		ext = ".png"
	case "image/jpeg":
		ext = ".jpg"
	default:
		return nil, "", errors.New("unsupported mimeType: " + mimeType)
	}

	return decodedData, ext, nil
}
