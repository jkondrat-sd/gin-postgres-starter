package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/dto"
	"github.com/nattakornwarisnarathorn/example-gin/pkg/response"
)

func (h *Handler) Health(c *gin.Context) {
	response.OK(c, dto.HealthResponse{
		App:    h.cfg.AppName,
		Env:    h.cfg.AppEnv,
		Status: "ok",
	})
}
