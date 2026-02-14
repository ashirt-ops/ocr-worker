package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/ashirt-ops/ocr-worker/internal/config"
	"github.com/ashirt-ops/ocr-worker/internal/handlers"
	"github.com/ashirt-ops/ocr-worker/internal/textextractor"
	"github.com/jrozner/weby/middleware"
	"github.com/jrozner/weby/rlog"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var handler slog.Handler = slog.NewTextHandler(os.Stdout, nil)

	handler = rlog.RequestIDHandler{Handler: handler}
	logger := slog.New(handler)
	slog.SetDefault(logger)

	var conf config.Config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatalf("unable to read config: %v", err)
	}

	extractor, err := initializeExtractor(&conf)
	if err != nil {
		log.Fatalf("unable to initialize extractor: %v", err)
	}

	env := handlers.New(extractor, conf.APIBase, conf.AccessKey, conf.SecretKey)
	mux := env.Routes()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.WrapResponse)
	mux.Use(middleware.Logger(logger))

	host := fmt.Sprintf(":%d", conf.Port)
	log.Fatal(http.ListenAndServe(host, mux))
}

func initializeExtractor(conf *config.Config) (textextractor.TextExtractor, error) {
	switch conf.Backend {
	case "tesseract":
		return textextractor.NewTesseract(), nil
	case "gcp":
		return textextractor.NewGCP(), nil
	case "aws":
		return textextractor.NewAWS(), nil
	default:
		return nil, fmt.Errorf("unsupported backend: %s", conf.Backend)
	}
}
