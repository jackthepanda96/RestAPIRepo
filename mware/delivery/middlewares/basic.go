package middlewares

import "github.com/labstack/echo/v4"

func BasicCheck(username, password string, ctx echo.Context) (bool, error) {
	if username == "admin" && password == "admin" {
		return true, nil
	}

	return false, nil
}
