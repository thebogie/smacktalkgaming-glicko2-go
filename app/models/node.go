// node
package models

import (
	//"fmt"
	"github.com/jmcvetta/neoism"
	"log"
	"reflect"
)

type Node interface {
	Create() neoism.Props
	Read() string
}

type Event struct {
	Start      string
	Stop       string
	Numplayers string
	Location   string
	UUID       string
}

type Game struct {
	Name      string
	Published string
	UUID      string
}

type Player struct {
	Firstname string
	Surname   string
	Nickname  string
	Birthdate string
	UUID      string
}

func getUUID() string {
	newUUID, uuidErr := newUUID()
	if uuidErr != nil {
		log.Panic("CREATE UUID ERROR:", uuidErr)
	}
	return newUUID

}

/******** EVENT NODE *************/

func (n *Event) Create() neoism.Props {

	log.Println("Creating  ", reflect.TypeOf(n))
	n.UUID = getUUID()
	return neoism.Props{
		"Start":      n.Start,
		"Stop":       n.Stop,
		"Numplayers": n.Numplayers,
		"Location":   n.Location,
		"UUID":       n.UUID,
	}

}

func (n *Event) Read() string {

	log.Println("Reading  ", reflect.TypeOf(n))
	//log.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "asdf"
	//"MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}

/******** GAME NODE *************/
func (n *Game) Create() neoism.Props {

	log.Println("Creating  ", reflect.TypeOf(n))
	n.UUID = getUUID()
	return neoism.Props{"Name": n.Name, "Published": n.Published, "UUID": n.UUID}

}

func (n *Game) Read() string {

	log.Println("Reading  ", reflect.TypeOf(n))
	log.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

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
