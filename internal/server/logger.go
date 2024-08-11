package server

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

type RequestLog struct {
	Method     string            `json:"method"`
	URL        string            `json:"url"`
	Proto      string            `json:"proto"`
	RemoteAddr string            `json:"remote_addr"`
	UserAgent  string            `json:"user_agent"`
	Headers    map[string]string `json:"headers"`
	Duration   time.Duration     `json:"duration"`
	Error      string            `json:"error,omitempty"`
	StackTrace string            `json:"stack_trace,omitempty"`
}

func JSONLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			if err := recover(); err != nil {
				var errMsg string
				if errStr, ok := err.(string); ok {
					errMsg = errStr
				} else if errObj, ok := err.(error); ok {
					errMsg = errObj.Error()
				} else {
					errMsg = fmt.Sprintf("%v", err)
				}

				stackTrace := string(debug.Stack())
				logEntry := RequestLog{
					Method:     r.Method,
					URL:        r.RequestURI,
					Proto:      r.Proto,
					RemoteAddr: r.RemoteAddr,
					UserAgent:  r.UserAgent(),
					Headers:    getHeaders(r),
					Duration:   time.Since(start),
					Error:      errMsg,
					StackTrace: stackTrace,
				}

				log.Error().
					Str("method", logEntry.Method).
					Str("url", logEntry.URL).
					Str("proto", logEntry.Proto).
					Str("remote_addr", logEntry.RemoteAddr).
					Str("user_agent", logEntry.UserAgent).
					Dur("duration", logEntry.Duration).
					Str("error", logEntry.Error).
					Str("stack_trace", logEntry.StackTrace).
					Fields(map[string]interface{}{"headers": logEntry.Headers}).
					Msg("Handled request with error")

				http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(rw, r)

		logEntry := RequestLog{
			Method:     r.Method,
			URL:        r.RequestURI,
			Proto:      r.Proto,
			RemoteAddr: r.RemoteAddr,
			UserAgent:  r.UserAgent(),
			Headers:    getHeaders(r),
			Duration:   time.Since(start),
		}

		log.Info().
			Str("method", logEntry.Method).
			Str("url", logEntry.URL).
			Str("proto", logEntry.Proto).
			Str("remote_addr", logEntry.RemoteAddr).
			Str("user_agent", logEntry.UserAgent).
			Dur("duration", logEntry.Duration).
			Fields(map[string]interface{}{"headers": logEntry.Headers}).
			Msg("ok")
	})
}

func getHeaders(r *http.Request) map[string]string {
	headers := make(map[string]string)
	for name, values := range r.Header {
		headers[name] = values[0]
	}
	return headers
}
