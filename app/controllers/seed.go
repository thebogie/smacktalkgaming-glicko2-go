package controllers

import (
	//"bytes"
	//"encoding/gob"
	"encoding/json"
	//"github.com/jmcvetta/neoism"
	"github.com/revel/revel"
	"io/ioutil"
	"log"
	"mitchgottlieb.com/smacktalkgaming/app/models"
	//"reflect"
	//"fmt"
	"encoding/xml"
	"net/http"
	"regexp"
	"time"
	"unicode"
)

type Seed struct {
	Application
}

type Config struct {
	PlayedList []PlayedElement
}

type PlayedElement struct {
	Game   models.Game
	Player models.Player
	//Played models.Played
}

/*
type Base struct {
	Query   string `json:"query"`
	Count   int    `json:"count"`
	Objects []struct {
		ItemID      string `json:"ITEM_ID"`
		ProdClassID string `json:"PROD_CLASS_ID"`
		Available   int    `json:"AVAILABLE"`
	}
}
*/

type SeedObj struct {
	Games   []models.Game
	Players []models.Player
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (c Seed) Index(kind string) revel.Result {
	log.Println("SEED", kind)

	devmode, _ := revel.Config.String("mode.dev")

	if devmode == "false" {
		return c.Redirect("/")
	}
	neo := new(models.Neo4jObj)

	if kind == "games" {

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

		result := new(bggXML)
		//get XML from api2
		getboardgamegeekGames := "http://www.boardgamegeek.com/xmlapi2/search?type=boardgame,boardgameextension&query="

		//TODO: how to get all the games? Run through alphabet?
		search := "a"

		res, err := http.Get(getboardgamegeekGames + search)
		if err == nil {
			log.Println("GAME GET", res)
			decoded := xml.NewDecoder(res.Body)

			err := decoded.Decode(result)
			if err != nil {
				log.Printf("Error: %v\n", err)
			}
			//http://boardgamegeek.com/boardgame/ID
			log.Println("HERE:", result)
			for _, value := range result.Items {

				var game models.Game

				game.Name = value.Name.Value
				game.Published = value.Yearpublished.Value
				game.BGGLink = "http://boardgamegeek.com/boardgame/" + value.ID

				UUIDnodeGame := neo.Create(&game)
				log.Println("game added:", UUIDnodeGame)
			}
		}
	}

	if kind == "players" {

		type name struct {
			First string `json:"first"`
			Last  string `json:"last"`
		}

		type user struct {
			Username string `json:"username"`
			NameObj  name   `json:"name"`
		}

		type results struct {
			User user `json:"user"`
		}

		type randomnameObj struct {
			Results []results
		}

		sum := 0
		for i := 0; i < 1000; i++ {
			var randomname randomnameObj

			//adjective!
			getnameURL := "http://api.randomuser.me/"

			res, err := http.Get(getnameURL)
			if err == nil {
				jsonDataFromHttp, err := ioutil.ReadAll(res.Body)
				if err != nil {
					panic(err)
				}
				log.Println("Random Name", string(jsonDataFromHttp))

				err = json.Unmarshal([]byte(jsonDataFromHttp), &randomname) // here!

				if err != nil {
					panic(err)
				}

				log.Println("Random Name after", randomname)

				var player models.Player

				capFirst := []rune(randomname.Results[0].User.NameObj.First)
				capFirst[0] = unicode.ToUpper(capFirst[0])

				capLast := []rune(randomname.Results[0].User.NameObj.Last)
				capLast[0] = unicode.ToUpper(capLast[0])

				player.Firstname = string(capFirst)
				player.Surname = string(capLast)
				player.Nickname = randomname.Results[0].User.Username

				log.Println("player", player)

				UUIDnodePlayer := neo.Create(&player)
				log.Println("player", UUIDnodePlayer)

			}
			sum += i
		}

	}

	if kind == "events" {

		seedaction(neo,
			&models.Event{
				Numplayers: "3",
				Start:      "201409122000",
				Stop:       "201409122100",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Payday", Published: "1973"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201409131000",
				Stop:  "201409131100",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "DROP", Place: ""},
			},
			[]*models.Game{
				&models.Game{Name: "London Cabbie Game", Published: "1971"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201409131400",
				Stop:  "201409131500",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Payday", Published: "1973"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201409131600",
				Stop:  "201409131700",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
				&models.Player{Firstname: "Penolope", Surname: ""},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "King of Toyko", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201409131800",
				Stop:  "201409131900",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
				&models.Player{Firstname: "Penolope", Surname: ""},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "King of Toyko", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201409271800",
				Stop:  "201409271900",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Takenoko", Published: "2012"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201409272000",
				Stop:  "201409272100",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
				&models.Player{Firstname: "Miguel", Surname: "Coronado"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "5",

				Start: "201409272100",
				Stop:  "201409272200",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Miguel", Surname: "Coronado"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
				&models.Player{Firstname: "Mario", Surname: "Munoz"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "DEMOLISH", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
				&models.Played_In{Result: "LOST", Place: "5"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201409272300",
				Stop:  "201409280000",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Miguel", Surname: "Coronado"},

				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "DEMOLISH", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201409280100",
				Stop:  "201409280200",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Miguel", Surname: "Coronado"},
				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201409281200",
				Stop:  "201409281300",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Tokaido", Published: "2012"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201409281700",
				Stop:  "201409281800",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "FRAG", Published: "2012"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201410092000",
				Stop:  "201410092200",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Netrunner", Published: "2012"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201410100100",
				Stop:  "201410100300",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Netrunner", Published: "2012"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201410131030",
				Stop:  "201410131130",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Sleeping Queen", Published: "2005"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201410142130",
				Stop:  "201410142230",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
				&models.Player{Firstname: "Gary", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201410142230",
				Stop:  "201410142330",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{

				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Gary", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201410142330",
				Stop:  "201410150030",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Frederique", Surname: "McAdam"},

				&models.Player{Firstname: "Gary", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201410162230",
				Stop:  "201410162330",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Netrunner", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201410172030",
				Stop:  "201410172130",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Miguel", Surname: "Coronado"},
				&models.Player{Firstname: "Mario", Surname: "Munoz"},

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "DEMOLISH", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201410172130",
				Stop:  "201410172230",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Miguel", Surname: "Coronado"},
				&models.Player{Firstname: "Mario", Surname: "Munoz"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201410172330",
				Stop:  "201410180130",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{

				&models.Player{Firstname: "Josephine", Surname: "McAdam"},

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Miguel", Surname: "Coronado"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201410261200",
				Stop:  "201410261300",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "DEMOLISH", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Carcassonne", Published: "2000"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201411081500",
				Stop:  "201411081600",
			},
			&models.Location{
				Locationname: "Emerald Tavern Games and Cafe, Research Boulevard, Austin, TX, United States",
				Locationlat:  "30.371085",
				Locationlng:  "-97.72470299999998",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Takenoko", Published: "2012"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201411081600",
				Stop:  "201411081700",
			},
			&models.Location{
				Locationname: "Emerald Tavern Games and Cafe, Research Boulevard, Austin, TX, United States",
				Locationlat:  "30.371085",
				Locationlng:  "-97.72470299999998",
			},
			[]*models.Player{
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Carcassonne", Published: "2000"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201411111400",
				Stop:  "201411111500",
			},
			&models.Location{
				Locationname: "Austin Java, Barton Springs Road, Austin, TX, United States",
				Locationlat:  "30.262408",
				Locationlng:  "-97.761731",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Lost Cities", Published: "1999"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201411111800",
				Stop:  "201411111900",
			},
			&models.Location{
				Locationname: "Austin Java, Barton Springs Road, Austin, TX, United States",
				Locationlat:  "30.262408",
				Locationlng:  "-97.761731",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Lost Cities", Published: "1999"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201411111900",
				Stop:  "201411112000",
			},
			&models.Location{
				Locationname: "Austin Java, Barton Springs Road, Austin, TX, United States",
				Locationlat:  "30.262408",
				Locationlng:  "-97.761731",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Lost Cities", Published: "1999"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201411271900",
				Stop:  "201411272000",
			},

			&models.Location{
				Locationname: "2505 Side Cove, Austin, TX, United States",
				Locationlat:  "30.251579",
				Locationlng:  "-97.79226399999999",
			},
			[]*models.Player{
				&models.Player{Firstname: "Simone", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Carcassonne", Published: "2000"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201411301800",
				Stop:  "201411301900",
			},
			&models.Location{
				Locationname: "2505 Side Cove, Austin, TX, United States",
				Locationlat:  "30.251579",
				Locationlng:  "-97.79226399999999",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Carcassonne", Published: "2000"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201411301900",
				Stop:  "201411302000",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "WON", Place: "1"},
			},
			[]*models.Game{
				&models.Game{Name: "Forbidden Island", Published: "2012"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201412061600",
				Stop:  "201412061700",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Tonga Island", Published: "2009"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201411111800",
				Stop:  "201411111900",
			},
			&models.Location{
				Locationname: "Austin Java, Barton Springs Road, Austin, TX, United States",
				Locationlat:  "30.262408",
				Locationlng:  "-97.761731",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Caesar & Cleopatra", Published: "1997"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201411111900",
				Stop:  "201411112000",
			},
			&models.Location{
				Locationname: "Austin Java, Barton Springs Road, Austin, TX, United States",
				Locationlat:  "30.262408",
				Locationlng:  "-97.761731",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Caesar & Cleopatra", Published: "1997"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201411261600",
				Stop:  "201411261700",
			},
			&models.Location{
				Locationname: "Austin Java, Barton Springs Road, Austin, TX, United States",
				Locationlat:  "30.262408",
				Locationlng:  "-97.761731",
			},
			[]*models.Player{
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Hot Tin Roof", Published: "2014"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201501301900",
				Stop:  "201501302000",
			},
			&models.Location{
				Locationname: "2505 Side Cove, Austin, TX, United States",
				Locationlat:  "30.251579",
				Locationlng:  "-97.79226399999999",
			},
			[]*models.Player{
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Carcassonne", Published: "2000"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201502152000",
				Stop:  "201502152010",
			},
			&models.Location{
				Locationname: "Radio Coffee & Beer, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.231362",
				Locationlng:  "-97.78786200000002",
			},
			[]*models.Player{
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Tsuro", Published: "2004"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201502152010",
				Stop:  "201502152020",
			},
			&models.Location{
				Locationname: "Radio Coffee & Beer, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.231362",
				Locationlng:  "-97.78786200000002",
			},
			[]*models.Player{
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Tsuro", Published: "2004"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201502152020",
				Stop:  "201502152030",
			},
			&models.Location{
				Locationname: "Radio Coffee & Beer, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.231362",
				Locationlng:  "-97.78786200000002",
			},
			[]*models.Player{
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Tsuro", Published: "2004"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201502152030",
				Stop:  "201502152040",
			},
			&models.Location{
				Locationname: "Radio Coffee & Beer, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.231362",
				Locationlng:  "-97.78786200000002",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Tsuro", Published: "2004"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201502160900",
				Stop:  "201502161000",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},

			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},

				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Small World", Published: "2009"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201503011500",
				Stop:  "201503011530",
			},
			&models.Location{
				Locationname: "2505 Side Cove, Austin, TX, United States",
				Locationlat:  "30.251579",
				Locationlng:  "-97.79226399999999",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},

				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "The Builders: Middle Ages", Published: "2013"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201503011700",
				Stop:  "201503011730",
			},

			&models.Location{
				Locationname: "2505 Side Cove, Austin, TX, United States",
				Locationlat:  "30.251579",
				Locationlng:  "-97.79226399999999",
			},
			[]*models.Player{
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "The Builders: Middle Ages", Published: "2013"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201503071400",
				Stop:  "201503071500",
			},
			&models.Location{
				Locationname: "Radio Coffee & Beer, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.231362",
				Locationlng:  "-97.78786200000002",
			},
			[]*models.Player{
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "The Builders: Middle Ages", Published: "2013"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201503071500",
				Stop:  "201503071600",
			},
			&models.Location{
				Locationname: "Radio Coffee & Beer, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.231362",
				Locationlng:  "-97.78786200000002",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "The Builders: Middle Ages", Published: "2013"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201503071600",
				Stop:  "201503071700",
			},
			&models.Location{
				Locationname: "Radio Coffee & Beer, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.231362",
				Locationlng:  "-97.78786200000002",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "The Builders: Middle Ages", Published: "2013"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201503081900",
				Stop:  "201503082000",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Gary", Surname: "McAdam"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201503082200",
				Stop:  "201503082300",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Gary", Surname: "McAdam"},

				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201503082100",
				Stop:  "201503082200",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Gary", Surname: "McAdam"},

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201503132300",
				Stop:  "201503140001",
			},

			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},

				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "The Builders: Middle Ages", Published: "2013"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201503140130",
				Stop:  "201503140200",
			},

			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "The Builders: Middle Ages", Published: "2013"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201503140030",
				Stop:  "201503140100",
			},

			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "The Builders: Middle Ages", Published: "2013"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201503172100",
				Stop:  "201503172200",
			},

			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Council of Verona", Published: "2013"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201503172300",
				Stop:  "201503172330",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Council of Verona", Published: "2013"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201503172200",
				Stop:  "201503172300",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Council of Verona", Published: "2013"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504041500",
				Stop:  "201504041600",
			},
			&models.Location{
				Locationname: "Emerald Tavern Games and Cafe, Research Boulevard, Austin, TX, United States",
				Locationlat:  "30.371085",
				Locationlng:  "-97.72470299999998",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Matthew", Surname: "Curley"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Small World Underground", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504042300",
				Stop:  "201504050030",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Matthew", Surname: "Curley"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Small World", Published: "2009"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504082000",
				Stop:  "201504082030",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Dungeoneer: Tomb of the Lich Lord", Published: "2003"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504082200",
				Stop:  "201504082300",
			},
			&models.Location{
				Locationname: "Radio Coffee & Beer, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.231362",
				Locationlng:  "-97.78786200000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Bohnanza", Published: "1997"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504082300",
				Stop:  "201504082345",
			},
			&models.Location{
				Locationname: "Radio Coffee & Beer, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.231362",
				Locationlng:  "-97.78786200000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Bohnanza", Published: "1997"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504112000",
				Stop:  "201504112100",
			},

			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{

				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Small World", Published: "2009"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201504121430",
				Stop:  "201504121530",
			},
			&models.Location{
				Locationname: "Emerald Tavern Games and Cafe, Research Boulevard, Austin, TX, United States",
				Locationlat:  "30.371085",
				Locationlng:  "-97.72470299999998",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Survive: Escape from Atlantis!", Published: "1982"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504131930",
				Stop:  "201504132030",
			},
			&models.Location{
				Locationname: "Radio Coffee & Beer, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.231362",
				Locationlng:  "-97.78786200000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Bohnanza", Published: "1997"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504132200",
				Stop:  "201504132230",
			},
			&models.Location{
				Locationname: "Radio Coffee & Beer, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.231362",
				Locationlng:  "-97.78786200000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Bohnanza", Published: "1997"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504132030",
				Stop:  "201504132130",
			},
			&models.Location{
				Locationname: "Radio Coffee & Beer, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.231362",
				Locationlng:  "-97.78786200000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Bohnanza", Published: "1997"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504242000",
				Stop:  "201504242030",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Bohnanza", Published: "1997"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504242030",
				Stop:  "201504242100",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Bohnanza", Published: "1997"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504242130",
				Stop:  "201504242200",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Bohnanza", Published: "1997"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504242300",
				Stop:  "201504242330",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Bohnanza", Published: "1997"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504242200",
				Stop:  "201504242230",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Bohnanza", Published: "1997"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504302000",
				Stop:  "201504302030",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "DEMOLISH", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Bohnanza", Published: "1997"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504302100",
				Stop:  "201504302200",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Mr. Jack Pocket", Published: "2010"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504302200",
				Stop:  "201504302300",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Mr. Jack Pocket", Published: "2010"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201504302300",
				Stop:  "201504302330",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Mr. Jack Pocket", Published: "2010"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201505061930",
				Stop:  "201505062000",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{

				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "The Builders: Middle Ages", Published: "2013"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201505062330",
				Stop:  "201505062345",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Mr. Jack Pocket", Published: "2010"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201505062300",
				Stop:  "201505062314",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Mr. Jack Pocket", Published: "2010"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201505062230",
				Stop:  "201505062245",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Mr. Jack Pocket", Published: "2010"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201505062200",
				Stop:  "201505062215",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Mr. Jack Pocket", Published: "2010"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201505062130",
				Stop:  "201505062145",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Mr. Jack Pocket", Published: "2010"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "5",

				Start: "201505101900",
				Stop:  "201505102000",
			},
			&models.Location{
				Locationname: "10202 Talleyran Drive, Austin, TX, United States",
				Locationlat:  "30.42509",
				Locationlng:  "-97.80212799999998",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Miguel", Surname: "Coronado"},
				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
				&models.Player{Firstname: "Gary", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
				&models.Played_In{Result: "LOST", Place: "5"},
			},
			[]*models.Game{
				&models.Game{Name: "King of New York", Published: "2014"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201505102000",
				Stop:  "201505102100",
			},
			&models.Location{
				Locationname: "10202 Talleyran Drive, Austin, TX, United States",
				Locationlat:  "30.42509",
				Locationlng:  "-97.80212799999998",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},

				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
				&models.Player{Firstname: "Gary", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201505102200",
				Stop:  "201505102300",
			},
			&models.Location{
				Locationname: "10202 Talleyran Drive, Austin, TX, United States",
				Locationlat:  "30.42509",
				Locationlng:  "-97.80212799999998",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Gary", Surname: "McAdam"},
				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "DEMOLISH", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "4",

				Start: "201505102300",
				Stop:  "201505102359",
			},
			&models.Location{
				Locationname: "10202 Talleyran Drive, Austin, TX, United States",
				Locationlat:  "30.42509",
				Locationlng:  "-97.80212799999998",
			},
			[]*models.Player{

				&models.Player{Firstname: "Gary", Surname: "McAdam"},
				&models.Player{Firstname: "Josephine", Surname: "McAdam"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Frederique", Surname: "McAdam"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "LOST", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "Kingdom Builder", Published: "2011"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201505142230",
				Stop:  "201505142330",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},

			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "San Juan", Published: "2004"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201505312030",
				Stop:  "201505312130",
			},

			&models.Location{
				Locationname: "2505 Side Cove, Austin, TX, United States",
				Locationlat:  "30.251579",
				Locationlng:  "-97.79226399999999",
			},
			[]*models.Player{

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "San Juan", Published: "2004"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201505312130",
				Stop:  "201505312230",
			},
			&models.Location{
				Locationname: "2505 Side Cove, Austin, TX, United States",
				Locationlat:  "30.251579",
				Locationlng:  "-97.79226399999999",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "San Juan", Published: "2004"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201506011930",
				Stop:  "201506011945",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{

				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Castle Keep", Published: "2005"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201506011945",
				Stop:  "201506012000",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Castle Keep", Published: "2005"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201506021945",
				Stop:  "201506022030",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{

				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "Payday", Published: "1973"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201506092030",
				Stop:  "201506092130",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "San Juan", Published: "2004"},
			},
		)

		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201506092130",
				Stop:  "201506092230",
			},

			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "San Juan", Published: "2004"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "2",

				Start: "201506092230",
				Stop:  "201506092330",
			},
			&models.Location{
				Locationname: "Strange Brew, Manchaca Road, Austin, TX, United States",
				Locationlat:  "30.217843",
				Locationlng:  "-97.79648600000002",
			},
			[]*models.Player{

				&models.Player{Firstname: "Myron", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
			},
			[]*models.Game{
				&models.Game{Name: "San Juan", Published: "2004"},
			},
		)
		seedaction(neo,
			&models.Event{
				Numplayers: "3",

				Start: "201506161900",
				Stop:  "201506162000",
			},
			&models.Location{
				Locationname: "2613 W 10th St, Austin, TX, United States",
				Locationlat:  "30.284825",
				Locationlng:  "-97.77455599999996",
			},
			[]*models.Player{
				&models.Player{Firstname: "Olivia", Surname: "Gottlieb"},
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
				&models.Player{Firstname: "Mitch", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
			},
			[]*models.Game{
				&models.Game{Name: "Hot Tin Roof", Published: "2014"},
			},
		)
	}
	return c.Render()
}

//TODO make these objects in events.go when new event comes in IF users are cleared of currentevenst
func seedaction(neo *models.Neo4jObj,
	evt *models.Event,
	loc *models.Location,
	players []*models.Player,
	playedin []*models.Played_In,
	games []*models.Game) {

	//convert old date to RFC3339
	if match, _ := regexp.MatchString("\\d{12}", evt.Start); match {

		fixtime, _ := time.Parse("200601021504 -0700", evt.Start+" -0500")
		evt.Start = fixtime.Format(time.RFC3339)
	}
	if match, _ := regexp.MatchString("\\d{12}", evt.Stop); match {
		fixtime, _ := time.Parse("200601021504 -0700", evt.Stop+" -0500")
		evt.Stop = fixtime.Format(time.RFC3339)
	}
	log.Println("EVT START", evt.Start)
	log.Println("EVT STOP", evt.Stop)

	UUIDEvt := neo.Create(evt)

	UUIDnodeLoc := neo.Create(loc)
	neo.CreateRelate(UUIDEvt, UUIDnodeLoc, &models.Played_At{})

	var UUIDnodeGame string
	for index := range games {
		UUIDnodeGame = neo.Create(games[index])
		neo.CreateRelate(UUIDEvt, UUIDnodeGame, &models.Played_With{})
	}

	var UUIDnodePlayer string
	for index := range players {
		UUIDnodePlayer = neo.Create(players[index])
		neo.CreateRelate(UUIDnodePlayer, UUIDEvt, playedin[index])
		neo.CreateRelate(UUIDEvt, UUIDnodePlayer, &models.Included{})
	}

}

func (c Seed) checkUser() revel.Result {

	log.Println("CHECKING IF FACEBOOKED IN AND ADMIN!!!!")
	user := c.Connected()
	if user == nil || len(user.AccessToken) == 0 {
		c.Flash.Error("Please log in first")
		return c.Redirect(Application.Index)
	}
	return nil
}
