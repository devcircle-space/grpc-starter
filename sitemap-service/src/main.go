package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	contentpb "sitemap-service/src/protos"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env file, proceeding with system environment variables.")
	}
}

func main() {
	loadEnv()
	// gRPC client setup
	grpcAddr := os.Getenv("CONTENT_SERVICE_ADDR")
	if grpcAddr == "" {
		grpcAddr = ":50051" // default
	}
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to content-service: %v", err)
	}
	defer conn.Close()
	client := contentpb.NewContentServiceClient(conn)

	r := gin.Default()

	r.GET("/sitemap.xml", func(c *gin.Context) {
		resp, err := client.ListContent(context.Background(), &contentpb.ListContentRequest{})
		if err != nil {
			log.Printf("Failed to fetch content list: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch content list"})
			return
		}
		var urls []string
		for _, content := range resp.GetContents() {
			urls = append(urls, "https://example.com/content/"+content.GetXId())
		}
		// Generate simple XML sitemap
		sitemap := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n"
		for _, url := range urls {
			sitemap += "  <url><loc>" + url + "</loc></url>\n"
		}
		sitemap += "</urlset>"
		c.Data(http.StatusOK, "application/xml", []byte(sitemap))
	})

	r.Run(":8080")
}
