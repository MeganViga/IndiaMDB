package api

import (
	"errors"
	"fmt"
	"net/http"

	db "github.com/MeganViga/IndiaMDB/db/sqlc"
	"github.com/MeganViga/IndiaMDB/util"
	"github.com/gin-gonic/gin"
)

type createUserRequestParams struct{
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type loginUserRequestParams struct{
	Email string `json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}


var ErrUserOrPasswordIncorrect = errors.New("email or password Incorrect")

func (s *Server)createUser(ctx *gin.Context){
	var req  createUserRequestParams
	if err := ctx.BindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadGateway,errResponse(err))
		return
	}
	hashedPassword ,_ := util.HashPassword(req.Password)
	arg := db.CreateUserParams{
		Name: req.Name,
		Email: req.Email,
		Role: "user",
		Passwordhash:hashedPassword ,
	}

	user, err := s.store.CreateUser(ctx,arg)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (s *Server)loginUser(ctx *gin.Context){
	var req loginUserRequestParams

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadGateway,errResponse(err))
		return
	}

	user, err := s.store.GetUserByEmail(ctx,req.Email)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,errResponse(err))
		return
	}

	err = util.CheckPasswordAndHash(user.Passwordhash,req.Password)
	
	if err != nil{
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized,errResponse(ErrUserOrPasswordIncorrect))
		return
	}

	ctx.JSON(http.StatusOK, user)
	
}