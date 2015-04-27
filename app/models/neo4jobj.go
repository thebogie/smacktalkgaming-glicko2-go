// neo4j.go
/*

questions:
How many times were you beaten by John?
How many times did you win Game X?
How many time have you played Game X?
Did I ever play John in Game X?
When was the last time i played Game X?
When was the last time I played John and Mary?
Where did we play Game X the most?

*/
package models

import (
	"github.com/jmcvetta/neoism"
	"github.com/revel/revel"
	"log"

	//"reflect"
)

/*
	Init() *neoism.Database
	Create(map[string]string) string
	Read() map[string]string
	Update(string) bool
	Delete(string) bool
*/

type Neo4jObj struct {
	dbc *neoism.Database
}

func (neo *Neo4jObj) init() {
	log.Println("NEO4j INIT")

	if neo.dbc == nil {

		neo4jdburl, _ := revel.Config.String("appvars.neo4jdburl")

		dbconnect, err := neoism.Connect(neo4jdburl + "/db/data")
		neo.dbc = dbconnect

		if err != nil {
			log.Panicln("Database not connecting", err)
		}

	}

}

/* A -> B relationship */
func (neo *Neo4jObj) CreateRelate(UUIDnodeA string, UUIDnodeB string, relate Relate) (UUID string) {
	log.Println("NEO4j CREATE RELATE")
	neo.init()

	relateProps := relate.Create()

	var statementStr string

	switch t := relate.(type) {
	case *Played_In:
		statementStr = `
			match a, b where a.UUID ={UUIDnodeA} 
			AND b.UUID = {UUIDnodeB} 
			CREATE (a)-[r:PLAYED_IN {relateProps}]->(b) RETURN r
		`
	case *Played_With:
		statementStr = `
			match a, b where a.UUID ={UUIDnodeA} 
			AND b.UUID = {UUIDnodeB} 
			CREATE (a)-[r:PLAYED_WITH {relateProps}]->(b) RETURN r
		`
	case *Included:
		statementStr = `
			match a, b where a.UUID ={UUIDnodeA} 
			AND b.UUID = {UUIDnodeB} 
			CREATE (a)-[r:INCLUDED {relateProps}]->(b) RETURN r
		`
	default:
		log.Println("NODE TYPE", t)
	}

	bundleProps := neoism.Props{"relateProps": relateProps, "UUIDnodeA": UUIDnodeA, "UUIDnodeB": UUIDnodeB}

	log.Println("relateProps:", bundleProps)

	res1 := []struct {
		Node neoism.Relationship `json:"relationship"`
	}{}

	cq := neoism.CypherQuery{
		Statement:  statementStr,
		Parameters: bundleProps,
		Result:     &res1,
	}
	neo.dbc.Session.Log = true
	neo.dbc.Cypher(&cq)

	return UUID
}

func (neo *Neo4jObj) ReadRelate(relate Relate) (UUID string) {
	log.Println("NEO4j READ RELATE")
	neo.init()
	return UUID
}

func (neo *Neo4jObj) Create(node Node) (UUID string) {
	log.Println("NEO4j CREATE")
	neo.init()

	//doesnt exist?
	cargo := neo.Read(node)
	log.Println("CARGO:", cargo.Data)
	log.Println("CARGO:", len(cargo.Data))

	if len(cargo.Data) == 0 {

		newProps := node.Create()
		log.Println("NewProps:", newProps)

		newNode, err := neo.dbc.CreateNode(newProps)
		if err != nil {
			log.Fatal(err)
		}

		var label string

		switch t := node.(type) {
		case *Game:
			label = "Game"
		case *Player:
			label = "Player"
		case *Event:
			label = "Event"
		default:
			log.Println("NODE TYPE", t)
		}

		newNode.AddLabel(label)

		UUID, _ = newNode.Property("UUID")

	}

	if len(UUID) == 0 {
		UUID, _ = cargo.Data["UUID"].(string)

	}

	return UUID

}

//read from UUID... find other ways to find objects later
func (neo *Neo4jObj) Read(node Node) neoism.Node {
	log.Println("NEO4j READ")
	neo.init()

	statementStr := node.Read()
	res1 := []struct {
		Node neoism.Node `json:"node"`
	}{}

	cq := neoism.CypherQuery{
		Statement:  statementStr,
		Parameters: neoism.Props{},
		Result:     &res1,
	}
	neo.dbc.Session.Log = false
	neo.dbc.Cypher(&cq)

	if len(res1) > 0 {
		return res1[0].Node
	}

	return neoism.Node{}

}
