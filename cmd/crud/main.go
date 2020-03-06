package main

import (
	"context"
	"crud/cmd/crud/app"
	"crud/pkg/crud/services/burgers"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.String("port", "9999", "Server port")
	dsn  = flag.String("dsn", "postgres://app:pass@localhost:5432/app", "Postgres DSN")
)

func main() {
	flag.Parse()
	Port, ok := os.LookupEnv("PORT")
	if !ok{
		Port = *port
	}
	Dsn, ok := os.LookupEnv("DATABASE_URL")
	if !ok{
		Dsn = *dsn
	}
	Host, ok := os.LookupEnv("HOST")
	if !ok{
		Host = *host
	}

	addr := net.JoinHostPort(Host,Port)
	log.Println("starting server!")
	log.Printf("host = %s, port = %s\n",Host,Port)
	start(addr, Dsn)
}

func start(addr string, dsn string) {
	log.Println("started server!")
	router := app.NewExactMux()
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	burgersSvc := burgers.NewBurgersSvc(pool)
	server := app.NewServer(
		router,
		pool,
		burgersSvc,
		filepath.Join("web", "templates"),
		filepath.Join("web", "assets"),
	)
	server.InitRoutes()
	panic(http.ListenAndServe(addr, server))
}
