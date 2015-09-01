package controllers

/*
JSON EVENT OBJECT
{
  "event": {
    "numberplayer": "x",
    "start": "date",
    "stop": "date"
  },
  "location": {
    "locationname": "x",
    "locationlng": "date",
    "locationlat": "date"
  },
  "players": [
    {
      "firstname": "",
      "surname": "",
      "nickname": ""
    },
    {
      "firstname": "",
      "surname": "",
      "nickname": ""
    }
  ],
  "playedin": [
    {
      "result": "",
      "place": ""
    },
    {
      "result": "",
      "place": ""
    }
  ],
  "games": [
    {
      "name": "",
      "published": ""
    }
  ]
}



*/

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	//"io/ioutil"
	//"golang.org/x/crypto/bcrypt"
	"mitchgottlieb.com/smacktalkgaming/app/models"
	//"mitchgottlieb.com/smacktalkgaming/app/routes"
	"strconv"
	"strings"
	//"time"
)

type Events struct {
	Application
}

type EventCargo struct {
	Players   []models.Player   `json:players`
	Games     []models.Game     `json:games`
	Locations []models.Location `json:locations`
}

type EventCargoCommit struct {
	Eventuuid string             `json:eventuuid`
	Players   []models.Player    `json:players`
	Playedin  []models.Played_In `json:playedin`
}

//TODO combine above stucts to one EVentCargo
type EventLoad struct {
	Event     models.Event       `json:event`
	Players   []models.Player    `json:players`
	Games     []models.Game      `json:games`
	Locations []models.Location  `json:locations`
	Playedin  []models.Played_In `json:playedin`
}

func (c Events) Commit() revel.Result {

	revel.INFO.Println("COMMIT EVENT POST REQUEST")

	cargo := EventLoad{}
	revel.TRACE.Println("Body", c.Request.Body)
	revel.TRACE.Println("params", c.Params)

	revel.TRACE.Println("cargo:", cargo)

	retEventStatus := make(map[string]string)
	retEventStatus["status"] = "PASS"

	err := json.NewDecoder(c.Request.Body).Decode(&cargo)
	if err != nil {
		//panic(err)
	}
	revel.INFO.Println("cargoload:", cargo)

	//TODO: check if user has done a game at the same time? might be playing two games at once....
	// better to have a system to prove the game happened at time X with results Y
	neo := new(models.Neo4jObj)

	UUIDEvt := neo.Create(&cargo.Event)

	var UUIDnodeLoc string
	for index := range cargo.Locations {
		UUIDnodeLoc = neo.Create(&cargo.Locations[index])
		neo.CreateRelate(UUIDEvt, UUIDnodeLoc, &models.Played_At{})
	}

	var UUIDnodeGame string
	for index := range cargo.Games {
		UUIDnodeGame = neo.Create(&cargo.Games[index])
		neo.CreateRelate(UUIDEvt, UUIDnodeGame, &models.Played_With{})
	}

	var UUIDnodePlayer string
	for index := range cargo.Players {
		UUIDnodePlayer = neo.Create(&cargo.Players[index])

		//the resluts might not be ready for a templated event
		//if len(playedin) > 0 {
		//	neo.CreateRelate(UUIDnodePlayer, UUIDEvt, playedin[index])
		//}
		neo.CreateRelate(UUIDEvt, UUIDnodePlayer, &models.Included{})
		neo.CreateRelate(UUIDnodePlayer, UUIDEvt, &cargo.Playedin[index])
		neo.CreateRelate(UUIDnodePlayer, UUIDEvt, &models.Last_Event{})
	}

	return c.RenderJson(retEventStatus)
}

func (c Events) Status(event string) revel.Result {
	revel.TRACE.Println("STRING", event)

	evtret := new(models.QueryObj).GetEvent(event)

	return c.RenderJson(evtret)
}

func (c Events) LastStatus(playeruuid string) revel.Result {
	revel.TRACE.Println("STRING", playeruuid)

	evtret := new(models.QueryObj).GetLastEvent(playeruuid)

	return c.RenderJson(evtret)
}

func (c Events) Create() revel.Result {
	//qobj := new(models.QueryObj)
	//players := (new(models.QueryObj)).GetAllPlayers()

	//players := (new(models.QueryObj)).MatchPlayersByName("Mi")
	//games := (new(models.QueryObj)).GetAllGames()
	//locations := (new(models.QueryObj)).GetAllEventLocations()

	//revel.TRACE.Println("locations", locations[2])
	playerUUID := c.Session["playerUUID"]
	//return c.Render(players, games, locations)
	return c.Render(playerUUID)
}

func (c Events) List() revel.Result {
	qobj := new(models.QueryObj)
	events := qobj.GetAllEvents()
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

	return c.Render(events)
}

/***** NOT USED *************/

func (c Events) Start() revel.Result {

	//var retval []byte

	revel.TRACE.Println("START EVENT POST REQUEST")

	cargo := EventCargo{}
	revel.TRACE.Println("Body", c.Request.Body)
	revel.TRACE.Println("params", c.Params)

	err := json.NewDecoder(c.Request.Body).Decode(&cargo)
	if err != nil {
		panic(err)
	}
	revel.TRACE.Println("cargo:", cargo)

	retPlayerStatus := make(map[string]string)
	retPlayerStatus["status"] = "PASS"

	query := new(models.QueryObj)

	for k, v := range cargo.Players {
		// check if player is in an event already.. if so return list of users that are in an event
		fmt.Printf("key=%v, value=%v", k, v)
		answer := query.GetPlayerCurrentEvent(v.UUID)
		revel.TRACE.Println("ret getplayercurrentevent", answer)
		if answer != "" {
			revel.TRACE.Println("user:", v.Firstname, " ", v.Surname, " is in an event")
			retPlayerStatus["status"] = "FAIL: players already in events"
			retPlayerStatus[v.UUID] = answer
			//query.SetValues(v)

			//fmt.Print("UUID = ", v.UUID)
		}
	}
	if strings.LastIndex(retPlayerStatus["status"], "FAIL") == 0 {
		//retval, _ = json.Marshal(retPlayerStatus)
		return c.RenderJson(retPlayerStatus)
	}

	//users are clear, lets get the event going.
	retPlayerStatus["UUIDEvt"] = recordStartOfEvent(
		models.Event{
			Numplayers: strconv.Itoa(len(cargo.Players)),
		},
		cargo.Players,
		cargo.Games,
		cargo.Locations,
	)

	//log users in currentevent
	//TODO: use SetValue to set aech user currently inveloved with evt X

	//return the event UUID
	return c.RenderJson(retPlayerStatus)
}

func recordStartOfEvent(
	evt models.Event,
	players []models.Player,
	games []models.Game,
	locations []models.Location) string {

	neo := new(models.Neo4jObj)

	UUIDEvt := neo.Create(&evt)

	var UUIDnodeLoc string
	for index := range locations {
		UUIDnodeLoc = neo.Create(&locations[index])
		neo.CreateRelate(UUIDEvt, UUIDnodeLoc, &models.Played_At{})
	}

	var UUIDnodeGame string
	for index := range games {
		UUIDnodeGame = neo.Create(&games[index])
		neo.CreateRelate(UUIDEvt, UUIDnodeGame, &models.Played_With{})
	}

	var UUIDnodePlayer string
	for index := range players {
		UUIDnodePlayer = neo.Create(&players[index])

		//the resluts might not be ready for a templated event
		//if len(playedin) > 0 {
		//	neo.CreateRelate(UUIDnodePlayer, UUIDEvt, playedin[index])
		//}
		neo.CreateRelate(UUIDEvt, UUIDnodePlayer, &models.Included{})
	}

	return UUIDEvt
}
