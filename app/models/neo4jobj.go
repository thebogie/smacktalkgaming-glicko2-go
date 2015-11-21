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
	//"log"
	"strings"
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
	//revel.TRACE.Println("NEO4j INIT")

	if neo.dbc == nil {

		neo4jdburl, _ := revel.Config.String("appvars.neo4jdburl")

		dbconnect, err := neoism.Connect(neo4jdburl + "/db/data")
		neo.dbc = dbconnect
		neo.dbc.Session.Log = false

		if err != nil {
			revel.TRACE.Panicln("Database not connecting", err)
		}

	}

}

/* A -> B relationship */
func (neo *Neo4jObj) CreateRelate(UUIDnodeA string, UUIDnodeB string, relate Relate) (UUID string) {
	revel.TRACE.Println("NEO4j CREATE RELATE")
	neo.init()

	statements := []string{}

	relateProps := relate.Create()
	bundleProps := neoism.Props{"relateProps": relateProps, "UUIDnodeA": UUIDnodeA, "UUIDnodeB": UUIDnodeB}

	revel.TRACE.Println("relateProps = ", relateProps)

	//var statementStr string

	switch t := relate.(type) {

	case *Starts_With:
		statements = append(statements, `
			match a, b where a.UUID ={UUIDnodeA}
			AND b.UUID = {UUIDnodeB} 
			CREATE (a)-[r:`+relateProps["Relatename"].(string)+` {relateProps}]->(b) RETURN r
		`)

	case *Played_At:
		statements = append(statements, `
			match a, b where a.UUID ={UUIDnodeA} 
			AND b.UUID = {UUIDnodeB} 
			CREATE (a)-[r:PLAYED_AT {relateProps}]->(b) RETURN r
		`)
	case *Played_In:
		statements = append(statements, `
			match a, b where a.UUID ={UUIDnodeA} 
			AND b.UUID = {UUIDnodeB} 
			CREATE (a)-[r:PLAYED_IN {relateProps}]->(b) RETURN r
		`)
	case *Rating_Glicko2_prev:
		statements = append(statements, `
			match a, b where a.UUID ={UUIDnodeA} 
			AND b.UUID = {UUIDnodeB} 
			CREATE (a)-[r:RATING_GLICKO2_PREV {relateProps}]->(b) RETURN r
		`)
	case *Rating_Glicko2:
		statements = append(statements, `
			match a, b where a.UUID ={UUIDnodeA} 
			AND b.UUID = {UUIDnodeB} 
			CREATE (a)-[r:RATING_GLICKO2 {relateProps}]->(b) RETURN r
		`)
	case *Played_With:
		statements = append(statements, `
			match a, b where a.UUID ={UUIDnodeA} 
			AND b.UUID = {UUIDnodeB} 
			CREATE (a)-[r:PLAYED_WITH {relateProps}]->(b) RETURN r
		`)
	case *Included:
		statements = append(statements, `
			match a, b where a.UUID ={UUIDnodeA} 
			AND b.UUID = {UUIDnodeB} 
			CREATE (a)-[r:INCLUDED {relateProps}]->(b) RETURN r
		`)
	case *Last_Event:
		statements = append(statements, `
			match (a:Player {UUID:{UUIDnodeA}})-[r:LAST_EVENT]->() DELETE r
		`)
		statements = append(statements, `
			match a, b where a.UUID ={UUIDnodeA} 
			AND b.UUID = {UUIDnodeB} 
			CREATE (a)-[r:LAST_EVENT {relateProps}]->(b) RETURN r
		`)

	default:
		revel.TRACE.Println("NODE TYPE", t)
	}

	//revel.TRACE.Println("bndleprops:", bundleProps)

	for _, statement := range statements {

		res1 := []struct {
			Node neoism.Relationship `json:"relationship"`
		}{}

		//revel.TRACE.Println("Runing statement:", statement)

		cq := neoism.CypherQuery{
			Statement:  statement,
			Parameters: bundleProps,
			Result:     &res1,
		}
		neo.dbc.Session.Log = false
		neo.dbc.Cypher(&cq)

		//revel.TRACE.Println("RES: ", res1)
	}
	return UUID
}

func (neo *Neo4jObj) ReadRelate(relate Relate) (UUID string) {
	revel.TRACE.Println("NEO4j READ RELATE")
	neo.init()
	return UUID
}

func (neo *Neo4jObj) Create(node Node) (UUID string) {
	revel.TRACE.Println("NEO4j CREATE")
	neo.init()

	var newNode *neoism.Node

	//doesnt exist?
	cargo := neo.Read(node)

	if len(cargo.Data) == 0 {
		var err error
		newProps := node.Create()
		revel.TRACE.Println("NewProps:", newProps)

		newNode, err = neo.dbc.CreateNode(newProps)
		if err != nil {
			revel.ERROR.Fatal(err)
		}

		var label string

		switch t := node.(type) {
		case *Gamesifter:
			label = "Gamesifter"
		case *Game:
			label = "Game"
		case *Player:
			label = "Player"

			//create rating node and relate
			UUIDGLICKO2 := neo.Create(&Glicko2{})
			neo.CreateRelate(UUID, UUIDGLICKO2, &Rating_Glicko2{})
		case *Event:
			label = "Event"

		case *Location:
			label = "Location"
		case *Glicko2:
			label = "Glicko2"
		default:
			revel.TRACE.Println("NODE TYPE", t)
		}

		newNode.AddLabel(label)

		UUID, _ = newNode.Property("UUID")

	}

	if len(UUID) == 0 {
		UUID, _ = cargo.Data["UUID"].(string)

	}

	switch node.(type) {

	case *Game:

		var gamename string
		if len(cargo.Data) == 0 {
			gamename, _ = newNode.Property("Name")

			//If game node connect the gamesifter
			//create or find the GameCatalog node. This node's only purpose is to split the games
			//up into alphabet and other for quicker searching.
			UUIDGamesifter := neo.Create(&Gamesifter{Name: "Gamesifter"})
			revel.TRACE.Println("-------gamesifter", UUIDGamesifter)

			capletter := []byte(gamename)
			revel.TRACE.Println("gamename + cap", gamename, string(capletter[0]))
			startswithletter := "STARTS_WITH_" + strings.ToUpper(string(capletter[0]))
			neo.CreateRelate(UUIDGamesifter, UUID, &Starts_With{Relatename: startswithletter})
		}

	}

	return UUID

}

//read from UUID... find other ways to find objects later
func (neo *Neo4jObj) Read(node Node) neoism.Node {
	revel.TRACE.Println("NEO4j READ")
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
