package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/campushq-official/campushq-api/src/internal/api/middlewares"
	"github.com/campushq-official/campushq-api/src/internal/api/routers"
	"github.com/campushq-official/campushq-api/src/internal/common"
	"github.com/campushq-official/campushq-api/src/internal/common/tracerr"
	"github.com/campushq-official/campushq-api/src/internal/config"
	services "github.com/campushq-official/campushq-api/src/internal/core/domain/services/auth0-services"
	repositories "github.com/campushq-official/campushq-api/src/internal/core/infrastructure/postgres/repositories/sqlc"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/mattes/migrate/source/file"
)

func main() {

	logger := common.NewLogger(nil)

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

	repositories.NewStore(connPool)
	auth0Service := services.NewAuth0Service(env)

	/*
	   |--------------------------------------------------------------------------
	   | Routers
	   |--------------------------------------------------------------------------
	   | Register all routers here.
	   |--------------------------------------------------------------------------

	*/

	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	auth := middlewares.NewAuth0Middleware(env, logger)

	r.Use(middlewares.CORS)
	r.Use(middlewares.Loggin)
	r.Use(auth.Auth0TokenValidation)

	routers.StudentRouter(r, logger, auth0Service)
	routers.DepartmentRouter(r, logger)

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
