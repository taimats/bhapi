package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type customValidator struct {
	validator *validator.Validate
}

func NewCustomValidator(v *validator.Validate) *customValidator {
	return &customValidator{validator: v}
}

func (cv *customValidator) Validate(s any) error {
	if err := cv.validator.Struct(s); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "不正なリクエストです")
	}

	return nil
}
