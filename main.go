package main

import (
	"log"
	"net/http"
/*	"strings"*/
	"flag"
/*	"github.com/google/go-github/github"*/
	"fmt"
	"html/template"
/*	"os"
	"io"
	"strings"*/
	"io/ioutil"
	/*
	"bytes"
	*/

	"strings"
	/*"strconv"*/
	"strconv"
)

func  get_body(URL string) string {
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("body content : %s", body)
	/*n := bytes.Index(body, []byte{0})*/
	s := string(body[:])
	return s
}
/*return an hash table*/
func hash_table(body string) map[string]int{

	/*s := '{"C++":20754080,"Python":17160192,"Jupyter Notebook":1833840,"TypeScript":769672,"Go":763825}'*/
	tmp := strings.Split(body, "{")
	res := tmp[1]
	tmp = strings.Split(res,"}")
	res = tmp[0]
	tmp = strings.Split(res, ",")
	m:= make(map[string]int)
	for elmt := range tmp  {
		tab := strings.Split(tmp[elmt], ":")
		tmp := strings.Split(tab[0], "\"")
		value, _ := strconv.Atoi(tab[1])
		m[tmp[1]] = value
	}

	return m
}

func request(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
	}

func search(w http.ResponseWriter, r *http.Request) {
	/*Get the form content*/
	/*r.ParseForm()
	// logic part of log in
	fmt.Println("language: ", r.Form["language"])
/*	language := strings.Join(r.Form["language"], "")*/


/*	client := github.NewClient(nil)

	fmt.Println("Repos that contain the query language.")

	*//*query := fmt.Sprintf("language:" +  language)*//*
	//get the 100 last repositories
	opt := &github.RepositoryListAllOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	// get all pages of results
	var allRepos []*github.Repository
	repos, _, err := client.Repositories.ListAll(opt)
	if (err !=nil){
		fmt.Println(err)
	}
	//gather them
	allRepos = append(allRepos, repos...)
	for _, repo := range allRepos {
	//filter them

		fmt.Println("repo: ", *repo.FullName)
		fmt.Println("owner: ", *repo.Owner.Login)
		fmt.Println("languages:", * repo.LanguagesURL)

	}*/
	//URL treatment
	body := get_body("http://api.github.com/repos/tensorflow/tensorflow/languages")
	fmt.Printf("s content : %s\n", body)
	res := hash_table(body)
	fmt.Printf("result : %d\n",res["C++"])
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



