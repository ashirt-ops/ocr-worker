package messages

type Request struct {
	ContentType   string `json:"contentType"`
	EvidenceUUID  string `json:"evidenceUuid"`
	OperationSlug string `json:"operationSlug"`
	Type          string `json:"type"`
}
