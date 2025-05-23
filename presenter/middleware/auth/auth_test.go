package auth_test

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taimats/bhapi/presenter/middleware/auth"
	"github.com/taimats/bhapi/testutils"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticate(t *testing.T) {
	//Arrange
	err := testutils.DotEnv()
	if err != nil {
		t.Fatal(err)
	}
	a := assert.New(t)
	tests := map[string]struct {
		key     string
		want    bool
		isErr   bool
		errWant error
	}{
		"OK:認証キーが承認": {
			key:     os.Getenv("BACK_API_KEY"),
			want:    true,
			isErr:   false,
			errWant: nil,
		},
		"NG:認証キーが不正": {
			key:     testAPIKey(t),
			want:    false,
			isErr:   true,
			errWant: auth.ErrAuthInvalidKey,
		},
		"NG:認証キーのデコード失敗": {
			key:     os.Getenv("BACK_API_KEY") + "ng",
			want:    false,
			isErr:   true,
			errWant: auth.ErrAuthDecFail,
		},
		"NG:認証キーが空": {
			key:     "",
			want:    false,
			isErr:   true,
			errWant: auth.ErrAuthEmty,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := auth.Authenticate(test.key)
			if test.isErr {
				a.Equal(false, got)
				a.ErrorIs(err, test.errWant)
				return
			}
			a.Equal(true, got)
			a.Nil(err)
		})
	}
}

// 認証キーを出力するヘルパー関数
func testAPIKey(t *testing.T) string {
	t.Helper()
	seed := "test"
	hashed, err := bcrypt.GenerateFromPassword([]byte(seed), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("ハッシュに失敗:%v", err)
	}
	return base64.URLEncoding.EncodeToString([]byte(hashed))
}
