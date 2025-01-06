package main

import (
	"fmt"
	"net/http"
	"text/template"
	"API_Pokemon\services\SetPokemon"
)

var temp, tempErr = template.ParseGlob("./templates/*.html")

func main() {

	// Récupération des templates
	if tempErr != nil {
		fmt.Println(fmt.Sprint("ERREUR => %s", tempErr.Error()))
		return
	}

	http.HandleFunc("/Carte", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/Set", func(w http.ResponseWriter, r *http.Request) {
		type DataStruct struct {
			Data []SetPokemon.Set
		}
		DataToSend := DataStruct{
			Data: SetPokemon.SetRequest(),
		}
		temp.ExecuteTemplate(w, "Set", DataToSend)
	})

	fmt.Println("http://localhost:6969/")

	http.ListenAndServe(":6969", nil)
}
