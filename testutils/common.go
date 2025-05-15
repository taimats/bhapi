package testutils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// 実行場所からルート直下のenvファイル（bhapi/.env）を設定
func DotEnv() error {
	usrDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("ユーザーディレクトリの取得に失敗:%w", err)
	}
	targetPath := filepath.Join(usrDir, "bhapi", ".env")

	envfn, err := RelPath(targetPath)
	if err != nil {
		return fmt.Errorf("相対パスの取得に失敗:%w", err)
	}

	if err := godotenv.Load(envfn); err != nil {
		return fmt.Errorf(".envファイルのロードに失敗:%w", err)
	}

	return nil
}

// 現在のディレクトリから指定パスまでの相対パスを生成
func RelPath(targetPath string) (relPath string, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("現在のディレクトリの取得に失敗:%w", err)
	}

	relPath, err = filepath.Rel(cwd, targetPath)
	if err != nil {
		return "", fmt.Errorf("相対パスの取得に失敗:%w", err)
	}

	return relPath, nil
}
