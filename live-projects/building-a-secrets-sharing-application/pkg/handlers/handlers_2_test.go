package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Garren/building-a-secrets-sharing-application/pkg/store"
	"github.com/Garren/building-a-secrets-sharing-application/pkg/types"
)

var myTestSecret = "My Test Secret"

func TestGetSecretSuccess(t *testing.T) {
	var id string
	store.Init("/tmp/test.json", "password", "one")
	mux := http.NewServeMux()
	SetupHandlers(mux)

	{
		writer := httptest.NewRecorder()

		payload, _ := json.Marshal(types.CreateSecretPayload{
			PlainText: myTestSecret,
		})
		p := bytes.NewReader(payload)

		request, err := http.NewRequest("POST", "/", p)
		if err != nil {
			t.Errorf("posting request %e", err)
		}
		mux.ServeHTTP(writer, request)
		if writer.Code != http.StatusOK {
			t.Errorf("Response code is %v", writer.Code)
		}
		body := writer.Body.Bytes()
		response := types.CreateSecretResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			t.Errorf("getting response body %e", err)
		}
		id = response.Id
	}

	{
		writer := httptest.NewRecorder()

		request, _ := http.NewRequest("GET", "/"+id, nil)
		mux.ServeHTTP(writer, request)
		if writer.Code != http.StatusOK {
			t.Errorf("Response code is %v", writer.Code)
		}
		body := writer.Body.Bytes()
		response := types.GetSecretResponse{}
		json.Unmarshal(body, &response)
		if response.Data != myTestSecret {
			t.Errorf("wrong response, expecting %s, got %s",
				myTestSecret, response.Data)
		}
	}
}

func TestGetSecretFail(t *testing.T) {
	mux := http.NewServeMux()
	SetupHandlers(mux)
	{
		writer := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/", bytes.NewReader(nil))
		mux.ServeHTTP(writer, request)
		if writer.Code != http.StatusBadRequest {
			t.Errorf("Response code is %v", writer.Code)
		}
	}
	{
		writer := httptest.NewRecorder()

		payload, _ := json.Marshal(types.GetSecretResponse{
			Data: "data",
		})
		p := bytes.NewReader(payload)

		request, _ := http.NewRequest("POST", "/", p)
		mux.ServeHTTP(writer, request)
		if writer.Code != http.StatusBadRequest {
			t.Errorf("Response code is %v", writer.Code)
		}
	}
}
