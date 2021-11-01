package middlewares

import "github.com/labstack/echo/v4"

func RountingTracker(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var requestUri = c.Request().RequestURI
		println("request uri ::: " + requestUri)
		return next(c)
	}

}
