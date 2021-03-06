package net

import "fmt"
import "time"
import "net/http"

import "github.com/labstack/gommon/log"

type HandlerFunc func(*RequestRuntime) error
type MiddlewareFunc func(HandlerFunc) HandlerFunc

type Server struct {
	*http.Server
	Runtime *ServerRuntime
}

func (server *Server) Run(host string) {
	server.Server = &http.Server{
		Addr: host,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.Handler = server.Runtime

	server.Logger().Debugf(fmt.Sprintf("binding to host[%s]", host))
	server.ListenAndServe()
}

func (server *Server) Logger() *log.Logger {
	return server.Runtime.Log
}

