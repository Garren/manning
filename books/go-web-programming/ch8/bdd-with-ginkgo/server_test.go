package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("handle get", func() {

		// create a request
		request, _ := http.NewRequest("GET", "/post/1", nil)

		// send the request and response recorder through the multiplexer
		mux.ServeHTTP(writer, request)

		if writer.Code != 200 {
			GinkgoT().Errorf("Response code is %v", writer.Code)
		}

		var post Post
		json.Unmarshal(writer.Body.Bytes(), &post)
		if post.Id != 1 {
			GinkgoT().Error("Cannot retrieve JSON post")
		}
	})
	It("handle put", func() {

		//writer := httptest.NewRecorder()
		json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
		request, _ := http.NewRequest("PUT", "/post/1", json)

		mux.ServeHTTP(writer, request)

		if writer.Code != 200 {
			GinkgoT().Errorf("Response code is %v", writer.Code)
		}
	})
})
var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))
	writer = httptest.NewRecorder()
}
