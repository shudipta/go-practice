package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

// Adapter type
type Adapter func(http.HandlerFunc) http.HandlerFunc

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
	<html>
	  <head>
		  <title>Hello World</title>
	  </head>
	  <body>
		  Hello World!
	  </body>
	</html>`,
	)
}

func addHeader(k, v string) Adapter {
	return func(init http.HandlerFunc) http.HandlerFunc {
		return func(res http.ResponseWriter, req *http.Request) {
			req.Header.Set(k, base64.StdEncoding.EncodeToString([]byte(v)))
			
			fmt.Println(01)
			fmt.Fprintln(res, 01)
			init(res, req)
		}
	}
}

func checkAuth(t string) Adapter {
	return func(init http.HandlerFunc) http.HandlerFunc {
		return func(res http.ResponseWriter, req *http.Request) {
			encodedInfo := req.Header.Get("a")
			fmt.Fprintln(res, 02)
			decodedInfo, err := base64.StdEncoding.DecodeString(encodedInfo)
			if err != nil {
				fmt.Fprintln(res, "failed decode")
				return
			}

			if string(decodedInfo) != t {
				fmt.Fprintln(res, "wrong header")
				return
			}
			init(res, req)
		}
	}
}

func adapt(h http.HandlerFunc, adapters ...Adapter) http.HandlerFunc {
	// init := func(res http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprintln(res, 00)
	// }
	for _, adapter := range adapters {
		// init = adapter(init)
		h = adapter(h)
	}
	return h
}

func main() {
	http.HandleFunc("/", adapt(hello, 
		checkAuth("b"),
		addHeader("a", "b"),
	))

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":9000", nil)
}
