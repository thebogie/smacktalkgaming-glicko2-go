// node
package models

import (
	"bytes"
	"encoding/json"
	"github.com/jmcvetta/neoism"
	"github.com/revel/revel"
	"io/ioutil"
	//"log"
	"net/http"
	"reflect"
	"time"
)

type Node interface {
	Create() neoism.Props
	Read() string
}

/*
type Event struct {
	Start      string
	Stop       string
	Numplayers string
	Location   string
	UUID       string
}
*/

type Glicko2 struct {
	UUID            string `json:"UUID"`
	Rating          string `json:"Rating"`
	RatingDeviation string `json:"RatingDeviation"`
	Volatility      string `json:"Volatility"`
	NumResults      string `json:"NumResults"`
	Date			string `json:"Date"`
}

type Gamesifter struct {
	Name string
	UUID string `json:"UUID"`
}

type Location struct {
	Locationname string `json:"Locationname"`
	Locationlng  string `json:"Locationlng"`
	Locationlat  string `json:"Locationlat"`
	Locationtz   string `json:"Locationtz"`
	UUID         string `json:"UUID"`
}

type Event struct {
	Eventname  string `json:"Eventname"`
	Numplayers string `json:"Numplayers"`
	Start      string `json:"Start"`
	Stop       string `json:"Stop"`
	UUID       string `json:"UUID"`
}

type Game struct {
	Name      string `json:"Name"`
	Published string `json:"Published"`
	UUID      string `json:"UUID"`
	BGGLink   string `json:"BGGLink`
}

type Competitor struct {
	Player Player
	Result Played_In
	Rating Glicko2
}

type ByPlace []Competitor

func (a ByPlace) Len() int           { return len(a) }
func (a ByPlace) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPlace) Less(i, j int) bool { return a[i].Result.Place < a[j].Result.Place }

type Player struct {
	Firstname    string `json:"Firstname"`
	Surname      string `json:"Surname"`
	Nickname     string `json:"Nickname"`
	Birthdate    string `json:"Birthdate"`
	UUID         string `json:"UUID"`
	CurrentEvent string `json:"Currentevent"`
	Alignment    string `json:"Alignment"`
	Admin        string `json:"Admin"`
}

func getUUID() string {
	newUUID, uuidErr := newUUID()
	if uuidErr != nil {
		revel.ERROR.Panic("CREATE UUID ERROR:", uuidErr)
	}
	return newUUID

}

func getEventName() string {

	var retval bytes.Buffer

	type wordObj struct {
		Id   int    `json:"id"`
		Word string `json:"word"`
	}

	type eventnameObj struct {
		noun wordObj
		adj  wordObj
	}

	var eventname eventnameObj

	//adjective!
	adjURL := "http://api.wordnik.com:80/v4/words.json/randomWord?hasDictionaryDef=true&includePartOfSpeech=adjective&minCorpusCount=0&maxCorpusCount=-1&minDictionaryCount=1&maxDictionaryCount=-1&minLength=0&maxLength=-1&api_key=a2a73e7b926c924fad7001ca3111acd55af2ffabf50eb4ae5"
	nounURL := "http://api.wordnik.com:80/v4/words.json/randomWord?hasDictionaryDef=true&includePartOfSpeech=noun&minCorpusCount=0&maxCorpusCount=-1&minDictionaryCount=1&maxDictionaryCount=-1&minLength=0&maxLength=-1&api_key=a2a73e7b926c924fad7001ca3111acd55af2ffabf50eb4ae5"

	res, err := http.Get(adjURL)
	eventname.adj.Word = "boring"
	if err == nil {

		//found an adjective
		jsonDataFromHttp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		revel.TRACE.Println("EVENT adj", string(jsonDataFromHttp))

		err = json.Unmarshal([]byte(jsonDataFromHttp), &eventname.adj) // here!

		if err != nil {
			panic(err)
		}
	}

	res, err = http.Get(nounURL)
	eventname.noun.Word = "PUTUUID"
	if err == nil {

		//found an adjective
		jsonDataFromHttp, err := ioutil.ReadAll(res.Body)

		if err != nil {
			panic(err)
		}

		revel.TRACE.Println("EVENT: noun", string(jsonDataFromHttp))

		err = json.Unmarshal([]byte(jsonDataFromHttp), &eventname.noun) // here!
		if err != nil {
			panic(err)
		}
	}
	retval.WriteString("the ")
	retval.WriteString(eventname.adj.Word + " ")
	retval.WriteString(eventname.noun.Word + " ")
	retval.WriteString("event")

	revel.TRACE.Println("ADJ RETURN", retval.String())
	revel.TRACE.Println("ADJ RETURN", eventname)

	return retval.String()
}

/******** GLICKO2 NODE *************/
func (n *Glicko2) Create() neoism.Props {
	revel.TRACE.Println("Creating  ", reflect.TypeOf(n))
	n.UUID = getUUID()
	
	if (n.Rating =="" ) {
	n.Rating = "1500"
	n.RatingDeviation = "350"
	n.Volatility = "0.06"
	n.NumResults = "0"
	n.Date = "0"
}
	return neoism.Props{
		"UUID":            n.UUID,
		"Rating":          n.Rating,
		"RatingDeviation": n.RatingDeviation,
		"Volatility":      n.Volatility,
		"NumResults":      n.NumResults,
		"Date":				n.Date,
	}
}

func (n *Glicko2) Read() string {

	revel.TRACE.Println("Reading  ", reflect.TypeOf(n))
	revel.TRACE.Println("Searching for  ", n.UUID)
	return "MATCH (node:Glicko2 { UUID:\"" + n.UUID + "\" }) RETURN node"
}

/******** LOCATION NODE *************/
func (n *Location) Create() neoism.Props {
	revel.TRACE.Println("Creating  ", reflect.TypeOf(n))
	n.UUID = getUUID()
	return neoism.Props{
		"Locationlat":  n.Locationlat,
		"Locationlng":  n.Locationlng,
		"Locationname": n.Locationname,
		"UUID":         n.UUID}
}

func (n *Location) Read() string {

	revel.TRACE.Println("Reading  ", reflect.TypeOf(n))
	revel.TRACE.Println("Searching for  ", n.Locationlat, n.Locationlng)
	return "MATCH (node:Location { Locationlat:\"" + n.Locationlat + "\", Locationlng:\"" + n.Locationlng + "\" }) RETURN node"
}

/******** EVENT NODE *************/

func (n *Event) Create() neoism.Props {

	revel.TRACE.Println("Creating  ", reflect.TypeOf(n))

	n.UUID = getUUID()
	n.Eventname = getEventName()
	if n.Start == "" {
		t := time.Now()
		n.Start = t.Format(time.RFC3339)
	}

	return neoism.Props{
		"Eventname":  n.Eventname,
		"Start":      n.Start,
		"Stop":       n.Stop,
		"Numplayers": n.Numplayers,
		"UUID":       n.UUID,
	}

}

func (n *Event) Read() string {

	revel.TRACE.Println("Reading  ", reflect.TypeOf(n))
	revel.TRACE.Println("Searching for  ", n.Eventname)
	return "MATCH (node:Event { Eventname:\"" + n.Eventname + "\" }) RETURN node"

}

/******** GAME NODE *************/
func (n *Game) Create() neoism.Props {

	revel.TRACE.Println("Creating  ", reflect.TypeOf(n))
	n.UUID = getUUID()

	//connect first letter in game name to gamesifter

	return neoism.Props{"Name": n.Name, "Published": n.Published, "UUID": n.UUID, "BGGLink": n.BGGLink}

}

func (n *Game) Read() string {

	revel.TRACE.Println("Reading  ", reflect.TypeOf(n))
	revel.TRACE.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}

/******** Gamesifter Node ***********/
func (n *Gamesifter) Create() neoism.Props {

	//create or find the GameCatalog node. This node's only purpose is to split the games
	//up into alphabet and other for quicker searching.

	revel.TRACE.Println("Creating  ", reflect.TypeOf(n))
	n.UUID = getUUID()
	return neoism.Props{"Name": n.Name, "UUID": n.UUID}

}

func (n *Gamesifter) Read() string {

	revel.TRACE.Println("Reading  ", reflect.TypeOf(n))
	revel.TRACE.Println("Searching for  ", n.UUID)
	return "MATCH (node:Gamesifter { Name:\"" + n.Name + "\" }) RETURN node"

}

/******** PLAYER NODE *************/
func (n *Player) Create() neoism.Props {

	revel.TRACE.Println("Creating  ", reflect.TypeOf(n))
	n.UUID = getUUID()

	return neoism.Props{
		"Firstname": n.Firstname,
		"Surname":   n.Surname,
		"Birthdate": n.Birthdate,
		"Nickname":  n.Nickname,
		"UUID":      n.UUID,
		"Admin":     n.Admin}

}

func (n *Player) Read() string {
	revel.TRACE.Println("Reading  ", reflect.TypeOf(n))

	revel.TRACE.Println("Searching for  ", n.Firstname, n.Surname, n.UUID)
	return "MATCH (node:Player { Firstname:\"" + n.Firstname + "\", Surname:\"" + n.Surname + "\" }) RETURN node"

}
