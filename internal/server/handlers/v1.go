package handlers

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"GAMELIB/internal/actions"
	"GAMELIB/internal/entities"
	"GAMELIB/pkg/web"
	"github.com/forbiddencoding/howlongtobeat"
	"github.com/gin-gonic/gin"
)

func (h *Handler) MainPage(ctx *gin.Context) {
	base, err := actions.GetAllGames(ctx, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	allCount := len(base)

	doneCount, err := actions.GetDoneCountGame(ctx, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"gameList":      base,
		"randomGame":    "",
		"allCountGame":  allCount,
		"doneCountGame": doneCount,
	})
}

func (h *Handler) GetGame(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	game, err := actions.GetGame(ctx, id, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": game,
	})
}

func (h *Handler) GetGameByName(ctx *gin.Context) {
	name := ctx.Param("name")
	game, err := actions.GetGameByName(ctx, name, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": game,
	})
}

func (h *Handler) GetAllGames(ctx *gin.Context) {
	games, err := actions.GetAllGames(ctx, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": games,
	})
}

func (h *Handler) GetRandomGame(ctx *gin.Context) {
	done := ctx.GetBool("done")
	randomGame, err := actions.GetRandomGame(ctx, done, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": randomGame,
	})
}

func (h *Handler) GetRandomByFavorite(ctx *gin.Context) {
	favorite := ctx.GetBool("favorite")
	favorite = !favorite
	randomGame, err := actions.GetFavoriteGame(ctx, favorite, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": randomGame,
	})
}

func (h *Handler) GetRandomListGames(ctx *gin.Context) {
	done := ctx.GetBool("done")
	image := ctx.GetBool("image")

	var err error
	var randomGames []*entities.Game
	if image {
		randomGames, err = actions.GetRandomListGames(ctx, done, h.Storage)
	} else {
		randomGames, err = actions.GetRandomListGamesWithImage(ctx, done, h.Storage)
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": randomGames,
	})
}

func (h *Handler) PostGame(ctx *gin.Context) {
	var err error
	var game *entities.CreateGame
	if err = ctx.ShouldBindJSON(&game); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	clearPathImage := game.ClearPathImage
	var resultHLTB *howlongtobeat.GameDetailSimple

	if game.HowLongToBeatID != 0 {
		resultHLTB, err = actions.GetHltbGame(ctx, game.HowLongToBeatID, h.HLTB)
		if err != nil {
			log.Println(err)
		}
	}

	if game.Image != nil {
		game.Image = actions.ParseLocalImage(*game.Image, clearPathImage)
	} else {
		if game.FindGrid && resultHLTB != nil {
			game.Image = actions.ParseHltbImage(resultHLTB.GameImage)
		} else if game.FindGrid {
			searchGame, err := actions.FindHltbGame(ctx, game.Name, h.HLTB)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
				return
			}

			if searchGame != nil {
				game.Image = actions.ParseHltbImage(searchGame.GameImage)
			}
		}
	}

	if resultHLTB != nil {
		game.HowLongToBeatID = resultHLTB.GameID
		game.HowLongToBeatMainTime = int(math.Round(resultHLTB.CompMain))
		game.HowLongToBeatFullTime = int(math.Round(resultHLTB.CompPlus))
	}

	check, result, err := actions.CreateGame(ctx, game, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	if check {
		ctx.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"data": result,
		})
	}
}

func (h *Handler) PutGame(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	var game *entities.UpdateGame
	if err := ctx.ShouldBindJSON(&game); err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	clearPathImage := game.ClearPathImage
	var resultHLTB *howlongtobeat.GameDetailSimple

	if game.HowLongToBeatID != 0 {
		resultHLTB, err = actions.GetHltbGame(ctx, game.HowLongToBeatID, h.HLTB)
		if err != nil {
			log.Println(err)
		}
	}

	if game.FindGrid {
		if resultHLTB != nil {
			game.Image = actions.ParseHltbImage(resultHLTB.GameImage)
		} else {
			searchGame, err := actions.FindHltbGame(ctx, game.Name, h.HLTB)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
				return
			}

			if searchGame != nil {
				game.Image = actions.ParseHltbImage(searchGame.GameImage)
			}
		}
	} else if game.Image != nil {
		game.Image = actions.ParseLocalImage(*game.Image, clearPathImage)
	}

	if resultHLTB != nil {
		game.HowLongToBeatID = resultHLTB.GameID
		game.HowLongToBeatMainTime = int(math.Round(resultHLTB.CompMain))
		game.HowLongToBeatFullTime = int(math.Round(resultHLTB.CompPlus))
	}

	result, err := actions.PutGame(ctx, id, game, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *Handler) DeleteGame(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	result, err := actions.DeleteGame(ctx, id, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	if result.ImageURL != nil {
		url := strings.Split(*result.ImageURL, "/")
		imgGroup := url[len(url)-2]
		imgName := url[len(url)-1]
		if hltbGroup != imgGroup {
			err = actions.DeleteImage(ctx, imgName, imgGroup, h.Minio.DataBase)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
				return
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h *Handler) ReverseDoneStatus(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	err = actions.ReverseDoneStatus(ctx, id, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) ReverseFavoriteStatus(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	err = actions.ReverseFavoriteStatus(ctx, id, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) GetAllCountGame(ctx *gin.Context) {
	count, err := actions.GetAllCountGame(ctx, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}

func (h *Handler) GetDoneCountGame(ctx *gin.Context) {
	count, err := actions.GetDoneCountGame(ctx, h.Storage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}
