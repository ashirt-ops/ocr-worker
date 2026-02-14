package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ashirt-ops/ocr-worker/internal/client"
	"github.com/ashirt-ops/ocr-worker/internal/messages"
	"github.com/ashirt-ops/ocr-worker/internal/textextractor"
	"github.com/jrozner/weby"
)

type Env struct {
	extractor textextractor.TextExtractor
	client    *client.Client
}

func New(extractor textextractor.TextExtractor, base, accessKey string, secretKey []byte) *Env {
	return &Env{
		extractor: extractor,
		client:    client.New(base, accessKey, secretKey),
	}
}

func (e *Env) Routes() *weby.ServeMux {
	mux := weby.NewServeMux()
	mux.HandleFunc("POST /process", e.Process)

	return mux
}

func (e *Env) Process(w http.ResponseWriter, r *http.Request) {
	var (
		request  messages.Request
		response messages.Response
		status   int
		data     []byte
		content  string
	)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.ErrorContext(r.Context(), "error decoding request body", "error", err)
		goto end
	}

	if request.Type == "test" {
		testResponse := messages.Test{Status: "ok"}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(testResponse)
		if err != nil {
			slog.ErrorContext(r.Context(), "error encoding request body", "error", err)
			return
		}
		return
	}

	if request.ContentType != "image" {
		slog.DebugContext(r.Context(), "unsupported content type, skipping processing", "content_type", request.ContentType)
		response.Action = "rejected"
		status = http.StatusNotAcceptable
		goto end
	}

	data, err = e.client.GetEvidenceContent(request.OperationSlug, request.EvidenceUUID)
	if err != nil {
		slog.ErrorContext(r.Context(), "error getting evidence content", "error", err)
		status = http.StatusInternalServerError
		response.Action = "error"
		goto end
	}

	content, err = e.extractor.ExtractText(r.Context(), data)
	if err != nil {
		slog.ErrorContext(r.Context(), "error extracting text", "error", err)
		status = http.StatusInternalServerError
		response.Action = "error"
		goto end
	}

	status = http.StatusOK
	response.Action = "processed"
	response.Content = content

end:
	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		slog.ErrorContext(r.Context(), "unable to serialize response", "error", err)
	}
}
