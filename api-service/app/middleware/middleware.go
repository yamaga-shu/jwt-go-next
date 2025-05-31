package middleware

import (
	"slices"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Set(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			skipPaths := []string{
				"/favicon.ico",
				"/.well-known/appspecific/com.chrome.devtools.json",
			}
			return slices.Contains(skipPaths, c.Request().URL.Path)
		},
		Format: `{
	"time":"${time_rfc3339}",
	"remote_ip":"${remote_ip}",
	"uri":"${uri}",
	"status":"${status}",
	"error":"${error}",
	"headers": {
		"Authorization": "${header:Authorization}"
	}
}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
}
