package textextractor

import (
	"context"

	"github.com/otiai10/gosseract/v2"
)

type Tesseract struct{}

// NewTesseract returns a new Tesseract TextExtractor
func NewTesseract() TextExtractor {
	return &Tesseract{}
}

func (t *Tesseract) ExtractText(ctx context.Context, img []byte) (string, error) {
	client := gosseract.NewClient()
	//nolint:errcheck
	defer client.Close()

	err := client.SetImageFromBytes(img)
	if err != nil {
		return "", err
	}

	return client.Text()
}
