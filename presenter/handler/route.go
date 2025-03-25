package handler

import "github.com/labstack/echo/v4"

func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface) {
	const baseURL = "/v1"

	router.POST(baseURL+"/auth/register", si.PostAuthRegister)
	router.GET(baseURL+"/charts/:authUserId", si.GetChartsWithAuthUserId)
	router.GET(baseURL+"/health", si.GetHealth)
	router.GET(baseURL+"/health/db", si.GetHealthDb)
	router.GET(baseURL+"/records/:authUserId", si.GetRecordsWithAuthUserId)
	router.GET(baseURL+"/search", si.GetSearch)
	router.DELETE(baseURL+"/shelf/:authUserId", si.DeleteShelfWithAuthUserId)
	router.GET(baseURL+"/shelf/:authUserId", si.GetShelfWithAuthUserId)
	router.POST(baseURL+"/shelf/:authUserId", si.PostShelfAuthUserId)
	router.PUT(baseURL+"/shelf/:authUserId", si.PutShelfWithAuthUserId)
	router.PUT(baseURL+"/users", si.PutUsers)
	router.DELETE(baseURL+"/users/:authUserId", si.DeleteUsersWithAuthUserId)
	router.GET(baseURL+"/users/:authUserId", si.GetUsersWithAuthUserId)
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
