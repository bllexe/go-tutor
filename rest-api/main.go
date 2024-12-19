package main

import(
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles [] Article 

func main(){
   handleRequest()
}


func testAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println(w,"Endpoint Hit : All Articles Endpoint")
}

func handleRequest(){

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/articles",allArticles).Methods("GET")
	myRouter.HandleFunc("/articles",testAllArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "HomePage endpoint")
}


func allArticles(w http.ResponseWriter, r *http.Request){

	articles := Articles{
		Article{"First Article", "First Article Desc", "First Article Content"},
		Article{"Second Article", "Second Article Desc", "Second Article Content"},
		Article{"Third Article", "Third Article Desc", "Third Article Content"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)

}



