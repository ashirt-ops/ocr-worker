package textextractor

// TextExtractor is an interface for extracting text from image files with various backends
type TextExtractor interface {
	// ExtractText identifies and returns all text found within an image file
	ExtractText(img []byte) (string, error)
}
