package testutils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func SetEnvForTest() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("CWDの取得に失敗:%s", err)
	}

	srcPath := "C:/Users/beo03/bookhistoryapi/.env"
	efn, err := filepath.Rel(cwd, srcPath)
	if err != nil {
		return fmt.Errorf("envファイルの相対パス生成に失敗:%s", err)
	}

	if err := godotenv.Load(efn); err != nil {
		return fmt.Errorf(".envファイルの取得に失敗:%s", err)
	}

	return nil
}

func NewRelativePath(targetPath string) (relativePath string, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("cwdの取得に失敗:%w", err)
	}

	relativePath, err = filepath.Rel(cwd, targetPath)
	if err != nil {
		return "", fmt.Errorf("相対パスの取得に失敗:%w", err)
	}

	return relativePath, err
}
