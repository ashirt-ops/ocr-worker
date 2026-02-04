FROM golang:1.25-alpine AS build

RUN mkdir app
COPY . ./app/
WORKDIR /go/app
RUN go build -v ./...

FROM alpine:latest

RUN apk add --no-cache tesseract-ocr && \
    adduser -h /home/ashirt -S -D ashirt

USER ashirt
WORKDIR /home/ashirt

COPY --from=build /go/app/ocr-worker /home/ashirt/ocr-worker

CMD ["ocr-worker"]