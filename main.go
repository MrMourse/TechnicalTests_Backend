package main

import (
	"log"
	"net/http"
/*	"strings"*/
	"flag"
	"github.com/google/go-github/github"
	"fmt"
	"html/template"
)

func request(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
	}

func search(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	// logic part of log in
	fmt.Println("language: ", r.Form["language"])
/*	language := strings.Join(r.Form["language"], "")*/
	client := github.NewClient(nil)

	fmt.Println("Repos that contain the query language.")

	/*query := fmt.Sprintf("language:" +  language)*/

	opt := &github.RepositoryListAllOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	// get all pages of results
	var allRepos []*github.Repository
	repos, _, _ := client.Repositories.ListAll(opt)

	allRepos = append(allRepos, repos...)
	for _, repo := range allRepos {

		fmt.Println("repo: ", *repo.FullName)
		fmt.Println("owner: ", *repo.Owner.Login)


	}



	t, _ := template.ParseFiles("result.html")
	t.Execute(w, nil)
}


func main() {

	//On initialise la valeur du port
	port := flag.String("port", "3000", "server port number")

	http.HandleFunc("/", request)
	http.HandleFunc("/search", search)

	log.Println("Listening on :" + *port)
	//manage request
	log.Println("Listenning...")
	err := http.ListenAndServe(":" + *port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}



