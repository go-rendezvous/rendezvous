package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func LogrusLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Log incoming request
		logrus.Debugf("Incoming Request: %s %s", c.Request().Method, c.Request().URL.String())

		// Call next handler
		if err := next(c); err != nil {
			c.Error(err)
		}

		logrus.Debugf("Outgoing Response: %d %s", c.Response().Status, http.StatusText(c.Response().Status))

		return nil
	}
}
