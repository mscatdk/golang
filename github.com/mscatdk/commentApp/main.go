package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mscatdk/commentApp/comment"
	"github.com/mscatdk/commentApp/pages"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", commentPages.LandingPage)
	router.Methods("GET").Path("/API/comment").Name("listComments").HandlerFunc(comment.ListComments)
	router.Methods("POST").Path("/API/comment").Name("createComment").HandlerFunc(comment.CreateComment)
	router.Methods("PUT").Path("/API/comment/{commentId}").Name("updateComment").HandlerFunc(comment.UpdateComment)

	log.Fatal(http.ListenAndServe(":8080", router))
}
