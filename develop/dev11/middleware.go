package main

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type responseData struct {
	status int
}

type logResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func (r *logResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

// MakeLoggingHandler возвращает http.HandlerFunc, логирующую запросы
func MakeLoggingHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		start := time.Now()

		responseData := responseData{200}
		lw := logResponseWriter{
			ResponseWriter: w,
			responseData:   &responseData,
		}
		fn(&lw, r)
		duration := time.Since(start)
		logrus.WithFields(logrus.Fields{
			"uri":      r.RequestURI,
			"method":   r.Method,
			"status":   responseData.status,
			"duration": duration,
		}).Info("request completed")
	}
}
