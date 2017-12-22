package book_server

import (
	"testing"
	"net/http/httptest"
	"bytes"
	"io"
	"encoding/json"
	"encoding/base64"
)

type test struct {
	url string
	method string
	body io.Reader
	book Book
	res Response
}

func TestHello(t *testing.T) {
	tests := [] test {
		{url: u + hello, res: Response{200, welcome} },
	}

	r := httptest.NewRequest("GET", tests[0].url, nil)
	res := Hello(r)
	if res != tests[0].res {
		t.Error("1 -> Expected,", tests[0].res, ",found", res)
	}
}

func TestShowBookList(t *testing.T) {
	res := Response {200, empty}

	r := httptest.NewRequest("GET", u + showBookList, nil)
	r.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte("ac:ac")))

	Books = []Book{}
	resp := ShowBookList(r)
	if resp != res {
		t.Error("1 -> Expected,", res, ",found", resp)
	}

	Books = append(Books, Book{1, "a", "a"})
	Books = append(Books, Book{2, "b", "b"})
	msg := `[
 {
  "Id": 1,
  "Title": "a",
  "Author": "a"
 },
 {
  "Id": 2,
  "Title": "b",
  "Author": "b"
 }
]`
	res.Msg = msg
	resp = ShowBookList(r)
	if resp != res {
		t.Error("2 -> Expected,", res, ",found", resp)
	}
}

func TestAddBook(t *testing.T) {
	tests := [] test {
		{url: u + addBook + "?Title=c&Author=", body: nil, method: "GET", res: Response{400, emptyField} },
		{url: u + addBook + "?Title=c&Author=c", body: nil, method: "GET", res: Response{200, added} },
		{url: u + addBook, method: "POST", book: Book{Title: "", Author: "d"}, res: Response{400, emptyField} },
		{url: u + addBook, method: "POST", book: Book{Title: "d", Author: "d"}, res: Response{200, added} },
		{url: u + addBook, method: "PUT", body: nil, res: Response{405, wrongMethod} },
	}


	for i, tc := range tests {
		if tc.method == "POST" {
			bk, err := json.Marshal(tc.book)
			if err != nil {
				t.Fatal(err)
			}
			tc.body = bytes.NewReader(bk)
		}

		r := httptest.NewRequest(tc.method, tc.url, tc.body)
		r.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte("ac:ac")))

		resp := AddBook(r)
		if resp != tc.res {
			t.Error(i, "-> Expected,", tc.res, ", found, ", resp)
		}
	}
}

func TestEditBook(t *testing.T) {
	tests := [] test {
		{url: u + editBook + "a", method: "PUT", body: nil, res: Response{400, wrongId} },
		{url: u + editBook + "1", method: "PUT", book: Book{Title: "", Author: "e"}, res: Response{400, emptyField} },
		{url: u + editBook + "1", method: "PUT", book: Book{Title: "e", Author: "e"}, res: Response{200, edited} },
		{url: u + editBook + "0", method: "PUT", book: Book{Title: "e", Author: "e"}, res: Response{400, notFound} },
		{url: u + editBook + "1", method: "GET", body: nil, res: Response{405, wrongMethod} },
	}

	for i, tc := range tests {
		if tc.method == "PUT" {
			bk, err := json.Marshal(tc.book)
			if err != nil {
				t.Fatal(err)
			}
			tc.body = bytes.NewReader(bk)
		}

		r := httptest.NewRequest(tc.method, tc.url, tc.body)
		r.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte("ac:ac")))
		resp := EditBook(r)
		if resp != tc.res {
			t.Error(i, "-> Expected,", tc.res, ", found, ", resp)
		}
	}
}

func TestDeleteBook(t *testing.T) {
	tests := [] test {
		{url: u + deleteBook + "a", method: "DELETE", body: nil, res: Response{400, wrongId} },
		{url: u + deleteBook + "1", method: "DELETE", res: Response{200, deleted} },
		{url: u + deleteBook + "0", method: "DELETE", res: Response{400, notFound} },
		{url: u + deleteBook + "1", method: "GET", body: nil, res: Response{405, wrongMethod} },
	}

	for i, tc := range tests {
		if tc.method == "PUT" {
			bk, err := json.Marshal(tc.book)
			if err != nil {
				t.Fatal(err)
			}
			tc.body = bytes.NewReader(bk)
		}

		r := httptest.NewRequest(tc.method, tc.url, tc.body)
		r.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte("ac:ac")))
		resp := DeleteBook(r)
		if resp != tc.res {
			t.Error(i, "-> Expected,", tc.res, ", found, ", resp)
		}
	}
}
