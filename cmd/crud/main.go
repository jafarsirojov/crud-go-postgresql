package main

// package
// import
// var + type
// method + function

import (
	"crud/cmd/crud/app"
	"flag"
	"net"
	"net/http"
	"path/filepath"
)

var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.String("port", "9999", "Server port")
)

func main() {
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	start(addr)
}

func start(addr string) {
	router := http.NewServeMux()
	server := app.NewServer(router, filepath.Join("web", "templates"), filepath.Join("web", "assets"))
	server.InitRoutes()

	// server'ы должны работать "вечно"
	panic(http.ListenAndServe(addr, server)) // поднимает сервер на определённом адресе и обрабатывает запросы
}
