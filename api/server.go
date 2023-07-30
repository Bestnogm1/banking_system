package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func NewServer(store *db.Store) *Server {

	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)
	router.POST("/transfer", server.CreateTransfer)
	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/listAcc", server.ListAccounts)
	router.GET("/allAccounts", server.GetAllAccounts)
	router.PUT("/update", server.UpdateAccount)
	router.DELETE("/accounts/:id", server.DeleteOneAccount)
	server.router = router
	return server
}
