package app

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func (receiver *server) handleBurgersList() func(writer http.ResponseWriter, request *http.Request) {
	// TODO: handle concurrency
	burgers := make([]Burger, 0)
	var nextId int64 = 0
	// handler + closure
	// TODO: make some initialization -> only once
	// glob -> * - всё, кроме /
	// glob -> ? - один символ, но не /
	tpl, err := template.ParseFiles(filepath.Join(receiver.templatePath, "index.gohtml"))
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			err := request.ParseForm()
			if err != nil {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
				return
			}

			action := request.PostForm.Get("action")
			if action == "" {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
				return
			}

			//switch action {
			//case "save":
			//	...
			//case "remove":
			//}

			log.Print(request)
			log.Print(request.URL.Query())
			log.Print(request.Form)
			log.Print(request.PostForm)
			log.Print(request.MultipartForm)

			name, ok := request.PostForm["name"]
			if !ok {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
				return
			}

			price, ok := request.PostForm["price"]
			if !ok {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
				return
			}

			description, ok := request.PostForm["description"]
			if !ok {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
				return
			}

			// TODO: всё, что приходит по HTTP - string
			parsedPrice, err := strconv.Atoi(price[0])
			if err != nil {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400 - Bad Request
				return
			}

			// TODO: общее соглашение
			// - id = 0, создание нового элемента
			// - id != 0, обновление существующего элемента
			nextId++
			burger := Burger{nextId, name[0], parsedPrice, description[0]}
			burgers = append(burgers, burger)
		}

		// TODO: fetch from DB
		data := &struct {
			Title string
			Burgers []Burger
		}{
			Title: "McDonalds",
			Burgers: burgers,
		}

		err = tpl.Execute(writer, data)
		if err != nil {
			log.Print(err)
		}
	}
}

func (receiver *server) handleFavicon() func(http.ResponseWriter, *http.Request) {
	// TODO: handle concurrency
	file, err := ioutil.ReadFile(filepath.Join(receiver.assetsPath, "favicon.ico"))
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write(file)
		if err != nil {
			log.Print(err)
		}
	}
}
