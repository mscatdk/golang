package comment

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CommentDAO Memory data store
var CommentDAO DAO

func jsonToComment(w http.ResponseWriter, r *http.Request) Comment {
	var comment Comment
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &comment); err != nil {
		fmt.Println("Wrong format")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(400)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		panic(err)
	}

	return comment
}

//UpdateComment - Implement the comment API put handler
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID := vars["commentId"]
	i, err := strconv.Atoi(commentID)
	if err != nil {
		panic(err)
	}

	comment := jsonToComment(w, r)
	if comment.ID != i {
		w.Header().Set("Content-Type", "plain/text; charset=UTF-8")
		w.WriteHeader(400)
		fmt.Fprintf(w, "Wrong ID")
		return
	}

	CommentDAO.UpdateComment(comment)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}

}

//CreateComment - implement the comment API POST handler
func CreateComment(w http.ResponseWriter, r *http.Request) {
	comment := jsonToComment(w, r)
	comment = CommentDAO.AddComment(comment)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

//ListComments - implement the GET handler
func ListComments(w http.ResponseWriter, r *http.Request) {
	comments := CommentDAO.List()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}
