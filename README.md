# gRPC Go Microservices Workspace

This repository demonstrates a hybrid Go microservices architecture using both gRPC and REST, with a focus on helping developers transition from REST to gRPC.

## Structure

- **content-service/**: Pure gRPC microservice in Go, backed by MongoDB. Exposes content APIs via gRPC.
- **sitemap-service/**: REST + gRPC service in Go. Exposes a REST endpoint for sitemap generation, fetching data from content-service via gRPC.
- **protos/**: Shared Protocol Buffer definitions for service contracts.

## Quick Start

1. **Install dependencies:**
   - Go (>= 1.18)
   - Protocol Buffers (`protoc`)
   - Go plugins for protobuf/gRPC

2. **Generate gRPC code:**
   - In both `content-service` and `sitemap-service`, run:
     ```sh
     make proto
     ```

3. **Set up environment:**
   - For `content-service`, create a `.env` file with your MongoDB connection string.

4. **Run services:**
   - Start `content-service` (gRPC server):
     ```sh
     cd content-service
     go run ./src/main.go
     ```
   - Start `sitemap-service` (REST + gRPC client):
     ```sh
     cd sitemap-service
     go run ./src/main.go
     ```

5. **Access the REST API:**
   - Visit [http://localhost:8080/sitemap.xml](http://localhost:8080/sitemap.xml) to see the generated sitemap (data fetched via gRPC).

## References
- [gRPC-Go Documentation](https://grpc.io/docs/languages/go/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)

---

For a step-by-step guide and deeper explanation, see the individual service READMEs and the included article in `Article.md`.
