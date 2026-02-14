package textextractor

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

// AWS implements a Textextractor using the rekognition service from AWS. This was
// originally tested with the textract service but results were inconsistent with
// screenshots.
type AWS struct{}

// NewAWS returns a new AWS TextExtractor
func NewAWS() *AWS {
	return &AWS{}
}

func (a *AWS) ExtractText(ctx context.Context, img []byte) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}

	client := rekognition.NewFromConfig(cfg)

	params := rekognition.DetectTextInput{
		Image: &types.Image{
			Bytes: img,
		},
	}

	resp, err := client.DetectText(ctx, &params)
	if err != nil {
		return "", err
	}

	text := strings.Builder{}

	for _, detection := range resp.TextDetections {
		// there seems to be duplication of text (each word within a line plus
		// the full line) when using this service. We'll filter on just lines
		// of text and move to only words if we end up missing content
		if detection.Type == types.TextTypesLine && detection.DetectedText != nil {
			text.WriteString(*detection.DetectedText)
		}
	}

	return text.String(), nil
}
