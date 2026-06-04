// Why: handlers translate HTTP requests into service calls and service results into HTTP responses.
// What to do: keep shared handler dependencies and request helpers here, not domain rules.
package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/config"
	"github.com/nattakornwarisnarathorn/example-gin/internal/middleware"
	"github.com/nattakornwarisnarathorn/example-gin/internal/service"
	"github.com/nattakornwarisnarathorn/example-gin/pkg/response"
)

type Dependencies struct {
	Config         *config.Config
	AuthService    *service.AuthService
	UserService    *service.UserService
	ProjectService *service.ProjectService
}

type Handler struct {
	cfg      *config.Config
	auth     *service.AuthService
	users    *service.UserService
	projects *service.ProjectService
}

func NewHandler(deps Dependencies) *Handler {
	return &Handler{
		cfg:      deps.Config,
		auth:     deps.AuthService,
		users:    deps.UserService,
		projects: deps.ProjectService,
	}
}

func userIDFromContext(c *gin.Context) (uint, bool) {
	value, exists := c.Get(middleware.UserIDKey)
	if !exists {
		response.Error(c, http.StatusUnauthorized, "missing user context")
		return 0, false
	}

	userID, ok := value.(uint)
	if !ok {
		response.Error(c, http.StatusUnauthorized, "invalid user context")
		return 0, false
	}

	return userID, true
}

func parseIDParam(c *gin.Context, name string) (uint, bool) {
	id, err := strconv.ParseUint(c.Param(name), 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid "+name)
		return 0, false
	}
	return uint(id), true
}

func writeServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrConflict):
		response.Error(c, http.StatusConflict, err.Error())
	case errors.Is(err, service.ErrInvalidLogin):
		response.Error(c, http.StatusUnauthorized, err.Error())
	case errors.Is(err, service.ErrNotFound):
		response.Error(c, http.StatusNotFound, err.Error())
	default:
		response.Error(c, http.StatusInternalServerError, "internal server error")
	}
}
