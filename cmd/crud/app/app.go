package app

import (
	"errors"
	"net/http"
	"github.com/jackc/pgx/v4/pgxpoll"
)

// описание сервиса, который хранит зависимости и выполняет работу
type server struct { // <- Alt + Enter -> Constructor
	// зависимости (dependencies)
	pool 		*pgxpoll.Pool
	router       http.Handler
	templatePath string
	assetsPath   string
}

// 1. required <-
// 2. optional

// crash early
func NewServer(router http.Handler, templatePath string, assetsPath string) *server {
	if router == nil {
		panic(errors.New("router can't be nil"))
	}
	if templatePath == "" {
		panic(errors.New("templatePath can't be empty"))
	}
	if assetsPath == "" {
		panic(errors.New("assetsPath can't be empty"))
	}

	return &server{router: router, templatePath: templatePath, assetsPath: assetsPath}
}

func (receiver *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	receiver.router.ServeHTTP(writer, request)
}

