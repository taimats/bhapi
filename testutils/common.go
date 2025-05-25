package testutils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// 実行場所からルート直下のenvファイル（bhapi/.env）を設定
func DotEnv() error {
	envfn, err := searchFilePath(".env")
	if err != nil {
		return err
	}
	if err := godotenv.Load(envfn); err != nil {
		return fmt.Errorf(".envファイルのロードに失敗:%w", err)
	}
	return nil
}

// 実行場所からtestutils/testdata以下にあるファイルを取得
func TestFile(fileName string) (testData []byte, err error) {
	root, err := moduleRoot()
	if err != nil {
		return nil, err
	}
	target := filepath.Join(root, "testutils", "testdata", fileName)
	testData, err = os.ReadFile(target)
	if err != nil {
		return nil, fmt.Errorf("テストデータの取得に失敗:%w", err)
	}
	return testData, nil
}

// 現在のディレクトリから親ディレクトリへと探索し、
// 指定のファイル名があればそのパス名を返す。
// エラーは指定したパス名が存在しない場合に発生。
func searchFilePath(filename string) (string, error) {
	current, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		path := filepath.Join(current, filename)
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			return path, nil
		}
		parent := filepath.Dir(current)
		if parent == current {
			return "", errors.New("一致するパスはありません")
		}
		current = parent
	}
}

// 現在のディレクトリから親ディレクトリへと探索し、
// 本モジュールのルートパスを返す。
// エラーは内部パッケージのエラー発生時のみ返す。
func moduleRoot() (string, error) {
	target := "bhapi"
	current, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		dirName := filepath.Base(current)
		if dirName == target {
			return current, nil
		}
		current = filepath.Dir(current)
	}
}
