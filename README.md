# ocr-worker

## Configuration

Configuration is managed through environment variables. Below are the recognized configuration options:

| env var    | description                    | valid values             | default | required |
|------------|--------------------------------|--------------------------| ------- | -------- |
| API_BASE   | URL for ASHIRT API server      |                          | | yes |
| ACCESS_KEY | ASHIRT access key              |                          | | yes |
| SECRET_KEY | ASHIRT secret key              |                          | | yes |
| BACKEND    | text extraction backend to use | tesseract                | tesseract | no |
| LOG_LEVEL  | logging level                  | debug, info, warn, error | info | no |
| PORT       | the tcp port to bind on        | any valid port           | 8080 | no |
