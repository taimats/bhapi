package auth_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/taimats/bhapi/presenter/middleware/auth"
)

func TestAuthenticate(t *testing.T) {
	//Arrange
	err := godotenv.Load(".env")
	if err != nil {
		t.Fatal(err)
	}
	apikey := os.Getenv("BACK_API_KEY")

	//Act
	ok, err := auth.Authenticate(apikey)
	if err != nil {
		t.Fatal(err)
	}

	//Assert
	if ok != true {
		t.Error(err)
	}
}
