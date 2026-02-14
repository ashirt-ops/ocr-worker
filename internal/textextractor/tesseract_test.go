package textextractor

import (
	"context"
	"os"
	"testing"
)

func TestTesseract(t *testing.T) {
	data, err := os.ReadFile("./testdata/ashirt.png")
	if err != nil {
		t.Fatalf("unable to read test data: %v", err)
	}
	extractor := NewTesseract()
	extracted, err := extractor.ExtractText(context.Background(), data)
	if err != nil {
		t.Fatalf("unable to extract text: %v", err)
	}

	if extracted != "ASHIRT" {
		t.Fatalf("extracted text does not match expected text, expected: %v, got: %v", "ashirt", extracted)
	}
}
