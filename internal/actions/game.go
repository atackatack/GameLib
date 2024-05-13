package actions

import (
	"GAMELIB/internal/entities"
	"GAMELIB/internal/storage/postgres"
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	maxLenGameName = 150
	PathGrids      = "http://localhost:8080/minio/image/"
)

func ParseLocalImage(path string, clearPathImage bool) *string {
	if !clearPathImage {
		result := PathGrids + path
		return &result
	}
	return &path
}

func GetGame(ctx *gin.Context, id int64, storage *postgres.Storage) (*entities.Game, error) {
	return postgres.GetGame(ctx, id, storage.DataBase)
}

func CreateGame(ctx *gin.Context, game *entities.CreateGame, storage *postgres.Storage) (bool, *entities.Game, error) {
	if len(game.Name) == 0 {
		return false, nil, fmt.Errorf("name can't be empty")
	}

	if len(game.Name) > maxLenGameName {
		return false, nil, fmt.Errorf("name of game is too long")
	}

	check, result, err := CheckGameByName(ctx, game.Name, storage)
	if err != nil {
		return false, nil, err
	}

	if check {
		return true, result, nil
	}

	result, err = postgres.CreateGame(ctx, game, storage.DataBase)
	return false, result, err
}

func PutGame(ctx *gin.Context, id int64, update *entities.UpdateGame, storage *postgres.Storage) (*entities.Game, error) {
	return postgres.PutGame(ctx, id, update, storage.DataBase)
}

func DeleteGame(ctx *gin.Context, id int64, storage *postgres.Storage) (*entities.Game, error) {
	game, err := GetGame(ctx, id, storage)
	if err != nil {
		return nil, err
	}

	_ = game

	return postgres.DeleteGame(ctx, id, storage.DataBase)
}

func GetGameByName(ctx *gin.Context, name string, storage *postgres.Storage) (*entities.Game, error) {
	return postgres.GetGameByName(ctx, name, storage.DataBase)
}

func GetAllGames(ctx *gin.Context, storage *postgres.Storage) ([]*entities.Game, error) {
	return postgres.GetAllGames(ctx, storage.DataBase)
}

func GetRandomGame(ctx *gin.Context, done bool, storage *postgres.Storage) (*entities.Game, error) {
	return postgres.GetRandomGame(ctx, done, storage.DataBase)
}

func GetFavoriteGame(ctx *gin.Context, favorite bool, storage *postgres.Storage) (*entities.Game, error) {
	return postgres.GetFavoriteGame(ctx, favorite, storage.DataBase)
}
func GetRandomListGames(ctx *gin.Context, done bool, storage *postgres.Storage) ([]*entities.Game, error) {
	return postgres.GetRandomListGames(ctx, done, storage.DataBase)
}

func GetRandomListGamesWithImage(ctx *gin.Context, done bool, storage *postgres.Storage) ([]*entities.Game, error) {
	return postgres.GetRandomListGamesWithImage(ctx, done, storage.DataBase)
}

func CheckGame(ctx *gin.Context, id int64, storage *postgres.Storage) (bool, *entities.Game, error) {
	return postgres.CheckGame(ctx, id, storage.DataBase)
}

func CheckGameByName(ctx *gin.Context, name string, storage *postgres.Storage) (bool, *entities.Game, error) {
	return postgres.CheckGameByName(ctx, name, storage.DataBase)
}

func ReverseDoneStatus(ctx *gin.Context, id int64, storage *postgres.Storage) error {
	return postgres.ReverseDoneStatus(ctx, id, storage.DataBase)
}

func ReverseFavoriteStatus(ctx *gin.Context, id int64, storage *postgres.Storage) error {
	return postgres.ReverseFavoriteStatus(ctx, id, storage.DataBase)
}

func GetAllCountGame(ctx *gin.Context, storage *postgres.Storage) (int64, error) {
	return postgres.GetAllCountGame(ctx, storage.DataBase)
}

func GetDoneCountGame(ctx *gin.Context, storage *postgres.Storage) (int64, error) {
	return postgres.GetDoneCountGame(ctx, storage.DataBase)
}
