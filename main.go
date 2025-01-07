package main

import (
	Carte "api/services/CartePokemon"
	"api/services/SetPokemon"
	"fmt"
	"net/http"
	"text/template"
)

var temp, tempErr = template.ParseGlob("./templates/*.html")

func main() {

	// Récupération des templates
	if tempErr != nil {
		fmt.Printf("ERREUR => %s\n", tempErr.Error())
		return
	}

	http.HandleFunc("/Carte", func(w http.ResponseWriter, r *http.Request) {
		type DataStruct struct {
			Data []Carte.Carte
		}
		DataToSend := DataStruct{
			Data: Carte.CarteP(),
		}
		temp.ExecuteTemplate(w, "Carte", DataToSend)
	})

	http.HandleFunc("/Set", func(w http.ResponseWriter, r *http.Request) {
		type DataStruct struct {
			Data []SetPokemon.Set
		}
		DataToSend := DataStruct{
			Data: SetPokemon.SetP(),
		}
		temp.ExecuteTemplate(w, "Set", DataToSend)
	})

	fmt.Println("http://localhost:6969/")

	http.ListenAndServe(":6969", nil)
}
