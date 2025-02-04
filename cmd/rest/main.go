package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/daverussell13/to-do-list-bts/internal/domain"
	"github.com/daverussell13/to-do-list-bts/internal/envvar"
	"github.com/daverussell13/to-do-list-bts/internal/postgresql"
	"github.com/daverussell13/to-do-list-bts/internal/rest"
	"github.com/daverussell13/to-do-list-bts/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	if err := envvar.Load(); err != nil {
		log.Fatalf("error loading .env file")
	}

	conf := envvar.New(nil)
	pgPool, err := newPostgreSQL(conf)
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	userRepo := postgresql.NewUser(pgPool)
	authSvc := service.NewAuth(userRepo)

	authHandler := rest.NewAuthHandler(authSvc)

	e.POST("/login", authHandler.Login)

	e.Logger.Fatal(e.Start(":1323"))
}

func newPostgreSQL(conf *envvar.Configuration) (*pgxpool.Pool, error) {
	get := func(v string) string {
		res, err := conf.Get(v)
		if err != nil {
			log.Fatalf("couldn't get configuration value for %s: %s", v, err)
		}

		return res
	}

	databaseHost := get("DATABASE_HOST")
	databasePort := get("DATABASE_PORT")
	databaseUsername := get("DATABASE_USERNAME")
	databasePassword := get("DATABASE_PASSWORD")
	databaseName := get("DATABASE_NAME")
	databaseSSLMode := get("DATABASE_SSLMODE")

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(databaseUsername, databasePassword),
		Host:   fmt.Sprintf("%s:%s", databaseHost, databasePort),
		Path:   databaseName,
	}

	q := dsn.Query()
	q.Add("sslmode", databaseSSLMode)

	dsn.RawQuery = q.Encode()

	pool, err := pgxpool.New(context.Background(), dsn.String())
	if err != nil {
		return nil, domain.WrapErrorf(err, domain.ErrorCodeUnknown, "pgxpool.Connect")
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, domain.WrapErrorf(err, domain.ErrorCodeUnknown, "db.Ping")
	}

	return pool, nil
}
