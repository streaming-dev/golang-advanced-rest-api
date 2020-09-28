package auth

import "github.com/labstack/echo"

// Auth Delivery interface
type Handlers interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetUserByID() echo.HandlerFunc
	FindByName() echo.HandlerFunc
	GetUsers() echo.HandlerFunc
}
