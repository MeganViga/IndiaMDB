package api

import (
	"net/http"
	"time"

	db "github.com/MeganViga/IndiaMDB/db/sqlc"
	"github.com/gin-gonic/gin"
)
type createMovieRequestParams struct{
	Name string `json:"name" binding:"required"`
	Summary string `json:"summary" binding:"required"`
	Genre string `json:"genre" binding:"required"`
	Language string `json:"language" binding:"required"`
	ReleaseDate time.Time `json:"releasedate" binding:"required"`
}
func (s *Server)createMovie(ctx *gin.Context){
	var req createMovieRequestParams
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadGateway, errResponse(err))
		return
	}
	arg := db.CreateMovieParams{
		Name: req.Name,
		Summary: req.Summary,
		Genre: req.Genre,
		Language: req.Language,
		ReleaseDate: req.ReleaseDate,
	}
	movie, err := s.store.CreateMovie(ctx,arg)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK,movie)
}