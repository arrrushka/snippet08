package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)
// handler
func home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}

	w.Write([]byte("Hello Tleuzhan Mukatayev"))
}

func showSnippet(w http.ResponseWriter, r *http.Request)  {
	id,err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}

	fmt.Fprintf(w, "Display a snippet with id %v", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Create a specific snippet..."))
}

func main()  {
	// router
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/snippet",showSnippet)
	mux.HandleFunc("/snippet/create",createSnippet)

	//default server
	err := http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}