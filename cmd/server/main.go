// Why: this entrypoint composes config, database, repositories, services, handlers, and routes into one API process.
// What to do: keep startup wiring here, and move domain logic into internal packages.
package main // Declares this file as the executable application package.

import ( // Starts the list of packages this file needs.
	"fmt" // Formats startup errors with useful context.

	"github.com/gin-gonic/gin"                                           // Provides the HTTP server and Gin mode settings.
	"github.com/nattakornwarisnarathorn/example-gin/internal/config"     // Loads environment-based application config.
	"github.com/nattakornwarisnarathorn/example-gin/internal/handler"    // Creates HTTP handlers from services.
	"github.com/nattakornwarisnarathorn/example-gin/internal/repository" // Creates database repositories.
	"github.com/nattakornwarisnarathorn/example-gin/internal/routes"     // Registers Gin routes and middleware.
	"github.com/nattakornwarisnarathorn/example-gin/internal/service"    // Creates business-logic services.
	"github.com/nattakornwarisnarathorn/example-gin/pkg/database"        // Connects to PostgreSQL and runs local migrations.
	"github.com/nattakornwarisnarathorn/example-gin/pkg/logger"          // Initializes the shared structured logger.
	"go.uber.org/zap"                                                    // Adds structured fields to log messages.
) // Ends the import list.

func main() { // Starts the API process.
	cfg, err := config.Load() // Reads .env/environment variables into a typed Config struct.
	if err != nil {           // Checks whether config loading failed.
		panic(fmt.Errorf("load config: %w", err)) // Stops startup because the app cannot run without config.
	} // Ends the config error check.

	logger.Init(cfg.AppEnv) // Configures the logger for development or production.
	defer logger.Log.Sync() // Flushes buffered log entries before the process exits.

	if cfg.AppEnv == "production" { // Checks whether the app should run with production Gin settings.
		gin.SetMode(gin.ReleaseMode) // Disables extra Gin debug output in production.
	} // Ends the Gin mode check.

	db, err := database.ConnectPostgres(cfg) // Opens the PostgreSQL connection using config values.
	if err != nil {                          // Checks whether the database connection failed.
		logger.Log.Fatal("database connection failed", zap.Error(err)) // Logs the database error and stops the process.
	} // Ends the database connection error check.

	if err := database.AutoMigrate(db); err != nil { // Runs GORM AutoMigrate for local starter development.
		logger.Log.Fatal("database migration failed", zap.Error(err)) // Logs migration failure and stops startup.
	} // Ends the migration error check.

	userRepo := repository.NewUserRepository(db)       // Creates the user database access layer.
	projectRepo := repository.NewProjectRepository(db) // Creates the project database access layer.

	authService := service.NewAuthService(userRepo, cfg)     // Creates auth business logic with users and JWT config.
	userService := service.NewUserService(userRepo)          // Creates user business logic.
	projectService := service.NewProjectService(projectRepo) // Creates project business logic.

	apiHandler := handler.NewHandler(handler.Dependencies{ // Creates HTTP handlers and injects their dependencies.
		Config:         cfg,            // Provides app metadata for handlers such as health.
		AuthService:    authService,    // Provides register/login behavior to auth handlers.
		UserService:    userService,    // Provides current-user behavior to user handlers.
		ProjectService: projectService, // Provides project CRUD behavior to project handlers.
	}) // Ends handler dependency wiring.

	router := routes.SetupRouter(apiHandler, cfg) // Builds the Gin router with routes and middleware.

	addr := ":" + cfg.AppPort                                   // Builds the listen address from the configured port.
	logger.Log.Info("server running", zap.String("addr", addr)) // Logs the address before starting the server.
	if err := router.Run(addr); err != nil {                    // Starts the HTTP server and watches for startup/runtime errors.
		logger.Log.Fatal("server stopped", zap.Error(err)) // Logs fatal server errors before exiting.
	} // Ends the server run error check.
} // Ends the API process.
