package api

import (
	"fmt"

	db "github.com/MeganViga/IndiaMDB/db/sqlc"
	"github.com/gin-gonic/gin"
)


type Server struct{
	store db.Store
	router *gin.Engine
}


func NewServer(dbs db.Store)(Server, error){
	server :=  Server{
		store: dbs,
	}
	server.setRouter()
	return server, nil
}

func (s *Server)Start(address string)error{
	return s.router.Run(address)
}

func (s *Server)setRouter(){
	router := gin.Default()
	router.POST("/usersignup",s.createUser)
	router.POST("usersignin",s.loginUser)
	router.POST("/createmovie",s.createMovie)
	s.router = router
}

func errResponse(err error)gin.H{
	e := gin.H{"Error": err.Error()}
	fmt.Println(e)
	return e
}