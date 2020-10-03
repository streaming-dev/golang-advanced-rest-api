package utils

import (
	"context"
	"github.com/AleksK1NG/api-mc/config"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

func GetCtxWithReqID(c echo.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Second*5)
	ctx = context.WithValue(ctx, "ReqID", GetRequestID(c))
	return ctx, cancel
}

func ConfigureJWTCookie(cfg *config.Config, jwtToken string) *http.Cookie {
	return &http.Cookie{
		Name:  cfg.Cookie.Name,
		Value: jwtToken,
		Path:  "/",
		// Domain: "/",
		// Expires:    time.Now().Add(1 * time.Minute),
		RawExpires: "",
		MaxAge:     cfg.Cookie.MaxAge,
		Secure:     cfg.Cookie.Secure,
		HttpOnly:   cfg.Cookie.HttpOnly,
		SameSite:   0,
	}
}
