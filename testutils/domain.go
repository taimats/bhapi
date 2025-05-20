package testutils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func PseudoGoogleBooksAPIServer(t *testing.T) *httptest.Server {
	t.Helper()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /books/v1/volumes", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testData, err := TestFile("response_body.json")
		if err != nil {
			t.Fatal(err)
		}
		jb := IndentForJSON(t, string(testData))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jb)
	}))

	return httptest.NewServer(mux)
}
