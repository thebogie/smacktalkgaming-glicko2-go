package controllers

import (
	//"fmt"
	"github.com/revel/revel"
	//"golang.org/x/crypto/bcrypt"
	"mitchgottlieb.com/smacktalkgaming/app/models"
	//"mitchgottlieb.com/smacktalkgaming/app/routes"
	//"strings"
)

type Games struct {
	Application
}

func (c Games) Start() revel.Result {
	//qobj := new(models.QueryObj)
	players := (new(models.QueryObj)).GetAllPlayers()
	return c.Render(players)
}

func (c Games) ListAutoComplete(auto string) revel.Result {

	playerUUID := c.Session["playerUUID"]

	games := (new(models.QueryObj)).MatchGamesByName(auto, playerUUID)

	return c.RenderJson(games)
}
