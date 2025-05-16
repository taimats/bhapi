package handler

import "github.com/labstack/echo/v4"

func RegisterHandlersWithBaseURL(router EchoRouter, hi HandlerInterface) {
	const baseURL = "/v1"

	router.POST(baseURL+"/auth/register", hi.PostAuthRegister)
	router.GET(baseURL+"/charts/:authUserId", hi.GetChartsWithAuthUserId)
	router.GET(baseURL+"/health", hi.GetHealth)
	router.GET(baseURL+"/health/db", hi.GetHealthDb)
	router.GET(baseURL+"/records/:authUserId", hi.GetRecordsWithAuthUserId)
	router.GET(baseURL+"/search", hi.GetSearch)
	router.DELETE(baseURL+"/shelf/:authUserId", hi.DeleteShelfWithAuthUserId)
	router.GET(baseURL+"/shelf/:authUserId", hi.GetShelfWithAuthUserId)
	router.POST(baseURL+"/shelf/:authUserId", hi.PostShelfAuthUserId)
	router.PUT(baseURL+"/shelf/:authUserId", hi.PutShelfWithAuthUserId)
	router.PUT(baseURL+"/users", hi.PutUsers)
	router.DELETE(baseURL+"/users/:authUserId", hi.DeleteUsersWithAuthUserId)
	router.GET(baseURL+"/users/:authUserId", hi.GetUsersWithAuthUserId)
}

type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}
