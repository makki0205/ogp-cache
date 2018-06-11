package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cros)
}

func cros(c *gin.Context) {
	headers := c.Request.Header.Get("Access-Control-Request-Headers")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", headers)
	if c.Request.Method == "OPTIONS" {
		c.Status(200)
		c.Abort()
	}
	c.Set("start_time", time.Now())
	c.Next()
}
