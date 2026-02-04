package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ashirt-ops/ocr-worker/internal/messages"
	"github.com/ashirt-ops/ocr-worker/internal/textextractor"
	"github.com/jrozner/weby"
)

type Env struct {
	extractor textextractor.TextExtractor
}

func New(extractor textextractor.TextExtractor) *Env {
	return &Env{
		extractor: extractor,
	}
}

func (e *Env) Routes() *weby.ServeMux {
	mux := weby.NewServeMux()
	mux.HandleFunc("/process", e.Process)

	return mux
}

func (e *Env) Process(w http.ResponseWriter, r *http.Request) {
	var request messages.Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.ErrorContext(r.Context(), "error decoding request body", "error", err)
		goto error
	}

	if request.ContentType != "IMAGE" {
		slog.DebugContext(r.Context(), "unsupported content type, skipping processing", "content_type", request.ContentType)
		return
	}

error:
	response := messages.Response{
		Action:  "rejected",
		Content: "",
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		slog.ErrorContext(r.Context(), "unable to serialize response", "error", err)
	}
}
