package auth_test

import (
	"os"
	"testing"

	"github.com/taimats/bhapi/presenter/middleware/auth"
	"github.com/taimats/bhapi/testutils"
)

func TestAuthenticate(t *testing.T) {
	//Arrange
	err := testutils.DotEnv()
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
