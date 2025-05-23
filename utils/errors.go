package utils

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound  = errors.New("リソースがありません")
	ErrAlrExists = errors.New("リソースはすでに存在します")
)

// エラー元（＝ origin）を最新エラー（= errNow）でwrapする。
// errNowには現時点で発生したerrorを入れる。originにはラップしたい元errorを入れる。
// ラップする必要がない場合、originはnilにしてよい。
func NewErrChains(errNow error, origin error) error {
	if origin == nil {
		return errNow
	}
	return fmt.Errorf("%w:%w", errNow, origin)
}
