package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yamaga-shu/jwt-go-next/api-service/app/middleware"
)

func Set(e *echo.Echo) {
	// middleware の設定
	// # ログの設定
	middleware.Logger(e)

	// ルートの設定
	// # ヘルスチェック
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health Check")
	})
}
