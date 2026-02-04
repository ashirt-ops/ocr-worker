package messages

type Request struct {
	ContentType   string `json:"content_type"`
	EvidenceUUID  string `json:"evidence_uuid"`
	OperationSlug string `json:"operation_slug"`
}
