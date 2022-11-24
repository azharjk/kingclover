package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GreetHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, everyone\n")
}
