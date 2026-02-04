package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/ashirt-ops/ocr-worker/internal/handlers"
	"github.com/ashirt-ops/ocr-worker/internal/textextractor"
	"github.com/jrozner/weby/middleware"
	"github.com/jrozner/weby/rlog"
)

func main() {
	var handler slog.Handler = slog.NewTextHandler(os.Stdout, nil)

	handler = rlog.RequestIDHandler{Handler: handler}
	logger := slog.New(handler)

	extractor := textextractor.NewTesseract()

	env := handlers.New(extractor)
	mux := env.Routes()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.WrapResponse)
	mux.Use(middleware.Logger(logger))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
