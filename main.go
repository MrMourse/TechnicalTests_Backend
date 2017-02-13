package main

import (
	"log"
	"net/http"
	"flag"
	"fmt"
	"html/template"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"strings"
	"strconv"
	"sort"
)


// Request print the header and the footer of the page without data
func request(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("head.html")
	t.Execute(w, nil)
	u, _ := template.ParseFiles("footer.html")
	u.Execute(w, nil)
}


// Structure which provides the url of the repository and the hash_table of the data language
type repos_hash struct{
	url string
	hashtable map[string]int
}

// Obtain the hash_table of languages and the html.url with GitHubAPI
func gather(repo *github.Repository, client *github.Client, out chan repos_hash) {
	var repos_tmp repos_hash
	// Gather them
	res, _, err := client.Repositories.ListLanguages(*repo.Owner.Login, *repo.Name)
	if err != nil {
		fmt.Println(err)
	}
	repos_tmp.hashtable = res
	repos_tmp.url = *repo.HTMLURL

	out <- repos_tmp
}



// Get the result of the post, get the 100 last repositories, take languages information, gathered them, print them
func search(w http.ResponseWriter, r *http.Request) {

	// Get the form content
	r.ParseForm()
	language := strings.Join(r.Form["Search"], "")

	// Start the query on github
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "b1abf4e31f153ad21e19cf70dabce2310a731b1c"},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	// Get the last id
	opts:= &github.SearchOptions{Order: "desc"}
	repo, _, _ :=client.Search.Repositories(">2017-02-13",opts)
	id:=*repo.Repositories[0].ID

	// Get the 100 last repositories
	opt := &github.RepositoryListAllOptions{
		Since:id-100,
		ListOptions: github.ListOptions{PerPage: 100},
	}

	// Get all pages of results
	var allRepos []*github.Repository
	repos, _, err := client.Repositories.ListAll(opt)
	if err != nil {
		fmt.Println(err)
	}
	allRepos = append(allRepos, repos...)
	reposChan := make(chan repos_hash, len(allRepos))

	// Get the languages information->by goroutine
	for _, repo := range allRepos {
		go gather(repo, client, reposChan)
	}

	// Gather the information->by channel
	result := make([]repos_hash, len(allRepos))

	for i := 0; i < len(allRepos); i++ {
		result[i] = <-reposChan
	}

	// Find the total of lines
	total :=0
	printres := make(map[int]string)
	for _,elmt:=range result  {
		if (elmt.hashtable[language]>0){
			total += elmt.hashtable[language]
			/*fmt.Printf("url :%s, number :%d\n",elmt.url,elmt.hashtable[language])*/
			number := elmt.hashtable[language]
			url := elmt.url
			printres[number] = url
		}
	}



	// Print the page

	t, _ := template.ParseFiles("head.html")
	headContent := map[string]string{
		"Title" : "Result",
		"Language" : language,
		"result": " total : "+strconv.Itoa(total)+" lines",
	}
	t.Execute(w, headContent)

	// Sort data
	var keys []int
	for k := range printres {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, k := range keys {
		t, _ := template.ParseFiles("data.html")
		dataContent := map[string]string{
			"url": printres[k],
			"name":"link of the github : "+printres[k],
			"number": strconv.Itoa(k)+" lines",
		}
		t.Execute(w, dataContent)
		/*fmt.Println("Key:", k, "Value:", printres[k])*/
	}
	u, _ := template.ParseFiles("footer.html")
	u.Execute(w, nil)

}


func main() {

	// Port id initialization
	port := flag.String("port", "3000", "server port number")

	// Requests
	http.HandleFunc("/", request)
	http.HandleFunc("/search", search)

	log.Println("Listening on :" + *port)

	// Manage request
	log.Println("Listenning...")
	err := http.ListenAndServe(":" + *port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}



