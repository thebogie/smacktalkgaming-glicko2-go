// node
package models

import (
	//"fmt"
	"github.com/jmcvetta/neoism"
	"log"
	"reflect"
)

type Relate interface {
	Create() neoism.Props
	Read() string
}

type Played_In struct {
	Result string
	Place  string
	UUID   string
}

type Played_With struct {
	UUID string
}

type Included struct {
	UUID string
}

/******** PLAYED RELATE *************/

func (r *Played_In) Create() neoism.Props {

	log.Println("Creating  ", reflect.TypeOf(r))
	r.UUID = getUUID()
	return neoism.Props{
		"Result": r.Result,
		"Place":  r.Place,
		"UUID":   r.UUID,
	}

}

func (r *Played_In) Read() string {

	log.Println("Reading  ", reflect.TypeOf(r))
	//log.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "asdf"
	//"MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}

func (r *Played_With) Create() neoism.Props {

	log.Println("Creating  ", reflect.TypeOf(r))
	r.UUID = getUUID()
	return neoism.Props{
		"UUID": r.UUID,
	}

}

func (r *Played_With) Read() string {

	log.Println("Reading  ", reflect.TypeOf(r))
	//log.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "asdf"
	//"MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}

func (r *Included) Create() neoism.Props {

	log.Println("Creating  ", reflect.TypeOf(r))
	r.UUID = getUUID()
	return neoism.Props{
		"UUID": r.UUID,
	}

}

func (r *Included) Read() string {

	log.Println("Reading  ", reflect.TypeOf(r))
	//log.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "asdf"
	//"MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}
