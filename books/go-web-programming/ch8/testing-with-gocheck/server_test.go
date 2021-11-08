package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "gopkg.in/check.v1" // a dot lets you access all exported identifiers without qualifying
)

type PostTestSuite struct{}

func init() {
	Suite(&PostTestSuite{}) // register our test suite
}

func Test(t *testing.T) { TestingT(t) } // run all registered suites

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func (s *PostTestSuite) TestHandleGet(c *C) {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))

	writer = httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)

	mux.ServeHTTP(writer, request)

	c.Check(writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	c.Check(post.Id, Equals, 1)
}
