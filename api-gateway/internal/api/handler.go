package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Handler handles API requests and proxies them to the appropriate service
type Handler struct {
	client *http.Client
}

// NewHandler creates a new API handler
func NewHandler() *Handler {
	return &Handler{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// ProxyRequest forwards a request to the specified service
func (h *Handler) ProxyRequest(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceURL := viper.GetString(fmt.Sprintf("services.%s.url", serviceName))
		if serviceURL == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Service %s not configured", serviceName),
			})
			return
		}

		// Build the target URL
		targetURL := fmt.Sprintf("%s%s", serviceURL, c.Request.URL.Path)

		// Create a new request
		req, err := http.NewRequest(c.Request.Method, targetURL, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to create request: %v", err),
			})
			return
		}

		// Copy headers
		for key, values := range c.Request.Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		// Copy body if it exists
		if c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": fmt.Sprintf("Failed to read request body: %v", err),
				})
				return
			}
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Send the request
		resp, err := h.client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to send request to %s: %v", serviceName, err),
			})
			return
		}
		defer resp.Body.Close()

		// Copy response headers
		for key, values := range resp.Header {
			for _, value := range values {
				c.Header(key, value)
			}
		}

		// Copy response body
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to read response body: %v", err),
			})
			return
		}

		// Set response status and body
		c.Status(resp.StatusCode)
		c.Writer.Write(bodyBytes)
	}
}

// HealthCheck handles health check requests
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}
