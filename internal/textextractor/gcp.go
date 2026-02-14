package textextractor

import (
	"context"
	"fmt"

	vision "cloud.google.com/go/vision/v2/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
)

type GCP struct{}

// NewGCP returns a new GCP TextExtractor
func NewGCP() *GCP {
	return &GCP{}
}

func (g *GCP) ExtractText(ctx context.Context, img []byte) (string, error) {
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return "", err
	}
	//nolint:errcheck
	defer client.Close()

	req := visionpb.BatchAnnotateImagesRequest{
		Requests: []*visionpb.AnnotateImageRequest{
			&visionpb.AnnotateImageRequest{
				Image: &visionpb.Image{
					Content: img,
				},
				Features: []*visionpb.Feature{
					&visionpb.Feature{
						Type: visionpb.Feature_TEXT_DETECTION,
					},
				},
			},
		},
	}

	resp, err := client.BatchAnnotateImages(ctx, &req)
	if err != nil {
		return "", err
	}

	if len(resp.Responses) != 1 {
		return "", fmt.Errorf("expected 1 response, got %d", len(resp.Responses))
	}

	return resp.Responses[0].FullTextAnnotation.Text, nil
}
