# Sitemap Service

This service provides a REST API endpoint (`/sitemap.xml`) that generates a sitemap by fetching content data from a gRPC content service.

## Features
- REST API using Gin (serves `/sitemap.xml`)
- gRPC client to fetch content from a content service
- Environment variable support via `.env` file (using `godotenv`)
- Sitemap URLs are generated using the `title` field of each content item

## Getting Started

### Prerequisites
- Go 1.18+
- A running gRPC content service (see `CONTENT_SERVICE_ADDR`)

### Installation
1. Clone this repository.
2. Install dependencies:
   ```zsh
   go mod tidy
   ```
3. (Optional) Create a `.env` file in the root directory:
   ```env
   CONTENT_SERVICE_ADDR=localhost:50051
   ```

### Running the Service
```zsh
go run src/main.go
```
The service will start on `http://localhost:8080` by default.

### REST Endpoint
- `GET /sitemap.xml` — Returns an XML sitemap with URLs based on content titles.

### Environment Variables
- `CONTENT_SERVICE_ADDR` — Address of the gRPC content service (default: `:50051`)

## Development
- Protobuf files are in `src/protos/`.
- To regenerate Go code from `.proto` files:
  ```zsh
  make proto
  ```

## License
MIT
