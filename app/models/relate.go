// node
package models

import (
	//"fmt"
	"github.com/jmcvetta/neoism"
	//"log"
	"github.com/revel/revel"
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

type Rating_Glicko2 struct {
	UUID string
}

type Played_With struct {
	UUID string
}

type Starts_With struct {
	UUID       string
	Relatename string
}

type Played_At struct {
	UUID string
}

type Included struct {
	UUID string
}

type Last_Event struct {
	UUID string
}

/******** PLAYED IN *************/
func (r *Played_In) Create() neoism.Props {

	revel.TRACE.Println("Creating  ", reflect.TypeOf(r))
	r.UUID = getUUID()
	return neoism.Props{
		"Result": r.Result,
		"Place":  r.Place,
		"UUID":   r.UUID,
	}

}

func (r *Played_In) Read() string {

	revel.TRACE.Println("Reading  ", reflect.TypeOf(r))
	//revel.TRACE.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "asdf"
	//"MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}

/******** PLAYED WITH *************/
func (r *Played_With) Create() neoism.Props {

	revel.TRACE.Println("Creating  ", reflect.TypeOf(r))
	r.UUID = getUUID()
	return neoism.Props{
		"UUID": r.UUID,
	}

}

func (r *Played_With) Read() string {

	revel.TRACE.Println("Reading  ", reflect.TypeOf(r))
	//revel.TRACE.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "asdf"
	//"MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}

/******** STARTS WITH *************/
func (r *Starts_With) Create() neoism.Props {

	revel.TRACE.Println("Creating  ", reflect.TypeOf(r))
	r.UUID = getUUID()
	return neoism.Props{
		"UUID":       r.UUID,
		"Relatename": r.Relatename,
	}

}

func (r *Starts_With) Read() string {

	revel.TRACE.Println("Reading  ", reflect.TypeOf(r))
	//revel.TRACE.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "asdf"
	//"MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}

/******** PLAYED AT *************/
func (r *Played_At) Create() neoism.Props {

	revel.TRACE.Println("Creating  ", reflect.TypeOf(r))
	r.UUID = getUUID()
	return neoism.Props{
		"UUID": r.UUID,
	}

}

func (r *Played_At) Read() string {

	revel.TRACE.Println("Reading  ", reflect.TypeOf(r))
	//revel.TRACE.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "asdf"
	//"MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}

/******** INCULDED *************/
func (r *Included) Create() neoism.Props {

	revel.TRACE.Println("Creating  ", reflect.TypeOf(r))
	r.UUID = getUUID()
	return neoism.Props{
		"UUID": r.UUID,
	}

}

func (r *Included) Read() string {

	revel.TRACE.Println("Reading  ", reflect.TypeOf(r))
	//revel.TRACE.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "asdf"
	//"MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}

/******** LAST EVENT *************/
func (r *Last_Event) Create() neoism.Props {

	revel.TRACE.Println("Creating  ", reflect.TypeOf(r))
	r.UUID = getUUID()
	return neoism.Props{
		"UUID": r.UUID,
	}

}

func (r *Last_Event) Read() string {

	revel.TRACE.Println("Reading  ", reflect.TypeOf(r))
	//revel.TRACE.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "asdf"
	//"MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}

/******** GLICKO2 Rating *************/
func (r *Rating_Glicko2) Create() neoism.Props {

	revel.TRACE.Println("Creating ", reflect.TypeOf(r))
	r.UUID = getUUID()
	return neoism.Props{
		"UUID": r.UUID,
	}

}

func (r *Rating_Glicko2) Read() string {

	revel.TRACE.Println("Reading  ", reflect.TypeOf(r))
	//revel.TRACE.Println("Searching for  ", n.Name, n.Published, n.UUID)
	return "asdf"
	//"MATCH (node:Game { Name:\"" + n.Name + "\", Published:\"" + n.Published + "\" }) RETURN node"

}
