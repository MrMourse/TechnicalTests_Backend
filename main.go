package main

import (
	"log"
	"net/http"
	"strings"
	"flag"
	"github.com/google/go-github/github"
	"fmt"
	"html/template"
)


type Content struct{
	Path string
	Synopsis string
	Language string
	Length float64
	GitHubFullName string
}

func request(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("language: ", r.Form["language"])
	}

}

func main() {

	//On initialise la valeur du port
	port := flag.String("port", "3000", "server port number")

	http.HandleFunc("/", request)

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {

		language := strings.TrimSpace(r.URL.Query().Get("language"))


		if language != "" {

			log.Println("Query: ", language)
		}

		client := github.NewClient(nil)

		fmt.Println("Repos that contain the query language.")

		query := fmt.Sprintf("language:"+language)

		opts := &github.SearchOptions{
			Sort: "stars",
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		}

		repos, _, err := client.Search.Repositories(query, opts)

		if err != nil {
			fmt.Printf("error: %v\n\n", err)
		} else {
			fmt.Printf("%v\n\n", github.Stringify(repos))
		}

	})

	log.Println("Listening on :" + *port)
	//manage request
	log.Println("Listenning...")
	err := http.ListenAndServe(":" + *port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}



