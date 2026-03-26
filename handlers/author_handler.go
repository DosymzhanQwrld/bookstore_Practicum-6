package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"

	_ "github.com/gorilla/mux"
)

var authors = make(map[int]models.Author)
var nextAuthorID = 1

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var authorList []models.Author
	for _, author := range authors {
		authorList = append(authorList, author)
	}
	json.NewEncoder(w).Encode(authorList)
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	author.ID = nextAuthorID
	nextAuthorID++
	authors[author.ID] = author
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}
