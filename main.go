package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/ogp-cache/service"
)

func main() {
	r := gin.Default()
	r.Use(cros)
	r.GET("/v1", func(c *gin.Context) {
		url := c.Query("url")
		ogp, _ := service.PurseOgp(url)

		//計測
		start, ok := c.Get("start_time")
		if ok {
			c.Header("X-Server-Latency", time.Since(start.(time.Time)).String())
		}
		c.JSON(http.StatusOK, gin.H{
			"ogp": ogp,
		})
	})
	r.Run()
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
