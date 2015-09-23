package controllers

import (
	//"fmt"
	"github.com/revel/revel"
	//"golang.org/x/crypto/bcrypt"
	"mitchgottlieb.com/smacktalkgaming/app/models"
	//"mitchgottlieb.com/smacktalkgaming/app/routes"
	//"strings"
	//"encoding/json"
	"encoding/xml"
	"net/http"
)

//structs for BBG grab
type NameXML struct {
	XMLName xml.Name `xml:"name"`
	Value   string   `xml:"value,attr"`
}
type YearpublishedXML struct {
	XMLName xml.Name `xml:"yearpublished"`
	Value   string   `xml:"value,attr"`
}

type Item struct {
	XMLName       xml.Name         `xml:"item"`
	ID            string           `xml:"id,attr"`
	Name          NameXML          `xml:"name"`
	Yearpublished YearpublishedXML `xml:"yearpublished"`
}

type bggXML struct {
	Total   string   `xml:"total,attr"`
	XMLName xml.Name `xml:"items"`
	Items   []Item   `xml:"item"`
}

type Games struct {
	Application
}

func (c Games) Start() revel.Result {
	//qobj := new(models.QueryObj)
	//players := (new(models.QueryObj)).GetAllPlayers()
	return c.Render("fish")
}

func (c Games) ListAutoComplete(auto string) revel.Result {

	playerUUID := c.Session["playerUUID"]

	games := (new(models.QueryObj)).MatchGamesByName(auto, playerUUID)

	return c.RenderJson(games)
}

func (c Games) Add(search string) revel.Result {

	type Listofgames struct {
		Addedgames   []models.Game
		Updatedgames []models.Game
	}

	lg := Listofgames{}
	neo := new(models.Neo4jObj)

	if !CheckAdmin(c.Session["playerUUID"]) {
		return c.Redirect("/errors/admin.html")
	}

	result := new(bggXML)
	//get XML from api2
	getboardgamegeekGames := "http://www.boardgamegeek.com/xmlapi2/search?type=boardgame,boardgameextension&query="

	revel.TRACE.Println("SEED: Searching BGG for Game:", search)
	res, err := http.Get(getboardgamegeekGames + search)
	revel.TRACE.Println("SEED: MEssage from  BGG for Game:", res, err)
	if err != nil {
		revel.ERROR.Fatal("SEED: MEssage from  BGG for Game:", res, err)
		return c.Redirect("/errors/admin.html")
	}

	decoded := xml.NewDecoder(res.Body)

	err = decoded.Decode(result)
	if err != nil {
		revel.ERROR.Fatal("SEED: MEssage from  BGG for Game:", res, err)
		return c.Redirect("/errors/admin.html")
	}
	//http://boardgamegeek.com/boardgame/ID
	revel.TRACE.Println("HERE:", result)
	for _, search := range result.Items {

		var game models.Game

		game.Name = search.Name.Value
		game.Published = search.Yearpublished.Value
		game.BGGLink = "http://boardgamegeek.com/boardgame/" + search.ID

		cargo := neo.Read(&game)
		revel.TRACE.Println("CARGO:", cargo)

		if len(cargo.Data) == 0 {
			neo.Create(&game)

			lg.Addedgames = append(lg.Addedgames, game)

		} else {
			revel.TRACE.Println("game updated:", game)

			lg.Updatedgames = append(lg.Updatedgames, game)

		}
	}
	revel.TRACE.Println("listofgames:", lg)

	return c.RenderJson(lg)
}
