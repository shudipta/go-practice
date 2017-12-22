package book_server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	//. "Go-Practice/Book-Server-by-glide/test"
	. "Go-Practice/Book-Server-by-glide/book_server"
	"io"
	"net/http/httptest"
	"encoding/base64"
	"net/http"
	"encoding/json"
	"bytes"
	"log"
)

var _ = Describe("BookServer", func() {
	type test struct {
		r *http.Request
		url string
		method string
		body io.Reader
		book Book
		res Response
	}

	var (
		res Response
		tests []test
		r *http.Request

		u = "http://localhost:10000"
		hello = "/"
		showBookList = "/showBookList"
		addBook = "/addBook"
		editBook = "/editBook/"
		deleteBook = "/deleteBook/"
		welcome = "Welcome to the \"Book Server\""
		empty = "There is no book"
		emptyField = "contains empty field"
		added = "added successfully"
		wrongMethod = "requested method is not allowed"
		wrongId = "id is required to be an integer"
		edited = "edited successfully"
		notFound = "requested book isn't found"
		deleted = "deleted successfully"
	)

	Describe("Hello Page", func() {
		Context("url only", func() {
			BeforeEach(func() {
				tests = append(tests, test {
					url: u + hello, res: Response{200, welcome},
				})
			})

			It("Should be welcome msg", func() {
				r := httptest.NewRequest("GET", tests[0].url, nil)
				Expect(Hello(r)).To(Equal(tests[0].res))
			})
		})
	})

	Describe("Show Book List Page", func() {
		Context("(1) any method and url only", func() {
			JustBeforeEach(func() {
				res = Response {200, empty}
				Books = []Book{}
				r = httptest.NewRequest("GET", u + showBookList, nil)
				r.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte("ac:ac")))
			})


			It("Should be empty msg", func() {
				Expect(ShowBookList(r)).To(Equal(res))
			})
		})

		Context("(2) any method and url only", func() {
			JustBeforeEach(func() {
				res = Response {200, empty}
				Books = []Book{}
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
				r = httptest.NewRequest("GET", u + showBookList, nil)
				r.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte("ac:ac")))
			})

			It("Should be shown the book list", func() {

				Expect(ShowBookList(r)).To(Equal(res))
			})
		})
	})

	Describe("Add Book Page", func() {
		Context("", func() {
			BeforeEach(func() {
				tests = [] test {
					{url: u + addBook + "?Title=c&Author=", body: nil, method: "GET", res: Response{400, emptyField} },
					{url: u + addBook + "?Title=c&Author=c", body: nil, method: "GET", res: Response{200, added} },
					{url: u + addBook, method: "POST", book: Book{Title: "", Author: "d"}, res: Response{400, emptyField} },
					{url: u + addBook, method: "POST", book: Book{Title: "d", Author: "d"}, res: Response{200, added} },
					{url: u + addBook, method: "PUT", body: nil, res: Response{405, wrongMethod} },
				}
				for i, _ := range tests {
					if tests[i].method == "POST" {
						bk, err := json.Marshal(tests[i].book)
						if err != nil {
							log.Fatal(err)
						}
						tests[i].body = bytes.NewReader(bk)
					}

					tests[i].r = httptest.NewRequest(tests[i].method, tests[i].url, tests[i].body)
					tests[i].r.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte("ac:ac")))

				}
			})

			It("Should be,", func() {
				Expect(AddBook(tests[0].r)).To(Equal(tests[0].res))
			})

			It("Should be,", func() {
				Expect(AddBook(tests[1].r)).To(Equal(tests[1].res))
			})

			It("Should be,", func() {
				Expect(AddBook(tests[2].r)).To(Equal(tests[2].res))
			})

			It("Should be,", func() {
				Expect(AddBook(tests[3].r)).To(Equal(tests[3].res))
			})

			It("Should be,", func() {
				Expect(AddBook(tests[4].r)).To(Equal(tests[4].res))
			})
		})
	})

	Describe("Edit Book Page", func() {
		Context("", func() {
			BeforeEach(func() {
				tests = []test{}
				tests = [] test {
					{url: u + editBook + "a", method: "PUT", res: Response{400, wrongId} },
					{url: u + editBook + "1", method: "PUT", book: Book{Title: "", Author: "e"}, res: Response{400, emptyField} },
					{url: u + editBook + "1", method: "PUT", book: Book{Title: "e", Author: "e"}, res: Response{200, edited} },
					{url: u + editBook + "0", method: "PUT", book: Book{Title: "e", Author: "e"}, res: Response{400, notFound} },
					{url: u + editBook + "1", method: "GET", body: nil, res: Response{405, wrongMethod} },
				}
				for i, _ := range tests {
					if tests[i].method == "PUT" {
						bk, err := json.Marshal(tests[i].book)
						if err != nil {
							log.Fatal(err)
						}
						tests[i].body = bytes.NewReader(bk)
					}

					tests[i].r = httptest.NewRequest(tests[i].method, tests[i].url, tests[i].body)
					tests[i].r.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte("ac:ac")))
				}
			})

			It("Should be,", func() {
				Expect(EditBook(tests[0].r)).To(Equal(tests[0].res))
			})

			It("Should be,", func() {
				Expect(EditBook(tests[1].r)).To(Equal(tests[1].res))
			})

			It("Should be,", func() {
				Expect(EditBook(tests[2].r)).To(Equal(tests[2].res))
			})

			It("Should be,", func() {
				Expect(EditBook(tests[3].r)).To(Equal(tests[3].res))
			})

			It("Should be,", func() {
				Expect(EditBook(tests[4].r)).To(Equal(tests[4].res))
			})
		})
	})

	Describe("Edit Book Page", func() {
		Context("", func() {
			BeforeEach(func() {
				tests = []test{}
				tests = [] test {
					{url: u + deleteBook + "a", method: "DELETE", body: nil, res: Response{400, wrongId} },
					{url: u + deleteBook + "1", method: "DELETE", res: Response{200, deleted} },
					{url: u + deleteBook + "0", method: "DELETE", res: Response{400, notFound} },
					{url: u + deleteBook + "1", method: "GET", body: nil, res: Response{405, wrongMethod} },
				}
				for i, _ := range tests {
					//By("" + strconv.Itoa(i + 1))
					if tests[i].method == "PUT" {
						bk, err := json.Marshal(tests[i].book)
						if err != nil {
							log.Fatal(err)
						}
						tests[i].body = bytes.NewReader(bk)
					}

					tests[i].r = httptest.NewRequest(tests[i].method, tests[i].url, tests[i].body)
					tests[i].r.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte("ac:ac")))
				}
			})

			It("Should be,", func() {
				Expect(DeleteBook(tests[0].r)).To(Equal(tests[0].res))
			})

			It("Should be,", func() {
				Expect(DeleteBook(tests[1].r)).To(Equal(tests[1].res))
			})

			It("Should be,", func() {
				Expect(DeleteBook(tests[2].r)).To(Equal(tests[2].res))
			})

			It("Should be,", func() {
				Expect(DeleteBook(tests[3].r)).To(Equal(tests[3].res))
			})

		})
	})
})
