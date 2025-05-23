package auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	"github.com/taimats/bhapi/utils"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrAuthEmty       = errors.New("キーが空です")
	ErrAuthDecFail    = errors.New("base64urlデコードに失敗")
	ErrAuthInvalidKey = errors.New("不正なトークンです")
)

// Bearerのapikeyを検証
func Authenticate(apikey string) (bool, error) {
	if apikey == "" {
		return false, utils.NewErrChains(ErrAuthEmty, nil)
	}

	decodedKey, err := base64.URLEncoding.DecodeString(apikey)
	if err != nil {
		return false, utils.NewErrChains(ErrAuthDecFail, err)
	}

	src := os.Getenv("TOKEN_SEED")
	err = bcrypt.CompareHashAndPassword(decodedKey, []byte(src))
	if err != nil {
		return false, utils.NewErrChains(ErrAuthInvalidKey, err)
	}

	return true, nil
}

// TOKEN_SEEDを変更したときに使用（サーバーの検証には使わない）
func GenerateSource() (string, error) {
	ts := os.Getenv("TOKEN_SEED")

	hashed, err := bcrypt.GenerateFromPassword([]byte(ts), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("ハッシュの生成に失敗:%w", err)
	}

	return string(hashed), nil
}

func IssueAPIKey(src string) string {
	apikey := base64.URLEncoding.EncodeToString([]byte(src))
	return apikey
}
