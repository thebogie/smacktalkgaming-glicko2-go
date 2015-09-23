package controllers

import (
	//"encoding/json"
	//"fmt"
	"github.com/revel/revel"
	//"github.com/revel/revel/cache"
	//"golang.org/x/oauth2"
	"mitchgottlieb.com/smacktalkgaming/app/models"
	//"mitchgottlieb.com/smacktalkgaming/app/routes"
	//"net/http"
	//"net/url"
	//"strconv"
	//"time"
)

type Stats struct {
	Application
}

//GrabNemesis - top 3 competitors that player has most losses and
type GrabNemesisCargo struct {
	Players []models.Player
}

func (c Stats) Grab(playeruuid, action string) revel.Result {
	revel.TRACE.Println("STRING", action, playeruuid)

	switch action {
	case "Nemesis":
		revel.TRACE.Println("Grab Nemesis")
		return c.RenderJson(GrabNemesis())
	default:
		revel.TRACE.Println("No Action")

	}
	return c.RenderJson("")
}

func GrabNemesis() (retval GrabNemesisCargo) {

	player := new(models.Player)
	player.UUID = "82130948203948203984"
	retval.Players = append(retval.Players, *player)

	return retval
}
