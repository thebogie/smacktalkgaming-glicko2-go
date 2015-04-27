package models

import (
	//"encoding/json"
	//"fmt"
	"github.com/jmcvetta/neoism"
	"log"
	"strconv"
	//"mitchgottlieb.com/smacktalkgaming/app/models"
	"reflect"
)

type QueryObj struct {
}

func query(load *neoism.CypherQuery) {

	log.Println("RELFECT:", reflect.ValueOf(load))

	neo := new(Neo4jObj)
	neo.init()

	neo.dbc.Session.Log = true
	neo.dbc.Cypher(load)

	log.Println("AFTER CYPHER", load.Result)

}

func (qobj *QueryObj) TotalNumberOfGamesPlayed(UUID string) string {

	res := []struct {
		// `json:` tags matches column names in query
		Rel_count int `json:"rel_count"`
	}{}

	prop := neoism.Props{
		"UUID": UUID}

	cq := neoism.CypherQuery{
		Statement: `
			start n=node(*)
			match (n:Player {UUID:{UUID}})-[r:PLAYED_IN]-(c)
			return count(r) as rel_count order by rel_count desc
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	r := res[0]
	//log.Println("RELFECT:", reflect.TypeOf(r.Rel_count))
	return strconv.Itoa(r.Rel_count)

}
func (qobj *QueryObj) TotalGamesWon(UUID string) string {

	/*
	 */

	res := []struct {
		// `json:` tags matches column names in query
		Rel_count int `json:"rel_count"`
	}{}

	prop := neoism.Props{
		"UUID": UUID}

	cq := neoism.CypherQuery{
		Statement: `
			start n=node(*)
			MATCH (n:Player { UUID:{UUID} })-[r:PLAYED_IN]->() 
			WHERE (r.Result = "WON" OR r.Result = "DEMOLISH")
			AND r.Place = "1"
			RETURN count(r) as rel_count
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	r := res[0]
	//log.Println("RELFECT:", reflect.TypeOf(r.Rel_count))
	return strconv.Itoa(r.Rel_count)

}
func (qobj *QueryObj) TotalGamesLost(UUID string) string {

	/*
	 */

	res := []struct {
		// `json:` tags matches column names in query
		Rel_count int `json:"rel_count"`
	}{}

	prop := neoism.Props{
		"UUID": UUID}

	cq := neoism.CypherQuery{
		Statement: `
			start n=node(*)
			MATCH (n:Player { UUID:{UUID} })-[r:PLAYED_IN]->() 
			WHERE (r.Result = "LOST" OR r.Result = "DROP")
			AND NOT r.Place = "1"
			RETURN count(r) as rel_count
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	r := res[0]
	//log.Println("RELFECT:", reflect.TypeOf(r.Rel_count))
	return strconv.Itoa(r.Rel_count)

}

func (qobj *QueryObj) GetAllPlayers() []Player {

	var retval []Player

	//retval := make([]Player, 10)

	res := []struct {
		// `json:` tags matches column names in query
		Firstname string `json:"n.Firstname"`
		Surname   string `json:"n.Surname"`
		UUID      string `json:"n.UUID"`
	}{}

	prop := neoism.Props{}

	cq := neoism.CypherQuery{
		Statement: `
			start n=node(*)
			MATCH (n:Player)
			where has(n.UUID)
			return n.Firstname, n.Surname, n.UUID
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	for _, node := range res {

		retval = append(retval, Player{Firstname: node.Firstname, Surname: node.Surname, UUID: node.UUID})
	}

	//test := res[0].Firstname
	////log.Println("RELFECT:", reflect.TypeOf(test))
	//log.Println("RELFECT:", test)

	/*keys := make([]string, 10)
	for k := range test {
		log.Println("KEYMAYE:", k)
		keys = append(keys, k)
	}

	//fmt.Printf("%+v\n", test)
	log.Println("RELFECT:", keys)
	*/
	//log.Println("RELFECT:", res[5])
	return retval

}

// Game Name ,WON, LOST
//TODO add demolish drop
func (qobj *QueryObj) OverallGameRecord(UUID string) map[string]map[string]int {

	res := []struct {
		// `json:` tags matches column names in query
		Result string `json:"Result"`
		Game   string `json:"Game"`
	}{}

	prop := neoism.Props{
		"UUID": UUID}

	cq := neoism.CypherQuery{
		Statement: `
			start n=node(*)
			match (n:Player {UUID:{UUID}})-[r:PLAYED_IN]->(c)
			match (c)-[r2:PLAYED_WITH]->(game)
			return r.Result as Result, game.Name as Game
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	retval := make(map[string]map[string]int)
	for _, v := range res {
		//log.Println("R:", v.Result, v.Game)

		if _, ok := retval[v.Game]; !ok {
			retval[v.Game] = make(map[string]int)
		}

		if _, ok := retval[v.Game][v.Result]; !ok {

			retval[v.Game] = map[string]int{v.Result: 1}

		} else {

			retval[v.Game][v.Result]++

		}

	}

	for k, v := range retval {
		log.Println("RETVAL V", v)
		log.Println("RETVAL K", k)
	}

	return retval

}
