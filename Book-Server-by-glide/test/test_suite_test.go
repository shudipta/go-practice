package book_server_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "Go-Practice/Book-Server-by-glide/book_server"

)

func TestTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Book Server Test Suite")
}

var _ = BeforeSuite(func() {
	HandleRequests()
	go StartServer("10000", false)
})

var _ = AfterSuite(func() {
	ShutdownServer()
})
