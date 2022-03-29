package web

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TestAPI(c *gin.Context) {
	buf, _ := json.Marshal("ok")
	c.Data(http.StatusOK, "application/json; charset=utf-8", buf)
	c.Abort()
}

func StartServer() {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	api := engine.Group("/api")
	api.GET("/test", TestAPI)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8090),
		Handler:        engine,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
