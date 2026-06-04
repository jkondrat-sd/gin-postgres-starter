// Why: routes centralize HTTP path groups and middleware wiring away from handler logic.
// What to do: register new route group files here when a feature grows new endpoints.
package routes // Declares this package as the API routing layer.

import ( // Starts the list of packages needed to build the router.
	"github.com/gin-gonic/gin"                                           // Provides the Gin engine, route groups, and middleware support.
	"github.com/nattakornwarisnarathorn/example-gin/internal/config"     // Provides typed config values such as the JWT secret.
	"github.com/nattakornwarisnarathorn/example-gin/internal/handler"    // Provides handler methods that routes connect to URLs.
	"github.com/nattakornwarisnarathorn/example-gin/internal/middleware" // Provides shared middleware such as CORS and JWT auth.
) // Ends the import list.

func SetupRouter(h *handler.Handler, cfg *config.Config) *gin.Engine { // Builds and returns the configured Gin router.
	r := gin.New()                                         // Creates a clean Gin engine without default middleware.
	r.Use(gin.Logger(), gin.Recovery(), middleware.CORS()) // Adds request logging, panic recovery, and browser CORS rules.

	RegisterHealthRoutes(r, h) // Registers the public health check route.

	api := r.Group("/api/v1")  // Groups versioned API routes under /api/v1.
	RegisterAuthRoutes(api, h) // Registers public auth routes such as login and register.

	protected := api.Group("/")                      // Creates a subgroup for routes that require authentication.
	protected.Use(middleware.JWTAuth(cfg.JWTSecret)) // Requires a valid JWT before protected handlers run.

	RegisterUserRoutes(protected, h)    // Registers authenticated current-user routes.
	RegisterProjectRoutes(protected, h) // Registers authenticated project routes.

	return r // Gives the fully configured router back to main.go.
} // Ends router setup.
