package server

import (
	"GAMELIB/internal/server/handlers"
	"github.com/gin-gonic/gin"
)

func (s *Server) configureRoutes() *gin.Engine {
	router := gin.New()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	h := handlers.NewHandler(s.Storage, s.Minio, s.HLTB)

	router.GET("/", h.MainPage)

	hltbGroup := router.Group("/hltb/")
	{
		hltbGroup.GET("/search/", h.SearchGameByName)
	}

	minioGroup := router.Group("minio")
	{
		minioGroup.POST("/image", h.PostMinioImage)
		minioGroup.GET("/image/:name", h.GetMinioImage)
		minioGroup.DELETE("/image/:name", h.DeleteMinioImage)
	}

	apiGroup := router.Group("/api/")
	{
		v1Group := apiGroup.Group("/v1/")
		{
			gameGroup := v1Group.Group("/game/")
			{
				gameGroup.GET("/:id", h.GetGame)
				gameGroup.GET("/name/:name", h.GetGameByName)

				gameGroup.POST("/", h.PostGame)
				gameGroup.PUT("/:id", h.PutGame)
				gameGroup.DELETE("/:id", h.DeleteGame)

				gameGroup.GET("/all", h.GetAllGames)

				gameGroup.GET("/random", h.GetRandomGame)
				gameGroup.GET("/random/list", h.GetRandomListGames)

				gameGroup.PUT("/reverse/status/:id", h.ReverseDoneStatus)
				gameGroup.PUT("/reverse/favorite/:id", h.ReverseFavoriteStatus)

				gameGroup.GET("/count/all", h.GetAllCountGame)
				gameGroup.GET("/count/done", h.GetDoneCountGame)
			}
		}
	}

	return router
}
