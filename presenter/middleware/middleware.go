package middleware

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/taimats/bhapi/presenter/handler"
	"github.com/taimats/bhapi/presenter/middleware/auth"
	"github.com/taimats/bhapi/presenter/middleware/loggers"
	"golang.org/x/time/rate"
)

var (
	allowedOrigins = []string{os.Getenv("FRONT_API_BASE_URL")}
	allowedMethods = []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete}
	allowedHeaders = []string{echo.HeaderContentType, echo.HeaderAuthorization}

	authSkippedPaths = map[string]struct{}{"/v1/health": {}, "/v1/health/db": {}}
)

// echoインスタンスに対して必要なすべてのmiddlewareを設定する。
func SetAll(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Recover())

	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	e.Use(middleware.RequestLoggerWithConfig(
		loggers.NewRequestLoggerConfig(l),
	))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowedOrigins,
		AllowMethods: allowedMethods,
		AllowHeaders: allowedHeaders,
	}))

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper: func(c echo.Context) bool {
			_, ok := authSkippedPaths[c.Request().URL.Path]
			return ok
		},

		KeyLookup:  "header:" + echo.HeaderAuthorization,
		AuthScheme: "Bearer",
		Validator: func(key string, c echo.Context) (bool, error) {
			ok, err := auth.Authenticate(key)
			return ok, err
		},
	}))

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(3))))

	e.Validator = handler.NewCustomValidator(validator.New())

	return e
}
