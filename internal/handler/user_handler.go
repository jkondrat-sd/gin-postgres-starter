// Why: user handlers expose endpoints for data about the authenticated user.
// What to do: read identity from middleware context and delegate user lookup to UserService.
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/pkg/response"
)

func (h *Handler) Me(c *gin.Context) {
	userID, ok := userIDFromContext(c)
	if !ok {
		return
	}

	res, err := h.users.Me(userID)
	if err != nil {
		writeServiceError(c, err)
		return
	}

	response.OK(c, res)
}
