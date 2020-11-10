package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	g *gin.Engine
}

func NewServer() *Server {
	s := new(Server)
	s.g = gin.New()

	s.g.GET("keys", s.getKeys)
	s.g.POST("encode", s.encodeMsg)
	s.g.POST("decode", s.decodeMsg)
	s.g.POST("sign", s.signMsg)

	return s
}

func (s *Server) Run() error {
	return s.g.Run()
}
