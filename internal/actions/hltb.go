package actions

import (
	"github.com/forbiddencoding/howlongtobeat"
	"github.com/gin-gonic/gin"
)

const imagePath = "https://howlongtobeat.com/games/"

func ParseHltbImage(path string) *string {
	result := imagePath + path
	return &result
}

func GetHltbGame(ctx *gin.Context, id int, hltb *howlongtobeat.Client) (*howlongtobeat.GameDetailSimple, error) {
	return hltb.DetailSimple(ctx, id)
}

func FindHltbGame(ctx *gin.Context, gameName string, hltb *howlongtobeat.Client) (*howlongtobeat.SearchGameSimple, error) {
	searchResults, err := hltb.SearchSimple(ctx, gameName, howlongtobeat.SearchModifierHideDLC)
	if err != nil {
		return nil, err
	}

	if len(searchResults) > 0 {
		return searchResults[0], nil
	}
	return nil, err
}

func FindHltbGames(ctx *gin.Context, gameName string, hltb *howlongtobeat.Client) ([]*howlongtobeat.SearchGameSimple, error) {
	searchResults, err := hltb.SearchSimple(ctx, gameName, howlongtobeat.SearchModifierHideDLC)
	if err != nil {
		return nil, err
	}

	return searchResults, err
}
