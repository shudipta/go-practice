# book-server

Run the following command to run the server in a terminal:

```console
go run main.go
```

Send requests via `curl` from another terminal terminal:

```console
# home page
$ curl http://localhost:8443
Welcome to the "Book Server"

# show existing book list
$ curl http://localhost:8443/showBookList \
    --header "Content-Type: application/json" \
    --request GET
There is no book

# add a book
$ curl http://localhost:8443/addBook \
    --header "Content-Type: application/json" \
    --request POST \
    --data "{\"Title\":\"aaa\",\"Author\":\"AAA\"}}"
added successfully

# now show the book list
$ curl http://localhost:8443/showBookList \
    --header "Content-Type: application/json" \
    --request GET
[
 {
  "Id": 1,
  "Title": "aaa",
  "Author": "AAA"
 }
]

# add another book
$ curl http://localhost:8443/addBook \
    --header "Content-Type: application/json" \
    --request POST \
    --data "{\"Title\":\"ccc\",\"Author\":\"CCC\"}"
added successfully

# now show the book list again
$ curl http://localhost:8443/showBookList \
    --header "Content-Type: application/json" \
    --request GET
[
 {
  "Id": 1,
  "Title": "aaa",
  "Author": "AAA"
 },
 {
  "Id": 2,
  "Title": "ccc",
  "Author": "CCC"
 }
]

# update the book with id 2
$ curl http://localhost:8443/editBook/2 \
    --header "Content-Type: application/json" \
    --request PUT \
    --data "{\"Title\":\"bbb\",\"Author\":\"BBB\"}"
edited successfully

# now show the book list for changes
$ curl http://localhost:8443/showBookList \
    --header "Content-Type: application/json" \
    --request GET
[
 {
  "Id": 1,
  "Title": "aaa",
  "Author": "AAA"
 },
 {
  "Id": 2,
  "Title": "bbb",
  "Author": "BBB"
 }
]

# delete the book with id 2
$ curl http://localhost:8443/deleteBook/2 \
    --request DELETE
deleted successfully

# finally show the book list that the book with id 2 is no more exists
$ curl http://localhost:8443/showBookList \
    --header "Content-Type: application/json" \
    --request GET
[
 {
  "Id": 1,
  "Title": "aaa",
  "Author": "AAA"
 }
]