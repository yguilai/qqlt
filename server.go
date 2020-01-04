package qqlt

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	Updates chan Update
	Buffer  int
}

type Update struct {
	Type       int32  `json:"Type"`
	TypeCode   string `json:"TypeCode"`
	FromGroup  string `json:"Fromgroup"`
	FromQQ     string `json:"Fromqq"`
	MessageId  string `json:"MessageId"`
	Message    string `json:"Message"`
	CreateTime string `json:"CreateTime"`
	Platform   int32  `json:"Platform"`
	Result     string `json:"Result"`
}

type UpdateResult struct {
	Ec      int32
	ErrCode int32
	Em      string
	Result  interface{}
}

var server *Server

// NewServer 新建一个服务, 接收插件上报信息
func NewServer(buffer int) *Server {
	server = &Server{
		Router:  gin.Default(),
		Updates: make(chan Update, buffer),
	}
	return server
}

func NewDefaultServer() *Server {
	server = &Server{
		Router:  gin.Default(),
		Updates: make(chan Update, 100),
	}

	server.Router.POST("/api/ReceiveMahuaOutput", DefaultServerHandler)
	return server
}

func (s *Server) DefaultRun() {
	err := s.Router.Run(":65321")
	if err != nil {
		panic(err)
	}
}

func (s *Server) Run(addr string) {
	err := s.Router.Run(addr)
	if err != nil {
		panic(err)
	}
}

func DefaultServerHandler(ctx *gin.Context) {
	bytes, err := ctx.GetRawData()
	if err != nil {
		panic(err)
	}
	u := Update{}
	err = json.Unmarshal(bytes, &u)
	if err != nil {
		panic(err)
	}
	server.Updates <- u
}
