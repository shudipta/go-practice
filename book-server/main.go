package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"encoding/base64"
)

type Book struct {
	Id	int	`json:"Id"`
	Title   string	`json:"Title"`
	Author	string	`json:"Author"`
}

var books []Book

func response(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode == http.StatusUnauthorized {
		w.Header().Add("WWW-Authenticate", `Basic realm="Authorization Required"`)
	}
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, msg)

}

func checkAuth(w http.ResponseWriter, r *http.Request) bool {

	encodedInfo := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(encodedInfo) != 2 {
		return false
	}

	decodedInfo, err := base64.StdEncoding.DecodeString(encodedInfo[1])
	if err != nil {
		return false
	}

	userInfo := strings.SplitN(string(decodedInfo), ":", 2)
	if len(userInfo) != 2 {
		return false
	}

	if userInfo[0] != "ac" || userInfo[1] != "ac" {
		return false
	}

	return true
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" \"/\" page")
	response(w, http.StatusOK, "Welcome to the \"Book Server\"")
}

func showBookList(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" \"/showBookList\" page")

	if !checkAuth(w, r) {
		response(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if len(books) == 0 {
		response(w, http.StatusOK, "There is no book")
		return
	}

	list, convertErr := json.MarshalIndent(books, "", " ")
	if convertErr != nil {
		response(w, http.StatusInternalServerError, "Error occured in converting into json is " + convertErr.Error())
		return
	}

	response(w, http.StatusOK, string(list))
}

func addBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" \"addBook\" page")

	if !checkAuth(w, r) {
		response(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var book Book

	if r.Method == "GET" {
		data := r.URL.Query()
		book = Book {Title: data["Title"][0], Author: data["Author"][0]}
	} else if r.Method == "POST" {
		convertErr := json.NewDecoder(r.Body).Decode(&book)
		defer r.Body.Close()

		if convertErr != nil {
			response(w, http.StatusInternalServerError, "error getting json data in PUT method")
			return
		}
	} else {
		response(w, http.StatusMethodNotAllowed, "requested method is not allowed")
		return
	}

	if book.Title == "" || book.Author == "" {
		response(w, http.StatusBadRequest, "contains empty field")
		return
	}

	book.Id = len(books) + 1
	books = append(books, book)

	response(w, http.StatusOK, "added successfully")
}

func editBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" \"editBook\" page")

	if !checkAuth(w, r) {
		response(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var book Book

	if r.Method == "PUT" {
		id, idErr := strconv.Atoi(r.URL.Path[len("/editBook/"):])

		if idErr != nil {
			response(w, http.StatusBadRequest, "id is required to be an integer")
			return
		}

		convertErr := json.NewDecoder(r.Body).Decode(&book)
		defer r.Body.Close()

		fmt.Println(book)
		if convertErr != nil {
			response(w, http.StatusInternalServerError, "error getting json data in PUT method")
			return
		}

		if book.Title == "" || book.Author == "" {
			response(w, http.StatusBadRequest, "contains empty field")
			return
		}

		book.Id = id

		for i, _ := range books{
			if i + 1 == id {
				books[i] = book

				response(w, http.StatusOK, "updated successfully")
				return
			}
		}

		response(w, http.StatusBadRequest, "requested book isn't found")
	} else {
		response(w, http.StatusMethodNotAllowed, "requested method is not allowed")
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" \"deleteBook\" page")

	if !checkAuth(w, r) {
		response(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if r.Method == "DELETE" {
		id, idErr := strconv.Atoi(r.URL.Path[len("/deleteBook/"):])

		if idErr != nil {
			response(w, http.StatusBadRequest, "id is required to be an integer")
			return
		}

		for i, _ := range books{
			if i + 1 == id {
				books = append(books[:i], books[i+1:]...)
				for j, _ := range books {
					books[j].Id = j + 1
				}

				response(w, http.StatusOK, "deleted successfully")
				return
			}
		}

		response(w, http.StatusBadRequest, "requested book isn't found")
	} else {
		response(w, http.StatusMethodNotAllowed, "requested method is not allowed")
	}
}

func handleRequests() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/showBookList", showBookList)
	http.HandleFunc("/addBook", addBook)
	http.HandleFunc("/editBook/", editBook)
	http.HandleFunc("/deleteBook/", deleteBook)

	serverErr :=http.ListenAndServe(":10000", nil)

	if serverErr != nil {
		log.Fatal("Server Error:", serverErr)
	}
}

func main() {
	handleRequests()
}
