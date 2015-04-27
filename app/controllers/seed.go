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

func (c Seed) Index() revel.Result {

	neo := new(models.Neo4jObj)

	seedaction(neo,
		&models.Event{
			Numplayers: "3",
			Location:   "2613 W10TH",
			Start:      "201409122000",
			Stop:       "201409122100",
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
			Location:   "2613 W10TH",
			Start:      "201409131000",
			Stop:       "201409131100",
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
			Location:   "2613 W10TH",
			Start:      "201409131400",
			Stop:       "201409131500",
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
			Location:   "2613 W10TH",
			Start:      "201409131600",
			Stop:       "201409131700",
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
			Location:   "2613 W10TH",
			Start:      "201409131800",
			Stop:       "201409131900",
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
			Location:   "2613 W10TH",
			Start:      "201409271800",
			Stop:       "201409271900",
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
			Location:   "2613 W10TH",
			Start:      "201409272000",
			Stop:       "201409272100",
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
			Location:   "2613 W10TH",
			Start:      "201409272100",
			Stop:       "201409272200",
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
			Location:   "2613 W10TH",
			Start:      "201409272300",
			Stop:       "201409280000",
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
			Location:   "2613 W10TH",
			Start:      "201409280100",
			Stop:       "201409280200",
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
			Location:   "2613 W10TH",
			Start:      "201409281200",
			Stop:       "201409281300",
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
			Location:   "2613 W10TH",
			Start:      "201409281700",
			Stop:       "201409281800",
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
			Location:   "2613 W10TH",
			Start:      "201410092000",
			Stop:       "201410092200",
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
			Location:   "2613 W10TH",
			Start:      "201410100100",
			Stop:       "201410100300",
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
			Location:   "2613 W10TH",
			Start:      "201410131030",
			Stop:       "201410131130",
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
			Location:   "2613 W10TH",
			Start:      "201410142130",
			Stop:       "201410142230",
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
			Location:   "2613 W10TH",
			Start:      "201410142230",
			Stop:       "201410142330",
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
			Location:   "2613 W10TH",
			Start:      "201410142330",
			Stop:       "201410150030",
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
			Location:   "2613 W10TH",
			Start:      "201410162230",
			Stop:       "201410162330",
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
			Location:   "2613 W10TH",
			Start:      "201410172030",
			Stop:       "201410172130",
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
			Location:   "2613 W10TH",
			Start:      "201410172130",
			Stop:       "201410172230",
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
			Location:   "2613 W10TH",
			Start:      "201410172330",
			Stop:       "201410180130",
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
			Location:   "2613 W10TH",
			Start:      "201410261200",
			Stop:       "201410261300",
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
			Location:   "Emerald Tavern Games & Cafe",
			Start:      "201411081500",
			Stop:       "201411081600",
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
			Location:   "Emerald Tavern Games & Cafe",
			Start:      "201411081600",
			Stop:       "201411081700",
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
			Location:   "Austin Java Barton Springs",
			Start:      "201411111400",
			Stop:       "201411111500",
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
			Location:   "Austin Java Barton Springs",
			Start:      "201411111800",
			Stop:       "201411111900",
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
			Location:   "Austin Java Barton Springs",
			Start:      "201411111900",
			Stop:       "201411112000",
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
			Location:   "2505 Side Cove",
			Start:      "201411271900",
			Stop:       "201411272000",
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
			Location:   "2505 Side Cove",
			Start:      "201411301800",
			Stop:       "201411301900",
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
			Location:   "2613 W10TH",
			Start:      "201411301900",
			Stop:       "201411302000",
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
			Location:   "2613 W10TH",
			Start:      "201412061600",
			Stop:       "201412061700",
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
			Location:   "Austin Java Barton Springs",
			Start:      "201411111800",
			Stop:       "201411111900",
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
			Location:   "Austin Java Barton Springs",
			Start:      "201411111900",
			Stop:       "201411112000",
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
			Location:   "2613 W10TH",
			Start:      "201411261600",
			Stop:       "201411261700",
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
			Location:   "2525 Side Cove",
			Start:      "201501301900",
			Stop:       "201501302000",
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
			Location:   "Radio Coffeehouse",
			Start:      "201502152000",
			Stop:       "201502152010",
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
			Location:   "Radio Coffeehouse",
			Start:      "201502152010",
			Stop:       "201502152020",
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
			Location:   "Radio Coffeehouse",
			Start:      "201502152020",
			Stop:       "201502152030",
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
			Location:   "Radio Coffeehouse",
			Start:      "201502152030",
			Stop:       "201502152040",
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
			Location:   "2613 W10TH",
			Start:      "201502160900",
			Stop:       "201502161000",
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
			Location:   "2525 Side Cove",
			Start:      "201503011500",
			Stop:       "201503011530",
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
			Location:   "2525 Side Cove",
			Start:      "201503011700",
			Stop:       "201503011730",
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
			Location:   "Radio Coffeehouse",
			Start:      "201503071400",
			Stop:       "201503071500",
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
			Location:   "Radio Coffeehouse",
			Start:      "201503071500",
			Stop:       "201503071600",
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
			Location:   "Radio Coffeehouse",
			Start:      "201503071600",
			Stop:       "201503071700",
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
			Location:   "2613 W10TH",
			Start:      "201503081900",
			Stop:       "201503082000",
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
			Location:   "2613 W10TH",
			Start:      "201503082200",
			Stop:       "201503082300",
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
			Location:   "2613 W10TH",
			Start:      "201503082100",
			Stop:       "201503082200",
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
			Location:   "2613 W10TH",
			Start:      "201503132300",
			Stop:       "201503140001",
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
			Location:   "2613 W10TH",
			Start:      "201503140130",
			Stop:       "201503140200",
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
			Location:   "2613 W10TH",
			Start:      "201503140030",
			Stop:       "201503140100",
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
			Location:   "2613 W10TH",
			Start:      "201503172100",
			Stop:       "201503172200",
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
			Location:   "2613 W10TH",
			Start:      "201503172300",
			Stop:       "201503172330",
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
			Location:   "2613 W10TH",
			Start:      "201503172200",
			Stop:       "201503172300",
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
			Location:   "Emerald Tavern Games & Cafe",
			Start:      "201504041500",
			Stop:       "201504041600",
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
			Location:   "2613 W10TH",
			Start:      "201504042300",
			Stop:       "201504050030",
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
			Location:   "2613 W10TH",
			Start:      "201504082000",
			Stop:       "201504082030",
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
			Location:   "Radio Coffeehouse",
			Start:      "201504082200",
			Stop:       "201504082300",
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
			Location:   "Radio Coffeehouse",
			Start:      "201504082300",
			Stop:       "201504082345",
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
			Location:   "2613 W10TH",
			Start:      "201504112000",
			Stop:       "201504112100",
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
			Location:   "Emerald Tavern Games & Cafe",
			Start:      "201504121430",
			Stop:       "201504121530",
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
			Location:   "Radio Coffeehouse",
			Start:      "201504131930",
			Stop:       "201504132030",
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
			Location:   "Radio Coffeehouse",
			Start:      "201504132200",
			Stop:       "201504132230",
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
			Location:   "Radio Coffeehouse",
			Start:      "201504132030",
			Stop:       "201504132130",
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
			Location:   "Strange Brew",
			Start:      "201504242000",
			Stop:       "201504242030",
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
			Location:   "Strange Brew",
			Start:      "201504242030",
			Stop:       "201504242100",
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
			Location:   "Strange Brew",
			Start:      "201504242130",
			Stop:       "201504242200",
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
			Location:   "Strange Brew",
			Start:      "201504242300",
			Stop:       "201504242330",
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
			Location:   "Strange Brew",
			Start:      "201504242200",
			Stop:       "201504242230",
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

	return c.Render()
}

func seedaction(neo *models.Neo4jObj,
	evt *models.Event,
	players []*models.Player,
	playedin []*models.Played_In,
	games []*models.Game) {

	UUIDEvt := neo.Create(evt)

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

func oldindex() {

	var jsoncollection SeedObj

	log.Println(revel.BasePath)

	contents, err := ioutil.ReadFile(revel.BasePath + "/app/db/seed/seed.json")
	check(err)

	json.Unmarshal(contents, &jsoncollection)
	log.Println(jsoncollection)

	neo := new(models.Neo4jObj)

	for _, game := range jsoncollection.Games {

		neo.Create(&game)

	}

	for _, player := range jsoncollection.Players {

		neo.Create(&player)

	}

	//for relationships
	//var jsonPlayed SeedObjPlayed
	contents, err = ioutil.ReadFile(revel.BasePath + "/app/db/seed/seed-played.json")
	check(err)
	var conf Config

	json.Unmarshal(contents, &conf)

	/*
		for _, count := range conf.PlayedList {
			neo.CreateRelate(neo.Read(&count.Game), neo.Read(&count.Player), count.Played)
			//fmt.Printf("%#v\n", count)
			//count.Played.Create(count.Game, count.Player)
		}

	*/

}

func workingaction() {

	neo := new(models.Neo4jObj)
	UUIDEvt := neo.Create(
		&models.Event{
			Numplayers: "3",
			Location:   "2613 W10TH",
			Start:      "201409122000",
			Stop:       "201409122100",
		},
	)
	UUIDnodeGame := neo.Create(&models.Game{
		Name:      "Payday",
		Published: "1973",
	})

	neo.CreateRelate(UUIDEvt, UUIDnodeGame, &models.Played_With{})

	UUIDnodePlayer := neo.Create(&models.Player{
		Firstname: "Olivia",
		Surname:   "Gottlieb",
	})
	neo.CreateRelate(UUIDnodePlayer, UUIDEvt, &models.Played_In{Result: "WON", Place: "1"})
	neo.CreateRelate(UUIDEvt, UUIDnodePlayer, &models.Included{})

	UUIDnodePlayer = neo.Create(&models.Player{
		Firstname: "Myron",
		Surname:   "Gottlieb",
	})

	neo.CreateRelate(UUIDnodePlayer, UUIDEvt, &models.Played_In{Result: "LOST", Place: "3"})
	neo.CreateRelate(UUIDEvt, UUIDnodePlayer, &models.Included{})

	UUIDnodePlayer = neo.Create(&models.Player{
		Firstname: "Mitch",
		Surname:   "Gottlieb",
	})

	neo.CreateRelate(UUIDnodePlayer, UUIDEvt, &models.Played_In{Result: "LOST", Place: "2"})
	neo.CreateRelate(UUIDEvt, UUIDnodePlayer, &models.Included{})
}
