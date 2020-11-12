package api

import (
	db "github.com/Alam3035/simple_bank_golang/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves all HTTP request for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creats a new server to use
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts", server.deleteAccount)
	router.POST("/accounts/update", server.updateAccount)

	server.router = router
	return server
}

// Start starts the server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
