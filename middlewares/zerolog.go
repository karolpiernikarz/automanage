package middlewares

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var AwsLogFile *os.File

func ZeroLogMiddlewareForOrderbox(c *gin.Context) {
	start := time.Now()
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	// Process the request
	c.Next()
	var response string
	// if status is start with 5, log the response
	if c.Writer.Status() >= 500 {
		response = blw.body.String()
	} else {
		response = ""
	}

	// log only the specified parameters
	queriesToLog := "a=" + c.Request.URL.Query().Get("a")
	if c.Request.URL.Query().Get("o") != "" {
		queriesToLog += "&o=" + c.Request.URL.Query().Get("o")
	}
	if c.Request.URL.Query().Get("ak") != "" {
		queriesToLog += "&ak=" + c.Request.URL.Query().Get("ak")
	}
	if c.Request.URL.Query().Get("m") != "" {
		queriesToLog += "&m=" + c.Request.URL.Query().Get("m")
	}

	// Log the request
	log.Info().
		Time("time", start.UTC()).
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.Path).
		Str("query", queriesToLog).
		Int("status", c.Writer.Status()).
		Int("content-length", blw.Size()).
		Str("ip", c.ClientIP()).
		Dur("latency", time.Since(start)).
		Str("response", response).
		Msg("")
}

func ZeroLogMiddleware(c *gin.Context) {
	start := time.Now()
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	// Process the request
	c.Next()
	var response string
	// if status is start with 5, log the response
	if c.Writer.Status() >= 500 {
		response = blw.body.String()
	} else {
		response = ""
	}
	// Log the request
	log.Info().
		Time("time", start.UTC()).
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.Path).
		Int("status", c.Writer.Status()).
		Int("content-length", blw.Size()).
		Str("ip", c.ClientIP()).
		Dur("latency", time.Since(start)).
		Str("response", response).
		Msg("")
}

func ZeroLogMiddlewareForStoreRawJson(c *gin.Context) {
	start := time.Now()
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	logger := zerolog.New(AwsLogFile).With().Timestamp().Logger()

	var message bytes.Buffer
	c.Request.Body = io.NopCloser(io.TeeReader(c.Request.Body, &message))

	c.Next()
	// use a different file for aws

	err := c.Request.Body.Close()
	if err != nil {
		return
	}
	// Log the request
	logger.Info().
		Time("time", start.UTC()).
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.Path).
		Int("status", c.Writer.Status()).
		Int("content-length", blw.Size()).
		Str("ip", c.ClientIP()).
		Dur("latency", time.Since(start)).
		RawJSON("request", message.Bytes()).
		Msg("")
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
