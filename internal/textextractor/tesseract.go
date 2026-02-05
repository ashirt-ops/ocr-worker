package textextractor

import (
	"github.com/otiai10/gosseract"
)

type Tesseract struct{}

// NewTesseract returns a new Tesseract TextExtractor
func NewTesseract() TextExtractor {
	return &Tesseract{}
}

func (t *Tesseract) ExtractText(img []byte) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	err := client.SetLanguage("eng")
	if err != nil {
		return "", err
	}

	err = client.SetImageFromBytes(img)
	if err != nil {
		return "", err
	}

	return client.Text()
}
