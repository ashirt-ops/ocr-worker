package textextractor

import (
	"github.com/otiai10/gosseract/v2"
)

type Tesseract struct{}

// NewTesseract returns a new Tesseract TextExtractor
func NewTesseract() TextExtractor {
	return &Tesseract{}
}

func (t *Tesseract) ExtractText(img []byte) (string, error) {
	client := gosseract.NewClient()
	defer func() {
		// TODO: we're doing this because golangci-lint complains about an unchecked error
		// We can either add an exclusion to it's config or used a named return to surface'
		// this up
		_ = client.Close()
	}()

	err := client.SetImageFromBytes(img)
	if err != nil {
		return "", err
	}

	return client.Text()
}
