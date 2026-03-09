# ocr-worker

OCR Worker is a service worker for ASHIRT that implements OCR capability for evidence. It supports multiple backends using tesseract as a fully open source option not tied to one of the cloud providers, GCP vision api for GCP deployments, and rekognition for AWS deployments.

## Configuration

Configuration is managed through environment variables. Below are the recognized configuration options:

| env var    | description                    | valid values             | default | required |
|------------|--------------------------------|--------------------------| ------- | -------- |
| API_BASE   | URL for ASHIRT API server      |                          | | yes |
| ACCESS_KEY | ASHIRT access key              |                          | | yes |
| SECRET_KEY | ASHIRT secret key              |                          | | yes |
| BACKEND    | text extraction backend to use | tesseract, gcp, aws      | tesseract | no |
| LOG_LEVEL  | logging level                  | debug, info, warn, error | info | no |
| PORT       | the tcp port to bind on        | any valid port           | 8080 | no |

## License

This project is licensed under the terms of the MIT open source license. Please refer to [LICENSE](LICENSE) for the full terms.
