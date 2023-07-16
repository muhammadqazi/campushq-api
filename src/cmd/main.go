package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/campushq-official/campushq-api/src/internal/api/middlewares"
	"github.com/campushq-official/campushq-api/src/internal/api/routers"
	validators "github.com/campushq-official/campushq-api/src/internal/api/validators/students-validator"
	"github.com/campushq-official/campushq-api/src/internal/common/logs"
	"github.com/campushq-official/campushq-api/src/internal/common/tracerr"
	"github.com/campushq-official/campushq-api/src/internal/config"
	auth0Service "github.com/campushq-official/campushq-api/src/internal/core/domain/services/auth0-services"
	services "github.com/campushq-official/campushq-api/src/internal/core/domain/services/student-services"

	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/mattes/migrate/source/file"
)

func main() {

	logger := logs.NewLogger(nil)

	/*
	   |--------------------------------------------------------------------------
	   | Viper Config
	   |--------------------------------------------------------------------------
	   | Load the environment variables from .env file
	   |--------------------------------------------------------------------------

	*/

	env, err := config.LoadConfig()
	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	}

	/*
	   |--------------------------------------------------------------------------
	   | Database Connection Pool & Migration
	   |--------------------------------------------------------------------------
	   | Establish connection & Migrate the database
	   |--------------------------------------------------------------------------

	*/

	connPool, err := pgxpool.New(context.Background(), env.DB_URL)
	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	}

	migration, err := migrate.New(env.MIGRATION_URL, env.DB_URL)
	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
	}

	/*
	   |--------------------------------------------------------------------------
	   | Initialize Services
	   |--------------------------------------------------------------------------
	   | Initialize all services here.
	   |--------------------------------------------------------------------------

	*/

	repository := repositories.NewStore(connPool)

	var (
		auth0Service   = auth0Service.NewAuth0Service(env)
		studentService = services.NewStudentService(repository)
	)

	/*
	   |--------------------------------------------------------------------------
	   | Mux initialization & Middlewares
	   |--------------------------------------------------------------------------
	   | Initialize mux router & middlewares
	   |--------------------------------------------------------------------------

	*/

	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	auth := middlewares.NewAuth0Middleware(env, logger)

	r.Use(middlewares.CORS)
	r.Use(middlewares.Loggin)
	r.Use(auth.Auth0Authentication)

	/*
	   |--------------------------------------------------------------------------
	   | Initialize the validators
	   |--------------------------------------------------------------------------
	   | Initialize all validators here.
	   |--------------------------------------------------------------------------

	*/

	var (
		studentValidator = validators.NewStudentValidators(repository)
	)

	/*
	   |--------------------------------------------------------------------------
	   | Router initialization
	   |--------------------------------------------------------------------------
	   | Initialize and register all routers here.
	   |--------------------------------------------------------------------------

	*/
	var (
		studentRouter    = routers.NewStudentRouter(r, logger, auth0Service, studentService, studentValidator)
		departmentRouter = routers.NewDepartmentRouter(r, logger)
	)

	studentRouter.StudentRouter()
	departmentRouter.DepartmentRouter()

	/*
	   |--------------------------------------------------------------------------
	   | Server
	   |--------------------------------------------------------------------------
	   | Start the server
	   |--------------------------------------------------------------------------
	*/
	s := http.Server{
		Addr:         env.PORT,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	logger.Info("Server started at port", env.PORT)

	err = s.ListenAndServe()
	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
		os.Exit(1)
	}

}
