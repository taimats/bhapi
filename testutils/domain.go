package testutils

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func GenerateTestJSON(t *testing.T) []byte {
	t.Helper()

	targetPath := "C:/Users/beo03/bookhistoryapi/domain/search_test/sample2.json"
	jb, err := os.ReadFile(targetPath)
	if err != nil {
		t.Fatal(err)
	}

	return jb
}

func PseudoGoogleBooksAPIServer(t *testing.T) *httptest.Server {
	t.Helper()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /books/v1/volumes", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jb := GenerateTestJSON(t)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jb)
	}))

	ts := httptest.NewServer(mux)

	return ts
}
