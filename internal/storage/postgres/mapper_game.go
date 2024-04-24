package postgres

import (
	"GAMELIB/internal/entities"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"strings"
)

const (
	GamesTabler = "gamelib.t_games"

	IDCol       = "id"
	DoneCol     = "done"
	NameCol     = "name"
	FavoriteCol = "Favorite"

	ImageUrlCol = "image_url"

	HowLongToBeatIDCol       = "hltb_id"
	HowLongToBeatMainTimeCol = "hltb_main_time"
	HowLongToBeatFullTimeCol = "hltb_full_time"

	CreateDTCol = "create_dt"
	UpdateDTCol = "update_dt"
)

var GamesMainCols = []string{NameCol, DoneCol}
var GamesBaseCols = []string{NameCol, DoneCol, FavoriteCol, ImageUrlCol, HowLongToBeatIDCol, HowLongToBeatMainTimeCol, HowLongToBeatFullTimeCol}
var GamesAllCols = append([]string{IDCol, CreateDTCol, UpdateDTCol}, GamesBaseCols...)

func GetGame(_ *gin.Context, id int64, repo *sqlx.DB) (*entities.Game, error) {
	rows, err := sq.Select(GamesAllCols...).
		From(GamesTabler).
		Where(sq.Eq{IDCol: id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var game entities.Game
	for rows.Next() {
		if err := rows.Scan(
			&game.ID,
			&game.CreateDt,
			&game.UpdateDt,
			&game.Name,
			&game.Done,
			&game.Favorite,
			&game.ImageURL,
			&game.HowLongToBeatID,
			&game.HowLongToBeatMainTime,
			&game.HowLongToBeatFullTime,
		); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &game, nil
}

func CreateGame(_ *gin.Context, createGame *entities.CreateGame, repo *sqlx.DB) (*entities.Game, error) {
	rows, err := sq.Insert(GamesTabler).
		Columns(GamesBaseCols...).
		Values(createGame.Name, createGame.Done, false, createGame.Image,
			createGame.HowLongToBeatID, createGame.HowLongToBeatMainTime, createGame.HowLongToBeatFullTime).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(GamesAllCols, ","))).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()

	if err != nil {
		return nil, err
	}

	var result entities.Game
	for rows.Next() {
		if err := rows.Scan(
			&result.ID,
			&result.CreateDt,
			&result.UpdateDt,
			&result.Name,
			&result.Done,
			&result.Favorite,
			&result.ImageURL,
			&result.HowLongToBeatID,
			&result.HowLongToBeatMainTime,
			&result.HowLongToBeatFullTime,
		); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func PutGame(_ *gin.Context, id int64, updateGame *entities.UpdateGame, repo *sqlx.DB) (*entities.Game, error) {
	query := sq.Update(GamesTabler).
		Where(sq.Eq{IDCol: id}).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(GamesAllCols, ",")))

	if updateGame.Done != nil {
		query = query.Set(DoneCol, *updateGame.Done)
	}

	if len(updateGame.Name) > 0 {
		query = query.Set(NameCol, updateGame.Name)
	}

	if updateGame.Image != nil {
		query = query.Set(ImageUrlCol, updateGame.Image)
	}

	if updateGame.HowLongToBeatID != 0 {
		query = query.Set(HowLongToBeatIDCol, updateGame.HowLongToBeatID)
		query = query.Set(HowLongToBeatMainTimeCol, updateGame.HowLongToBeatMainTime)
		query = query.Set(HowLongToBeatFullTimeCol, updateGame.HowLongToBeatFullTime)
	}

	rows, err := query.PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	if err != nil {
		return nil, err
	}

	var result entities.Game
	for rows.Next() {
		if err := rows.Scan(
			&result.ID,
			&result.CreateDt,
			&result.UpdateDt,
			&result.Name,
			&result.Done,
			&result.Favorite,
			&result.ImageURL,
			&result.HowLongToBeatID,
			&result.HowLongToBeatMainTime,
			&result.HowLongToBeatFullTime,
		); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func DeleteGame(_ *gin.Context, id int64, repo *sqlx.DB) (*entities.Game, error) {
	query := sq.Delete(GamesTabler).
		Where(sq.Eq{IDCol: id}).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(GamesAllCols, ",")))

	rows, err := query.PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	if err != nil {
		return nil, err
	}

	var result entities.Game
	for rows.Next() {
		if err := rows.Scan(
			&result.ID,
			&result.CreateDt,
			&result.UpdateDt,
			&result.Name,
			&result.Done,
			&result.Favorite,
			&result.ImageURL,
			&result.HowLongToBeatID,
			&result.HowLongToBeatMainTime,
			&result.HowLongToBeatFullTime,
		); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

// ============================================================================================================

func GetGameByName(_ *gin.Context, name string, repo *sqlx.DB) (*entities.Game, error) {
	rows, err := sq.Select(GamesAllCols...).
		From(GamesTabler).
		Where(sq.Eq{NameCol: name}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var game entities.Game
	for rows.Next() {
		if err := rows.Scan(
			&game.ID,
			&game.CreateDt,
			&game.UpdateDt,
			&game.Name,
			&game.Done,
			&game.Favorite,
			&game.ImageURL,
			&game.HowLongToBeatID,
			&game.HowLongToBeatMainTime,
			&game.HowLongToBeatFullTime,
		); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &game, nil
}

func GetAllGames(_ *gin.Context, repo *sqlx.DB) ([]*entities.Game, error) {
	rows, err := sq.Select(GamesAllCols...).
		From(GamesTabler).
		OrderBy(fmt.Sprintf("%s %s", FavoriteCol, "DESC"), DoneCol, fmt.Sprintf("LOWER(%v)", NameCol)).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var gameList []*entities.Game
	for rows.Next() {
		var tmp entities.Game
		if err := rows.Scan(
			&tmp.ID,
			&tmp.CreateDt,
			&tmp.UpdateDt,
			&tmp.Name,
			&tmp.Done,
			&tmp.Favorite,
			&tmp.ImageURL,
			&tmp.HowLongToBeatID,
			&tmp.HowLongToBeatMainTime,
			&tmp.HowLongToBeatFullTime,
		); err != nil {
			return gameList, err
		}
		gameList = append(gameList, &tmp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return gameList, nil
}

func GetRandomGame(_ *gin.Context, done bool, repo *sqlx.DB) (*entities.Game, error) {
	rows, err := sq.Select(GamesMainCols...).
		From(GamesTabler).
		Where(sq.Eq{DoneCol: done}).
		OrderBy("random()").
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var game entities.Game
	for rows.Next() {
		if err := rows.Scan(
			&game.Name,
			&game.Done,
		); err != nil {
			return nil, err
		}
	}

	return &game, nil
}

func GetRandomListGames(_ *gin.Context, done bool, repo *sqlx.DB) ([]*entities.Game, error) {
	rows, err := sq.Select(GamesMainCols...).
		From(GamesTabler).
		Where(sq.Eq{DoneCol: done}).
		OrderBy("random()").
		Limit(30).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var gameList []*entities.Game
	for rows.Next() {
		var tmp entities.Game
		if err := rows.Scan(&tmp.Name, &tmp.Done); err != nil {
			return gameList, err
		}
		gameList = append(gameList, &tmp)
	}

	if err = rows.Err(); err != nil {
		return gameList, err
	}
	return gameList, nil
}

func GetRandomListGamesWithImage(_ *gin.Context, done bool, repo *sqlx.DB) ([]*entities.Game, error) {
	rows, err := sq.Select(NameCol, DoneCol, ImageUrlCol).
		From(GamesTabler).
		Where(sq.Eq{DoneCol: done}).
		OrderBy("random()").
		Limit(30).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var gameList []*entities.Game
	for rows.Next() {
		var tmp entities.Game
		if err := rows.Scan(&tmp.Name, &tmp.Done, &tmp.ImageURL); err != nil {
			return gameList, err
		}
		gameList = append(gameList, &tmp)
	}

	if err = rows.Err(); err != nil {
		return gameList, err
	}
	return gameList, nil
}

func CheckGame(_ *gin.Context, id int64, repo *sqlx.DB) (bool, *entities.Game, error) {
	rows, err := sq.Select(GamesMainCols...).
		From(GamesTabler).
		Where(sq.Eq{IDCol: id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return false, nil, err
	}

	check := false
	var game entities.Game
	for rows.Next() {
		check = true
		if err := rows.Scan(
			&game.Name,
			&game.Done,
		); err != nil {
			return false, nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return false, nil, err
	}

	return check, &game, nil
}

func CheckGameByName(_ *gin.Context, name string, repo *sqlx.DB) (bool, *entities.Game, error) {
	rows, err := sq.Select(GamesMainCols...).
		From(GamesTabler).
		Where(sq.Eq{NameCol: name}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return false, nil, err
	}

	check := false
	var game entities.Game
	for rows.Next() {
		check = true
		if err := rows.Scan(
			&game.Name,
			&game.Done,
		); err != nil {
			return false, nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return false, nil, err
	}

	return check, &game, nil
}

func ReverseDoneStatus(_ *gin.Context, id int64, repo *sqlx.DB) error {
	_, err := repo.DB.Exec("UPDATE gamelib.t_games SET DONE = NOT DONE WHERE id = $1 "+
		fmt.Sprintf("RETURNING %s", strings.Join(GamesAllCols, ",")), id)

	if err != nil {
		return err
	}
	return nil
}

func ReverseFavoriteStatus(_ *gin.Context, id int64, repo *sqlx.DB) error {
	_, err := repo.DB.Exec("UPDATE gamelib.t_games SET FAVORITE = NOT FAVORITE WHERE id = $1 "+
		fmt.Sprintf("RETURNING %s", strings.Join(GamesAllCols, ",")), id)

	if err != nil {
		return err
	}
	return nil
}

func GetAllCountGame(_ *gin.Context, repo *sqlx.DB) (int64, error) {
	rows, err := sq.Select("COUNT(*)").
		From(GamesTabler).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return 0, err
	}

	var cnt int64
	for rows.Next() {
		if err := rows.Scan(&cnt); err != nil {
			return 0, err
		}
	}

	return cnt, nil
}

func GetDoneCountGame(_ *gin.Context, repo *sqlx.DB) (int64, error) {
	rows, err := sq.Select("COUNT(*)").
		From(GamesTabler).
		Where(sq.Eq{DoneCol: true}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return 0, err
	}

	var cnt int64
	for rows.Next() {
		if err := rows.Scan(&cnt); err != nil {
			return 0, err
		}
	}

	return cnt, nil
}
