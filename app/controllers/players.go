package controllers

import (
	//"fmt"
	"github.com/revel/revel"
	//"github.com/revel/revel/cache"
	//"golang.org/x/crypto/bcrypt"
	"mitchgottlieb.com/smacktalkgaming/app/models"
	//"mitchgottlieb.com/smacktalkgaming/app/routes"
	//"strings"
	"encoding/json"
)

type Players struct {
	Application
}

func (c Players) Status(player string) revel.Result {

	playerret := new(models.QueryObj).GetPlayer(player)

	return c.RenderJson(playerret)
}

func (c Players) OverallStats(playerUUID string) revel.Result {

	//var events []models.Event
	//var playedin []models.Played_In
	//var games []models.Game

	type Row struct {
		Stat  string `json:"stat"`
		Value int    `json:"value"`
	}

	retval := []Row{}

	events, playedins, games := new(models.QueryObj).GetOverallStats(playerUUID)

	if len(events) > 0 && len(playedins) > 0 && len(games) > 0 {
		retval = append(retval, Row{Stat: "Total Events Played", Value: len(events)})
		retval = append(retval, Row{Stat: "Total Games Played", Value: len(games)})

	}

	return c.RenderJson(events)
}

/* TODO: make update player info based on properites versus hardocded */
/********************* only birthdate, nickname, alignment ************/
func (c Players) Update() revel.Result {

	playerdelta := models.Player{}
	query := new(models.QueryObj)
	revel.TRACE.Println("UPDATE PLAYER POST REQUEST")

	revel.TRACE.Println("Body", c.Request.Body)
	revel.TRACE.Println("params", c.Params)

	err := json.NewDecoder(c.Request.Body).Decode(&playerdelta)
	if err != nil {
		//panic(err)
	}
	revel.TRACE.Println("cargocommit:", playerdelta)

	if playerdelta.UUID == c.Session["playerUUID"] {

		currentplayer := query.GetPlayer(playerdelta.UUID)

		if currentplayer.Nickname != playerdelta.Nickname {
			query.SetValue("Player", playerdelta.UUID, "Nickname", playerdelta.Nickname)
		}
		if currentplayer.Birthdate != playerdelta.Birthdate {
			query.SetValue("Player", playerdelta.UUID, "Birthdate", playerdelta.Birthdate)
		}
		if currentplayer.Alignment != playerdelta.Alignment {
			query.SetValue("Player", playerdelta.UUID, "Alignment", playerdelta.Alignment)
		}

	}

	return c.RenderJson(query.GetPlayer(playerdelta.UUID))
}

func (c Players) CreateList(setting string) revel.Result {
	revel.TRACE.Println("HERE", setting)
	c.Validation.Required(setting)
	if c.Validation.HasErrors() {
		// Sets the flash parameter `error` which will be sent by a flash cookie
		c.Flash.Error("Settings invalid!")
		// Keep the validation error from above by setting a flash cookie
		c.Validation.Keep()
		// Copies all given parameters (URL, Form, Multipart) to the flash cookie
		c.FlashParams()
		return c.Redirect("ERROR")
	}

	c.Flash.Success("Settings saved!")

	return c.RenderJson(setting)
}

func (c Players) List() revel.Result {

	qobj := new(models.QueryObj)
	players := qobj.GetAllPlayers()
	/*
		results, err := c.Txn.Select(models.Booking{},
			`select * from Booking where UserId = ?`, c.connected().UserId)
		if err != nil {
			panic(err)
		}

		var bookings []*models.Booking
		for _, r := range results {
			b := r.(*models.Booking)
			bookings = append(bookings, b)
		}
	*/

	return c.Render(players)
}

func (c Players) ListAutoComplete(auto string) revel.Result {
	playerUUID := c.Session["playerUUID"]

	players := (new(models.QueryObj)).MatchPlayersByName(auto, playerUUID)

	return c.RenderJson(players)
}
