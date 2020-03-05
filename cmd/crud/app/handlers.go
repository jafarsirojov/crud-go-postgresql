package app

import (
	"crud/pkg/crud/models"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func (receiver *server) handleBurgersList() func(http.ResponseWriter, *http.Request) {
	tpl, err := template.ParseFiles(filepath.Join(receiver.templatesPath, "index.gohtml"))
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("viewing burgers list")
		list, err := receiver.burgersSvc.BurgersList()
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		data := struct {
			Title   string
			Burgers []models.Burger
		}{
			Title:   "McDonalds",
			Burgers: list,
		}
		err = tpl.Execute(writer, data)
		if err != nil {
			log.Printf("can't execute: %v\n",err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		log.Println("viewed burger list")
	}
}

func (receiver *server) handleBurgersSave() func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request)
		log.Println("saving burger!")
		name := request.FormValue("name")
		priceString := request.FormValue("price")
		price, err := strconv.Atoi(priceString)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		model := models.Burger{
			Name:  name,
			Price: price,
		}
		log.Printf("name burger = %s, price = %d\n", model.Name, model.Price)
		receiver.burgersSvc.Save(model)
		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
		log.Println("saved burger!")
		return
	}
}

func (receiver *server) handleBurgersRemove() func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request)
		log.Println("removing burger by id!")
		idValue := request.FormValue("id")
		idInt, err := strconv.Atoi(idValue)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		log.Printf("removing burger id = %d\n",idInt)
		err = receiver.burgersSvc.RemoveById(idInt)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
		log.Println("removed burger")
		return
	}
}

func (receiver *server) handleFavicon() func(http.ResponseWriter, *http.Request) {
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
