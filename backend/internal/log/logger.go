package log

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

//Config Configuration for the log application
func Config(e *echo.Echo) {

	e.Logger.SetLevel(log.ERROR)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339_nano} ${remote_ip} ${host} ${method} ${uri} ${user_agent} ` +
			`${status} ${error} ${latency_human}` + "\n",
	}))

}
