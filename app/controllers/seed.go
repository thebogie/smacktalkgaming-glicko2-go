package controllers

import (
	//"bytes"
	//"encoding/gob"
	"encoding/json"
	//"github.com/jmcvetta/neoism"
	"github.com/revel/revel"
	"io/ioutil"
	//"log"
	"mitchgottlieb.com/smacktalkgaming/app/models"
	//"reflect"
	//"fmt"
	//"encoding/xml"
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
	revel.TRACE.Println("SEED", kind)

	devmode, _ := revel.Config.String("mode.dev")

	if devmode == "false" {
		return c.Redirect("/")
	}
	neo := new(models.Neo4jObj)

	if kind == "games" {
	}

	//monthly (complete month!) check update the numbers (also the non competed game deviation changes)
	if kind == "refreshrating" {
		//TODO: fix for montly runs
		months := []string{"2014-09", "2014-10", "2014-11", "2014-12", "2015-01", "2015-02", "2015-03", "2015-04", "2015-05", "2015-06", "2015-07", "2015-08", "2015-09"}
		
		//months := []string{"2014-09"}
		//month := "2014-09-13"

		type PlayerStats struct {
			Player []models.Player
			Rating []models.Glicko2
		}

		var playerstats PlayerStats
		//update all players with a ratings if they dont have it

		var qobj models.QueryObj

		playerstats.Player, playerstats.Rating = qobj.GetAllRatings()
		
		revel.TRACE.Println("SEED:PLAYERS:", playerstats.Player, "SEED:RATING:", playerstats.Rating)

		
		//TODO: this is done by months, but.. the set of games SHOULD be ALL the games up to this point 
		// and then each game after that... NOT a month, 12 times....
		var totalevents []string
		for _, month := range months {

			for _, eventuuid := range  qobj.GetAllEventUUIDsByMonth(month) {
				//get each event and churn through the ratings
				totalevents = append(totalevents, eventuuid)

			}
/* USE THIS FOR MONTHLY CRON ADJUSTEMNT
			//now readjust the deviation for those that played less then 5 games

			for index, player := range playerstats.Player {
				if qobj.GetPlayersTotalPlaysByMonth(month, player.UUID) < 5 {
					afterPeriodNonCompeteDevUpdate(
						playerstats.Rating[index].UUID,
						playerstats.Rating[index].RatingDeviation,
						playerstats.Rating[index].Volatility)

				}
			}
*/
		}
		
		for _, eventuuid := range totalevents {
			afterEventRankingUpdate(eventuuid)
		}

		//revel.TRACE.Println("playerstats", playerstats)
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
		for i := 0; i < 200; i++ {
			var randomname randomnameObj

			//adjective!
			getnameURL := "http://api.randomuser.me/"

			res, err := http.Get(getnameURL)
			if err == nil {
				jsonDataFromHttp, err := ioutil.ReadAll(res.Body)
				if err != nil {
					panic(err)
				}
				revel.TRACE.Println("Random Name", string(jsonDataFromHttp))

				err = json.Unmarshal([]byte(jsonDataFromHttp), &randomname) // here!

				if err != nil {
					panic(err)
				}

				revel.TRACE.Println("Random Name after", randomname)

				var player models.Player

				capFirst := []rune(randomname.Results[0].User.NameObj.First)
				capFirst[0] = unicode.ToUpper(capFirst[0])

				capLast := []rune(randomname.Results[0].User.NameObj.Last)
				capLast[0] = unicode.ToUpper(capLast[0])

				player.Firstname = string(capFirst)
				player.Surname = string(capLast)
				player.Nickname = randomname.Results[0].User.Username

				revel.TRACE.Println("player", player)

				UUIDnodePlayer := neo.Create(&player)
				revel.TRACE.Println("player", UUIDnodePlayer)

			}
			sum += i
		}

	}
	if kind == "onevent" {
		seedaction(neo,
			&models.Event{
				Numplayers: "4",
				Start:      "2015-10-20T03:35:00+01:00",
				Stop:       "2015-10-20T04:35:00+01:00",
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
				&models.Player{Firstname: "Lilyan", Surname: "Gottlieb"},
			},
			[]*models.Played_In{
				&models.Played_In{Result: "WON", Place: "1"},
				&models.Played_In{Result: "LOST", Place: "2"},
				&models.Played_In{Result: "LOST", Place: "3"},
				&models.Played_In{Result: "DROP", Place: "4"},
			},
			[]*models.Game{
				&models.Game{Name: "Payday", Published: "1973"},
			},
		)

	}

	if kind == "events" {

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
	revel.TRACE.Println("EVT START", evt.Start)
	revel.TRACE.Println("EVT STOP", evt.Stop)

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

	//Update rankings
	revel.TRACE.Println("UPDATE RANKINGS:", UUIDEvt)
	afterEventRankingUpdate(UUIDEvt)

}

func (c Seed) checkUser() revel.Result {

	revel.TRACE.Println("CHECKING IF FACEBOOKED IN AND ADMIN!!!!")
	user := c.Connected()
	if user == nil || len(user.AccessToken) == 0 {
		c.Flash.Error("Please log in first")
		return c.Redirect(Application.Index)
	}
	return nil
}
