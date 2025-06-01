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

	// # APIルートの設定
	api := e.Group("/api")
	// ## APIルート の version 設定
	v1 := api.Group("/v1")

	// # 非認証ルートの設定
	// ## User 関連のルートグループ
	ug := v1.Group("/user")
	ug.POST("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Create User")
	})
}
