// Why: auth handlers own the HTTP details for registration and login requests.
// What to do: validate request bodies here, then delegate password and token logic to AuthService.
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/dto"
	"github.com/nattakornwarisnarathorn/example-gin/pkg/response"
)

func (h *Handler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.auth.Register(req)
	if err != nil {
		writeServiceError(c, err)
		return
	}

	response.Created(c, res)
}

func (h *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.auth.Login(req)
	if err != nil {
		writeServiceError(c, err)
		return
	}

	response.OK(c, res)
}
