package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

// note - we need the db server to be running for these tests to pass
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)
	writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
	//// every test case runs independently and starts its own web server. You need
	//// to create a multiplexer and attach a handler.
	//mux := http.NewServeMux()
	//mux.HandleFunc("/post/", handleRequest)

	//// NewRecorder captures the request and response for inspection
	//writer := httptest.NewRecorder()

	// create a request
	request, _ := http.NewRequest("GET", "/post/1", nil)

	// send the request and response recorder through the multiplexer
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/post/1", handleRequest)

	//writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"Updated post","author:"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)

	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
