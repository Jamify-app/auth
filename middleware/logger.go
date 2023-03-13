package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func StructuredLogger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		method := ctx.Request.Method

		ctx.Next()

		params := gin.LogFormatterParams{}

		params.TimeStamp = start
		params.Path = path
		params.Method = method
		params.StatusCode = ctx.Writer.Status()
		params.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()

		var logEvent *zerolog.Event
		if params.StatusCode >= http.StatusInternalServerError {
			logEvent = logger.Error()
		} else if params.StatusCode >= http.StatusBadRequest {
			logEvent = logger.Warn()
		} else {
			logEvent = logger.Info()
		}

		logEvent.
			Str("time", params.TimeStamp.String()).
			Int("status_code", params.StatusCode).
			Str("method", params.Method).
			Str("path", params.Path).
			Msg(params.ErrorMessage)
	}
}
