// Why: response helpers keep API success and error JSON shapes consistent across handlers.
// What to do: add new response helpers here when many handlers need the same format.
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Body{Success: true, Data: data})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Body{Success: true, Data: data})
}

func Message(c *gin.Context, status int, message string) {
	c.JSON(status, Body{Success: status < http.StatusBadRequest, Message: message})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, Body{Success: false, Error: message})
}
