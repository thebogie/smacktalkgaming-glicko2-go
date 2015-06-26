// node
package models

import (
	"bytes"
	"encoding/json"
	"github.com/jmcvetta/neoism"
	"io/ioutil"
	"log"
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

type Gamesifter struct {
	Name string
	UUID string `json:"UUID"`
}

type Location struct {
	Locationname string `json:"Locationname"`
	Locationlng  string `json:"Locationlng"`
	Locationlat  string `json:"Locationlat"`
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

type Player struct {
	Firstname    string `json:"Firstname"`
	Surname      string `json:"Surname"`
	Nickname     string `json:"Nickname"`
	Birthdate    string `json:"Birthdate"`
	UUID         string `json:"UUID"`
	CurrentEvent string `json:"Currentevent`
}

func getUUID() string {
	newUUID, uuidErr := newUUID()
	if uuidErr != nil {
		log.Panic("CREATE UUID ERROR:", uuidErr)
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
		log.Println("EVENT adj", string(jsonDataFromHttp))

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

		log.Println("EVENT: noun", string(jsonDataFromHttp))

		err = json.Unmarshal([]byte(jsonDataFromHttp), &eventname.noun) // here!
		if err != nil {
			panic(err)
		}
	}
	retval.WriteString("the ")
	retval.WriteString(eventname.adj.Word + " ")
	retval.WriteString(eventname.noun.Word + " ")
	retval.WriteString("event")

	log.Println("ADJ RETURN", retval.String())
	log.Println("ADJ RETURN", eventname)

	return retval.String()
}

/******** LOCATION NODE *************/
func (n *Location) Create() neoism.Props {
	log.Println("Creating  ", reflect.TypeOf(n))
	n.UUID = getUUID()
	return neoism.Props{
		"Locationlat":  n.Locationlat,
		"Locationlng":  n.Locationlng,
		"Locationname": n.Locationname,
		"UUID":         n.UUID}
}

func (n *Location) Read() string {

	log.Println("Reading  ", reflect.TypeOf(n))
	log.Println("Searching for  ", n.Locationlat, n.Locationlng)
	return "MATCH (node:Location { Locationlat:\"" + n.Locationlat + "\", Locationlng:\"" + n.Locationlng + "\" }) RETURN node"
}

/******** EVENT NODE *************/

func (n *Event) Create() neoism.Props {

	log.Println("Creating  ", reflect.TypeOf(n))

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

	log.Println("Reading  ", reflect.TypeOf(n))
	log.Println("Searching for  ", n.Eventname)
	return "MATCH (node:Event { Eventname:\"" + n.Eventname + "\" }) RETURN node"

}

/******** GAME NODE *************/
func (n *Game) Create() neoism.Props {

	log.Println("Creating  ", reflect.TypeOf(n))
	n.UUID = getUUID()

	//connect first letter in game name to gamesifter

	return neoism.Props{"Name": n.Name, "Published": n.Published, "UUID": n.UUID, "BGGLink": n.BGGLink}

}

func (n *Game) Read() string {

	log.Println("Reading  ", reflect.TypeOf(n))
	log.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}

/******** Gamesifter Node ***********/
func (n *Gamesifter) Create() neoism.Props {

	//create or find the GameCatalog node. This node's only purpose is to split the games
	//up into alphabet and other for quicker searching.

	log.Println("Creating  ", reflect.TypeOf(n))
	n.UUID = getUUID()
	return neoism.Props{"Name": n.Name, "UUID": n.UUID}

}

func (n *Gamesifter) Read() string {

	log.Println("Reading  ", reflect.TypeOf(n))
	log.Println("Searching for  ", n.UUID)
	return "MATCH (node:Gamesifter { Name:\"" + n.Name + "\" }) RETURN node"

}

/******** PLAYER NODE *************/
func (n *Player) Create() neoism.Props {

	log.Println("Creating  ", reflect.TypeOf(n))
	n.UUID = getUUID()
	return neoism.Props{
		"Firstname": n.Firstname,
		"Surname":   n.Surname,
		"Birthdate": n.Birthdate,
		"Nickname":  n.Nickname,
		"UUID":      n.UUID}

}

func (n *Player) Read() string {
	log.Println("Reading  ", reflect.TypeOf(n))

	log.Println("Searching for  ", n.Firstname, n.Surname, n.UUID)
	return "MATCH (node:Player { Firstname:\"" + n.Firstname + "\", Surname:\"" + n.Surname + "\" }) RETURN node"

}
