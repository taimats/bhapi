package utils

import (
	"errors"
)

var ErrNotFound = errors.New("リソースがありません")
var ErrAlrExists = errors.New("すでに存在します")

func NewErrNotFound() error {
	return ErrNotFound
}

func NewErrAlrExists() error {
	return ErrNotFound
}
