package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)
    

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        handle := c.Query("handle")
        apiKey := os.Getenv("YOUTUBE_API_KEY")

        if apiKey == "" {
            c.JSON(500, gin.H{"error": "API key is not set"})
            return
        }

        if handle == "" {
            c.JSON(400, gin.H{"error": "Handle is required"})
            return
        }


        _, _ = youtube.NewService(c, option.WithAPIKey(apiKey));

        c.String(200, "Hello, Gin! Handle: %s", handle)
    })
    r.Run(":8080") // Starts the server
}