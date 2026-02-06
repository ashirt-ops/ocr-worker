FROM golang:1.25-alpine AS build

RUN mkdir app && \
    apk add --no-cache build-base tesseract-ocr-dev

COPY . ./app/
WORKDIR /go/app

RUN go build -v ./cmd/...

FROM alpine:latest

RUN apk add --no-cache tesseract-ocr tesseract-ocr-data-eng && \
    adduser -h /home/ashirt -S -D ashirt

USER ashirt
WORKDIR /home/ashirt

COPY --from=build /go/app/ocr-worker /home/ashirt/ocr-worker

CMD ["ocr-worker"]