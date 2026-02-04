package textextractor

type Tesseract struct{}

// NewTesseract returns a new Tesseract TextExtractor
func NewTesseract() TextExtractor {
	return &Tesseract{}
}

func (t *Tesseract) ExtractText(img []byte) (string, error) {
	return "", nil
}
