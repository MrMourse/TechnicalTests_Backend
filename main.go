package main

import (
	"log"
	"net/http"
	"fmt"
	"html/template"
)

func Home( w http.ResponseWriter, req *http.Request){
	render(w,"index.html")
}
func Search( w http.ResponseWriter, req *http.Request){
	render(w, "result.html")
}

func render(w http.ResponseWriter, tmpl string){
	tmpl =  fmt.Sprintf("templates/%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, "")
	if err!= nil {
		log.Print("template executing error: ", err)
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/search", Search)
	//manage request
	log.Println("Listenning...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


}
