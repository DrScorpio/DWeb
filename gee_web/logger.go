package gee_web

import (
	"github.com/sirupsen/logrus"
	"time"
)

var DefaultLogger *logrus.Logger

func init() {
	DefaultLogger = logrus.New()
	DefaultLogger.SetFormatter(&logrus.JSONFormatter{})
}

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		DefaultLogger.Infof("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
