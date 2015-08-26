package models

import (
	//"encoding/json"
	"errors"
	//"fmt"
	"github.com/jmcvetta/neoism"
	//"log"
	"strconv"
	//"mitchgottlieb.com/smacktalkgaming/app/models"
	"bytes"
	"encoding/gob"
	"github.com/revel/revel"
	"reflect"
	"strings"
)

type QueryObj struct {
}

func query(load *neoism.CypherQuery) {

	//revel.TRACE.Println("RELFECT:", reflect.ValueOf(load))

	neo := new(Neo4jObj)
	neo.init()

	neo.dbc.Session.Log = false
	neo.dbc.Cypher(load)

	//revel.TRACE.Println("AFTER CYPHER", load.Result)

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
	//revel.TRACE.Println("RELFECT:", reflect.TypeOf(r.Rel_count))
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
	//revel.TRACE.Println("RELFECT:", reflect.TypeOf(r.Rel_count))
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
	//revel.TRACE.Println("RELFECT:", reflect.TypeOf(r.Rel_count))
	return strconv.Itoa(r.Rel_count)

}

func (qobj *QueryObj) MatchPlayersByName(find string, playerUUID string) []Player {

	/*
		   find all the player's events and the players for each of those events
		   MATCH (n:Player {UUID: "e1426765-cf35-4df4-840d-b07f810a8eb9"})-[:PLAYED_IN]-(e:Event)-[:INCLUDED]-(p)


				MATCH (n:Player) RETURN CASE SUBSTRING( n.Firstname + ' ' + n.Surname, 0, {sizeofstring})
				WHEN {findstring} THEN  n ELSE null END as result
	*/
	retval := []Player{}

	res := []struct {
		// `json:` tags matches column names in query
		NodeReturned neoism.Node `json:"result"`
	}{}

	//localres := []Player{}
	prop := neoism.Props{"sizeofstring": len(find), "findstring": find, "PROPUUID": playerUUID}

	cq := neoism.CypherQuery{
		Statement: `
		
			MATCH (n:Player {UUID: {PROPUUID}})-[:PLAYED_IN]-(e:Event)-[:INCLUDED]-(p) RETURN
			CASE SUBSTRING( p.Firstname + ' ' + p.Surname, 0, {sizeofstring}) 
			WHEN {findstring} THEN  p  END as result
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	revel.TRACE.Println("LOCAL USERS SIZE", len(res))
	for _, node := range res {

		//revel.TRACE.Println("reflect ", reflect.TypeOf(node.NodeReturned.Data))
		if len(node.NodeReturned.Data) > 0 {
			//how do i get the Data struct into a Player stuct to send back
			//read reflect again....

			playerObj := new(Player)

			player := reflect.ValueOf(&playerObj).Elem()
			//tempPlayer = (Player)

			//retval = append(retval, node.NodeReturned.Data.(Player))

			for key, v := range node.NodeReturned.Data {
				//revel.TRACE.Println("key , v ", key, v)
				if len(v.(string)) > 0 {
					player.Elem().FieldByName(key).SetString(v.(string))
				}

			}

			alreadyexists := false
			for _, player := range retval {
				if player.UUID == playerObj.UUID {
					alreadyexists = true
				}
			}
			if !alreadyexists {
				retval = append(retval, *playerObj)
			}
		}

	}

	//find all the other users in the system
	res = []struct {
		// `json:` tags matches column names in query
		NodeReturned neoism.Node `json:"result"`
	}{}

	//localres := []Player{}
	prop = neoism.Props{"sizeofstring": len(find), "findstring": find}

	cq = neoism.CypherQuery{
		Statement: `
			MATCH (p:Player) RETURN
			CASE SUBSTRING( p.Firstname + ' ' + p.Surname, 0, {sizeofstring}) 
			WHEN {findstring} THEN  p  END as result
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	revel.TRACE.Println("LOCAL USERS SIZE", len(res))

	for _, node := range res {

		//revel.TRACE.Println("reflect ", reflect.TypeOf(node.NodeReturned.Data))
		if len(node.NodeReturned.Data) > 0 {
			//how do i get the Data struct into a Player stuct to send back
			//read reflect again....
			playerObj := new(Player)

			player := reflect.ValueOf(&playerObj).Elem()
			//tempPlayer = (Player)

			//retval = append(retval, node.NodeReturned.Data.(Player))

			for key, v := range node.NodeReturned.Data {
				//revel.TRACE.Println("key , v ", key, v)
				if len(v.(string)) > 0 {
					player.Elem().FieldByName(key).SetString(v.(string))
				}

			}

			alreadyexists := false
			for _, player := range retval {
				if player.UUID == playerObj.UUID {
					alreadyexists = true
				}
			}
			if !alreadyexists {
				retval = append(retval, *playerObj)
			}

		}

	}
	revel.TRACE.Println("END OF PLAYER SEARCH", retval)
	return retval
}

func (qobj *QueryObj) MatchGamesByName(find string, playerUUID string) []Game {

	if len(find) == 0 {

		return []Game{}
	}

	retval := []Game{}

	//dont search for 4 or less.. too many ogames
	//if len(find) < 4 {
	//	return retval
	//}

	res := []struct {
		// `json:` tags matches column names in query
		NodeReturned neoism.Node `json:"result"`
	}{}

	prop := neoism.Props{"sizeofstring": len(find), "findstring": find, "PROPUUID": playerUUID}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (p:Player {UUID: {PROPUUID}})-[:PLAYED_IN]-(e:Event)-[:PLAYED_WITH]-(g) RETURN
			CASE SUBSTRING( g.Name , 0, {sizeofstring}) 
			WHEN {findstring} THEN  g  END as result
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)
	for _, node := range res {

		//revel.TRACE.Println("reflect ", reflect.TypeOf(node.NodeReturned.Data))
		if len(node.NodeReturned.Data) > 0 {
			//how do i get the Data struct into a Player stuct to send back
			//read reflect again....
			gameObj := new(Game)

			game := reflect.ValueOf(&gameObj).Elem()

			for key, v := range node.NodeReturned.Data {
				//revel.TRACE.Println("key , v ", key, v)
				if len(v.(string)) > 0 {
					game.Elem().FieldByName(key).SetString(v.(string))
				}

			}
			alreadyexists := false
			for _, game := range retval {
				if game.UUID == gameObj.UUID {
					alreadyexists = true
				}
			}
			if !alreadyexists {
				retval = append(retval, *gameObj)
			}

		}

	}

	//find all the other games in the system
	res = []struct {
		// `json:` tags matches column names in query
		NodeReturned neoism.Node `json:"result"`
	}{}

	//localres := []Player{}
	prop = neoism.Props{}
	capletter := []byte(find)
	startswithletter := "STARTS_WITH_" + strings.ToUpper(string(capletter[0]))

	cq = neoism.CypherQuery{
		Statement:  "match (n:Gamesifter)-[r:" + startswithletter + "]-(g:Game) return g as result",
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	revel.TRACE.Println("LOCAL USERS SIZE", len(res))

	for _, node := range res {

		//revel.TRACE.Println("reflect ", reflect.TypeOf(node.NodeReturned.Data))
		if len(node.NodeReturned.Data) > 0 {
			//how do i get the Data struct into a Player stuct to send back
			//read reflect again....
			gameObj := new(Game)

			game := reflect.ValueOf(&gameObj).Elem()
			//tempPlayer = (Player)

			//retval = append(retval, node.NodeReturned.Data.(Player))

			for key, v := range node.NodeReturned.Data {
				//revel.TRACE.Println("SEARCHBYLETTER: key , v ", key, v)
				if len(v.(string)) > 0 {
					game.Elem().FieldByName(key).SetString(v.(string))
				}

			}

			alreadyexists := false
			for _, game := range retval {
				if game.UUID == gameObj.UUID {
					alreadyexists = true
				}
			}
			if !alreadyexists {
				retval = append(retval, *gameObj)
			}

		}

	}

	revel.TRACE.Println("END OF GAME SEARCH", retval)
	return retval
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (qobj *QueryObj) GetPlayerCurrentEvent(uuid string) string {

	var retval string

	res := []Player{}

	prop := neoism.Props{"UUID": uuid}

	cq := neoism.CypherQuery{
		Statement: `
			start n=node(*)
			MATCH (n:Player { UUID:{UUID} })
			return n.CurrentEvent as Currentevent
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	for _, node := range res {
		retval = node.CurrentEvent
	}
	return retval
}

/******************************************************/

func (qobj *QueryObj) GetOverallStats(uuid string) ([]Event, []Played_In, []Game) {

	type OverallStatsObj struct {
		Playedins []neoism.Relationship `json:"playedins"`
		Events    []neoism.Node         `json:"events"`
		Games     []neoism.Node         `json:"games"`
	}

	var events []Event
	var playedins []Played_In
	var games []Game

	res := []OverallStatsObj{}
	//res := []interface{}{}

	//res := make(map[string]resObj)

	prop := neoism.Props{"PROPUUID": uuid}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (p:Player {UUID: {PROPUUID}})-[r:PLAYED_IN]-(m:Event)-[t:PLAYED_WITH]-(g:Game)
			RETURN COLLECT(r) AS playedins , COLLECT(m) AS events , COLLECT( DISTINCT g ) AS games
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	if len(res) < 1 {
		return events, playedins, games
	}

	for _, node := range res {

		for _, event := range node.Events {
			result := &Event{}
			err := FillStruct(event.Data, result)
			if err != nil {
				revel.ERROR.Panic(err)
			}

			events = append(events, *result)
		}

		for _, game := range node.Games {

			result := &Game{}
			err := FillStruct(game.Data, result)
			if err != nil {
				revel.ERROR.Panic(err)
			}

			games = append(games, *result)
		}

		for _, playedin := range node.Playedins {

			result := &Played_In{}
			err := FillStruct(playedin.Data.(map[string]interface{}), result)
			if err != nil {
				revel.ERROR.Panic(err)
			}

			playedins = append(playedins, *result)
		}

	}
	return events, playedins, games
}

/******************************************************/

func (qobj *QueryObj) GetPlayer(uuid string) Player {

	retval := Player{}

	res := []struct {
		// `json:` tags matches column names in query
		NodeReturned neoism.Node `json:"result"`
	}{}

	prop := neoism.Props{"PROPUUID": uuid}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (p:Player {UUID: {PROPUUID}}) RETURN p as result
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	for _, node := range res {

		if len(node.NodeReturned.Data) > 0 {
			playerObj := new(Player)

			player := reflect.ValueOf(&playerObj).Elem()

			for key, v := range node.NodeReturned.Data {
				revel.TRACE.Println("key , v ", key, v)
				if len(v.(string)) > 0 {
					player.Elem().FieldByName(key).SetString(v.(string))
				}

			}
			retval = *playerObj

		}

	}

	return retval
}

/****************************************************************/

/***************************************************************/
func (qobj *QueryObj) GetEvent(uuid string) Event {

	res := []Event{}

	prop := neoism.Props{"UUID": uuid}

	cq := neoism.CypherQuery{
		Statement: `
			start n=node(*)
			MATCH (n:Event { UUID:{UUID} })
			return 	n.Eventname as Eventname, 
					n.Numplayers as Numplayers,
					n.Start     as Start,
					n.Stop as Stop,
					n.UUID as UUID
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	for _, node := range res {
		revel.TRACE.Println(res)
		revel.TRACE.Println(node)

	}

	if len(res) < 1 {
		return Event{}
	}

	return res[0]
}

func (qobj *QueryObj) GetValue(nodeType string, UUID string, key string) string {
	var prop neoism.Props
	var retval string

	res := []struct {
		// `json:` tags matches column names in query
		Result string `json:"Result"`
	}{}
	prop = neoism.Props{"UUID": UUID, "NODE": nodeType, "KEY": key}

	cq := neoism.CypherQuery{
		Statement: `
			start n=node(*)
			MATCH (n:{NODE} { UUID:{UUID} })
			return n.{KEY} as Result
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	//for _, node := range res {
	//retval = node
	//}
	return retval
}

func (qobj *QueryObj) SetValue(nodeType string, UUID string, key string, value string) error {
	var prop neoism.Props
	revel.TRACE.Println("SetValue:", nodeType, UUID, key, value)

	res := []struct {
		// `json:` tags matches column names in query
		Result string `json:"Result"`
	}{}
	prop = neoism.Props{}

	statementStr := `start n=node(*) MATCH (n:` + nodeType + `{ UUID:"` + UUID + `"}) set n.` + key + `= "` + value + `" return n.` + key + ` as Result `

	cq := neoism.CypherQuery{
		Statement:  statementStr,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	for _, v := range res {
		if v.Result != value {
			return errors.New("setvalue: value not set properly")
		}
	}
	return nil
}

func (qobj *QueryObj) GetAllPlayers() []Player {

	var retval []Player

	//retval := make([]Player, 10)

	res := []Player{}

	/*[]struct {
		// `json:` tags matches column names in query
		Firstname string `json:"n.Firstname"`
		Surname   string `json:"n.Surname"`
		UUID      string `json:"n.UUID"`
	}{} */

	prop := neoism.Props{}

	cq := neoism.CypherQuery{
		Statement: `
			start n=node(*)
			MATCH (n:Player)
			where has(n.UUID)
			return n.Firstname as Firstname, 
			       n.Surname as Surname, 
				   n.UUID as UUID
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	for _, node := range res {

		retval = append(retval, Player{Firstname: node.Firstname, Surname: node.Surname, UUID: node.UUID})
	}

	//test := res[0].Firstname
	////revel.TRACE.Println("RELFECT:", reflect.TypeOf(test))
	//revel.TRACE.Println("RELFECT:", test)

	/*keys := make([]string, 10)
	for k := range test {
		revel.TRACE.Println("KEYMAYE:", k)
		keys = append(keys, k)
	}

	//fmt.Printf("%+v\n", test)
	revel.TRACE.Println("RELFECT:", keys)
	*/
	//revel.TRACE.Println("RELFECT:", res[5])
	return retval

}

func (qobj *QueryObj) GetAllGames() []Game {

	res := []Game{}

	prop := neoism.Props{}

	cq := neoism.CypherQuery{
		Statement: `
			start n=node(*)
			MATCH (n:Game)
			return n.Name as Name, 
			       n.Published as Published,  
				   n.UUID as UUID
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	for _, node := range res {
		revel.TRACE.Println("res", node)
		//retval = append(retval, Event{Location: node.Location, Surname: node.Surname, UUID: node.UUID})
	}

	return res

}

func (qobj *QueryObj) GetAllEventLocations() (retval []string) {

	res := []Event{}

	prop := neoism.Props{}

	cq := neoism.CypherQuery{
		Statement: `
			start n=node(*)
			MATCH (n:Event)
			
			return  n.Location as Location
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	for _, node := range res {
		revel.TRACE.Println("res", node)

		var found = false
		//for _, item := range retval {
		//if item == node.Location {
		//	found = true
		//}
		//}
		if !found {
			retval = append(retval, "")
		}
	}

	return retval

}

func (qobj *QueryObj) GetAllEvents() []Event {

	//var retval []Event

	res := []Event{}

	/*
		res := []struct {
			// `json:` tags matches column names in query
			Location   string `json:"n.Location"`
			Numplayers string `json:"n.Numplayers"`
			Start      string `json:"n.Start"`
			Stop       string `json:"n.Stop"`
			UUID       string `json:"n.UUID"`
		}{}
	*/

	prop := neoism.Props{}

	cq := neoism.CypherQuery{
		Statement: `
			start n=node(*)
			MATCH (n:Event)
			
			return n.Eventname as Eventname, 
			       n.Location as Location, 
				   n.Numplayers as Numplayers, 
				   n.Start as Start, 
				   n.Stop as Stop, 
				   n.UUID as UUID
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	for _, node := range res {
		revel.TRACE.Println("res", node)
		//retval = append(retval, Event{Location: node.Location, Surname: node.Surname, UUID: node.UUID})
	}

	return res

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
		revel.TRACE.Println("R:", v.Result, v.Game)

		if _, ok := retval[v.Game]; !ok {
			retval[v.Game] = make(map[string]int)
		}

		if _, ok := retval[v.Game][v.Result]; !ok {

			retval[v.Game][v.Result] = 1
			//map[string]int{v.Result: 1}

		} else {

			retval[v.Game][v.Result]++

		}
		revel.TRACE.Println("RETVAL", retval)

	}

	for k, v := range retval {
		revel.TRACE.Println("RETVAL V", v)
		revel.TRACE.Println("RETVAL K", k)
	}

	return retval

}
