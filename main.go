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

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main()  {
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet", showSnippet)
	http.HandleFunc("/snippet/create", createSnippet)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}