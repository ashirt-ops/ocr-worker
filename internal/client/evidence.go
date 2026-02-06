package client

import (
	"fmt"
	"io"
)

func (c *Client) GetEvidenceContent(operationSlug, evidenceUUID string) ([]byte, error) {
	response, err := c.Get(fmt.Sprintf("%s/api/operations/%s/evidence/%s/media", c.base, operationSlug, evidenceUUID))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
