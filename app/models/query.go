package models

import (
	//"encoding/json"
	"errors"
	//"fmt"
	"github.com/jmcvetta/neoism"
	//"log"
	"bytes"
	"encoding/gob"
	"github.com/revel/revel"
	//"mitchgottlieb.com/smacktalkgaming/app/models"
	"reflect"
	"sort"
	"strconv"
	"strings"
	//"time"
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

func (qobj *QueryObj) GetAllEventUUID() (retval []string) {

	//format of cypher return
	res := []struct {
		EventUUID string `json:"UUID"`
	}{}
	//res := []map[string]map[string]map[string]interface{}{}

	prop := neoism.Props{}

	cq := neoism.CypherQuery{
		Statement:  `match (e:Event) return e.UUID as UUID`,
		Parameters: prop,
		Result:     &res,
	}
	query(&cq)

	for _, v := range res {
		retval = append(retval, v.EventUUID)
	}

	return retval
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
				//revel.TRACE.Println("key , v ", key, v)
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
func (qobj *QueryObj) GetLastEvent(uuid string) Event {

	//find lastevent through player
	res := []Event{}

	prop := neoism.Props{"UUID": uuid}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (m:Player { UUID:{UUID} })-[r:LAST_EVENT]->(n)
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

	/* for _, node := range res {
		revel.TRACE.Println(res)
		revel.TRACE.Println(node)

	} */

	if len(res) < 1 {
		return Event{}
	}

	return res[0]

}

/***************************************************************/
func (qobj *QueryObj) GetEvent(uuid string) Event {

	res := []Event{}

	prop := neoism.Props{"UUID": uuid}

	cq := neoism.CypherQuery{
		Statement: `
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

	/*
		for _, node := range res {
			revel.TRACE.Println(res)
			revel.TRACE.Println(node)

		}
	*/

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
	prop = neoism.Props{"UUID": UUID, "NODELABEL": nodeType, "KEY": key}

	statementStr := `MATCH (n:` + nodeType + `{ UUID:"` + UUID + `"}) return n.` + key + ` as Result `

	cq := neoism.CypherQuery{
		Statement:  statementStr,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	for _, v := range res {
		retval = v.Result
	}

	return retval
}

func (qobj *QueryObj) SetValue(nodeType string, UUID string, key string, value string) error {
	var prop neoism.Props
	//revel.TRACE.Println("SetValue:", nodeType, UUID, key, value)

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

func (qobj *QueryObj) GetAllRatings() (players []Player, ratings []Glicko2) {

	players = qobj.GetAllPlayers()

	for _, player := range players {
		ratings = append(ratings, qobj.GetPlayerGlicko2Rating(player.UUID))

	}

	return players, ratings
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

func (qobj *QueryObj) GetLastLocationNode(playeruuid string) (retval Location) {

	res := []map[string]map[string]map[string]interface{}{}

	prop := neoism.Props{"UUID": playeruuid}

	cq := neoism.CypherQuery{
		Statement:  `MATCH (n:Player {UUID:{UUID}})-[r:LAST_EVENT]-(m), m-[rr:PLAYED_AT]-(l) return l`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	//reflect the struct instead of writing it out eveytime.
	//assume: all strings
	for _, v := range res {
		revel.TRACE.Println("VTYPE", v, reflect.TypeOf(v))

		for ukey, u := range v {
			revel.TRACE.Println("ukey", ukey, u)
			if ukey == "l" {
				for wkey, w := range u {
					revel.TRACE.Println("wkey", wkey, w)
					if wkey == "data" {

						retval = Location{}
						for key, value := range w {
							element := reflect.ValueOf(&retval).Elem().FieldByName(key)
							if element.IsValid() {
								element.SetString(value.(string))
							}

						}

					}

				}
			}
		}
	}

	return retval
}

func (qobj *QueryObj) GetLocation(evt Event) string {

	res := []Location{}

	prop := neoism.Props{"UUID": evt.UUID}

	cq := neoism.CypherQuery{
		Statement: `
			
			MATCH (n:Event { UUID:{UUID} })-[r:PLAYED_AT]->(m)
			return 	m.Locationname as Locationname
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	if len(res) < 1 {
		return "Failed to find last location"
	}
	//revel.TRACE.Println("res", res[0])
	return res[0].Locationname

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

func (qobj *QueryObj) GetPlayerStatsByEvent(eventUUID string) (retval []Player, results []Played_In) {

	//format of cypher return
	res := []map[string]map[string]map[string]interface{}{}

	prop := neoism.Props{"UUID": eventUUID}

	cq := neoism.CypherQuery{
		Statement:  `MATCH (p:Event { UUID:{UUID} })<-[rr:PLAYED_IN]-(n) return n, rr`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	//reflect the struct instead of writing it out eveytime.
	//assume: all strings
	//revel.TRACE.Println("RES", res)
	for _, v := range res {
		//revel.TRACE.Println("V : TYPE", v, reflect.TypeOf(v))
		for _, u := range v {
			//revel.TRACE.Println("UKEY:U", ukey, u)
			for skey, s := range u {
				//revel.TRACE.Println("SKEY:S", skey, s)
				if skey == "data" {
					newevent := Player{}
					for key, value := range s {

						v := reflect.ValueOf(&newevent).Elem().FieldByName(key)
						if v.IsValid() {
							v.SetString(value.(string))
						}

					}
					retval = append(retval, newevent)
				}
				if skey == "rr" {
					newplayedin := Played_In{}
					for key, value := range s {

						v := reflect.ValueOf(&newplayedin).Elem().FieldByName(key)
						if v.IsValid() {
							v.SetString(value.(string))
						}

					}
					results = append(results, newplayedin)
				}

			}
		}

	}

	return retval, results

}

func (qobj *QueryObj) GetRatingByUUID(ratingUUID string) (retval Glicko2) {

	res := []map[string]map[string]map[string]interface{}{}

	prop := neoism.Props{"UUID": ratingUUID}

	cq := neoism.CypherQuery{
		Statement:  `MATCH (n:GLICKO2 { UUID:{UUID} }) return n`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	//reflect the struct instead of writing it out eveytime.
	//assume: all strings
	for _, v := range res {
		//revel.TRACE.Println("VTYPE", v, reflect.TypeOf(v))

		for ukey, u := range v {
			//revel.TRACE.Println("ukey", ukey, u)
			if ukey == "n" {
				for wkey, w := range u {
					//revel.TRACE.Println("wkey", wkey, w)
					if wkey == "data" {

						retval := Glicko2{}
						for key, value := range w {
							element := reflect.ValueOf(&retval).Elem().FieldByName(key)
							if element.IsValid() {
								element.SetString(value.(string))
							}

						}

					}

				}
			}
		}
	}
	return retval
}

func (qobj *QueryObj) SetGlicko2Rating(node Glicko2) {

	res := []map[string]map[string]map[string]interface{}{}

	//revel.TRACE.Println("SETGLICKO2", node)

	prop := neoism.Props{"UUID": node.UUID,
		"NumResults":      node.NumResults,
		"Rating":          node.Rating,
		"RatingDeviation": node.RatingDeviation,
		"Volatility":      node.Volatility}

	cq := neoism.CypherQuery{
		Statement: `MATCH (n:Glicko2 { UUID:{UUID} }) ` +
			`set n.NumResults={NumResults}, ` +
			` n.Rating={Rating}, ` +
			` n.RatingDeviation={RatingDeviation}, ` +
			` n.Volatility={Volatility} ` +
			`return n`,
		Parameters: prop,
		Result:     &res,
	}
	query(&cq)

}

func (qobj *QueryObj) CreateGlicko2PrevRating(ratinguuid string, node Glicko2, date string) (retval string, nodeexists bool) {

	nodeexists = false

	//does the node already exist?
	res := []struct {
		// `json:` tags matches column names in query
		Date string `json:"m.Date"`
		UUID string `json:"m.UUID"`
	}{}

	prop := neoism.Props{"RATINGUUID": ratinguuid}

	cq := neoism.CypherQuery{
		Statement:  `MATCH (n:Glicko2 {UUID:{RATINGUUID}})-[r:RATING_GLICKO2_PREV]->(m) return m.Date, m.UUID`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	revel.TRACE.Println("res", res)

	for _, value := range res {
		if value.Date == date {
			nodeexists = true
			retval = value.UUID
		}
	}

	if !nodeexists {
		neo := new(Neo4jObj)
		neo.init()

		node.Date = date
		retval = neo.Create(&node)
		neo.CreateRelate(ratinguuid, retval, &Rating_Glicko2_prev{})
	}
	return retval, nodeexists

}

//return competitors in placed order!
func (qobj *QueryObj) GetPlayerGlicko2Rating(playerUUID string) (retval Glicko2) {

	res := []map[string]map[string]map[string]interface{}{}

	prop := neoism.Props{"UUID": playerUUID}

	cq := neoism.CypherQuery{
		Statement:  `MATCH (p:Player { UUID:{UUID} })-[:RATING_GLICKO2]->(n) return n`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	//revel.TRACE.Println("did we find a rating?", res)

	if len(res) <= 0 {
		//TODO too low level?
		neo := new(Neo4jObj)
		neo.init()
		glicko2UUID := neo.Create(&Glicko2{})
		neo.CreateRelate(playerUUID, neo.Create(&Glicko2{}), &Rating_Glicko2{})
		retval = qobj.GetRatingByUUID(glicko2UUID)

	} else {
		// get the glicko2 node

		//reflect the struct instead of writing it out eveytime.
		//assume: all strings
		for _, v := range res {
			//revel.TRACE.Println("VTYPE", v, reflect.TypeOf(v))

			for ukey, u := range v {
				//revel.TRACE.Println("ukey", ukey, u)
				if ukey == "n" {
					for wkey, w := range u {
						//revel.TRACE.Println("wkey", wkey, w)
						if wkey == "data" {

							retval = Glicko2{}
							for key, value := range w {
								element := reflect.ValueOf(&retval).Elem().FieldByName(key)
								//revel.TRACE.Println("element and key and value", element, key, value)
								if element.IsValid() {
									element.SetString(value.(string))
								}

							}

						}

					}
				}
			}
		}
	}
	return retval
}

//return competitors in placed order!
//assumption: the rating node exists
func (qobj *QueryObj) GetCompetitorsByEvent(eventUUID string) (retval []Competitor) {

	//format of cypher return
	res := []map[string]map[string]map[string]interface{}{}

	prop := neoism.Props{"UUID": eventUUID}

	cq := neoism.CypherQuery{
		Statement: `MATCH (e:Event { UUID:{UUID} })
					-[r:INCLUDED]->(n), 
					(n)-[rr:RATING_GLICKO2]->(rating),
					(n)-[p:PLAYED_IN]->(e)
					return n, p, rating`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)
	//revel.TRACE.Println("RES", res)
	//reflect the struct instead of writing it out eveytime.
	//assume: all strings
	for _, v := range res {
		//revel.TRACE.Println("v, VTYPE", v, reflect.TypeOf(v))
		newcompetitor := Competitor{}

		for ukey, u := range v {
			//revel.TRACE.Println("ukey,u", ukey, u)
			if ukey == "n" {
				for wkey, w := range u {
					if wkey == "data" {
						newobj := Player{}
						for key, value := range w {

							element := reflect.ValueOf(&newobj).Elem().FieldByName(key)
							if element.IsValid() {
								element.SetString(value.(string))
							}

						}
						newcompetitor.Player = newobj

					}

				}
			}

			if ukey == "p" {
				for wkey, w := range u {
					if wkey == "data" {
						newobj := Played_In{}
						for key, value := range w {

							element := reflect.ValueOf(&newobj).Elem().FieldByName(key)
							if element.IsValid() {
								element.SetString(value.(string))
							}

						}

						newcompetitor.Result = newobj

					}

				}
			}
			if ukey == "rating" {
				for wkey, w := range u {
					if wkey == "data" {
						newobj := Glicko2{}

						for key, value := range w {

							element := reflect.ValueOf(&newobj).Elem().FieldByName(key)

							if element.IsValid() {
								element.SetString(value.(string))
							}

						}

						newcompetitor.Rating = newobj

					}

				}
			}

		}

		//revel.TRACE.Println("NewCompetitor Rating", newcompetitor.Rating)

		//if !ratingfound {
		//	newcompetitor.Rating = new(QueryObj).GetPlayerGlicko2Rating(newcompetitor.Player.UUID)
		//}

		retval = append(retval, newcompetitor)

	}
	sort.Sort(ByPlace(retval))

	return retval

}

func (qobj *QueryObj) GetGamesByEvent(eventUUID string) (retval []Game) {

	//format of cypher return
	res := []map[string]map[string]map[string]interface{}{}

	prop := neoism.Props{"UUID": eventUUID}

	//revel.TRACE.Println("prop", prop)

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (p:Event { UUID:{UUID} })-[r:PLAYED_WITH]->(n)
			
			return n
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	//reflect the struct instead of writing it out eveytime.
	//assume: all strings
	for _, v := range res {

		for _, u := range v {
			for s, t := range u {
				if s == "data" {
					newevent := Game{}
					for key, value := range t {

						element := reflect.ValueOf(&newevent).Elem().FieldByName(key)
						if element.IsValid() {
							element.SetString(value.(string))
						}

					}
					//revel.TRACE.Println("gamebyevent", newevent)
					retval = append(retval, newevent)
				}

			}
		}

	}

	if len(retval) == 0 {
		retval = nil
	}

	return retval

}

func (qobj *QueryObj) GetEventsByPlayer(playerUUID string) (retval []Event) {

	//format of cypher return
	res := []map[string]map[string]map[string]interface{}{}

	prop := neoism.Props{"UUID": playerUUID}

	cq := neoism.CypherQuery{
		Statement: `
			MATCH (p:Player { UUID:{UUID} })-[r:PLAYED_IN]->(n)
			return n
			`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	//reflect the struct instead of writing it out eveytime.
	//assume: all strings
	for _, v := range res {
		//revel.TRACE.Println("TYPE", k, reflect.TypeOf(v))
		for _, u := range v {
			for s, t := range u {
				if s == "data" {
					newevent := Event{}
					for key, value := range t {

						element := reflect.ValueOf(&newevent).Elem().FieldByName(key)
						if element.IsValid() {
							element.SetString(value.(string))
						}

					}
					retval = append(retval, newevent)
				}

			}
		}

	}

	return retval

}

//gather total number games played in a month
//input: "2014-01", player UUID
func (qobj *QueryObj) GetPlayersTotalPlaysByMonth(month string, playeruuid string) (retval int) {
	retval = 0

	res := []map[string]int{}

	prop := neoism.Props{"UUID": playeruuid, "MONTH": month}

	cq := neoism.CypherQuery{
		Statement:  `MATCH (n:Event)<-[r:PLAYED_IN]-(m:Player {UUID:{UUID}}) WHERE n.Start STARTS WITH {MONTH} return count(r)`,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)
	//revel.TRACE.Println("RES", res)
	//assume: all strings
	//revel.TRACE.Println("retval = ", (res[0])["count(r)"])
	if len(res) > 0 {
		retval = (res[0])["count(r)"]
	}
	/*
		for _, v := range res {
			revel.TRACE.Println("V:", v, "REFLECT:" , reflect.TypeOf(v))
			for ukey, u := range v {
			revel.TRACE.Println("u:", u, "ukey:" , ukey)
				for skey, s := range u {
					revel.TRACE.Println("s:", s, "skey:" , skey)
					if skey == "data" {
						//retval = s
					}

				}
			}
		}
	*/

	return retval

}

//gather all the events in the month and year
//input: "2014-01"
//return list of eventuuids
func (qobj *QueryObj) GetAllEventUUIDsByMonth(month string) (retval []string) {

	res := []struct {
		// `json:` tags matches column names in query
		UUID string `json:"n.UUID"`
	}{}

	prop := neoism.Props{"MONTH": month}

	cq := neoism.CypherQuery{
		Statement:  `MATCH (n:Event) WHERE n.Start STARTS WITH {MONTH} return n.UUID `,
		Parameters: prop,
		Result:     &res,
	}

	query(&cq)

	//revel.TRACE.Println("res", res)

	for _, value := range res {

		retval = append(retval, value.UUID)
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

	//for _, node := range res {
	//revel.TRACE.Println("res", node)
	//retval = append(retval, Event{Location: node.Location, Surname: node.Surname, UUID: node.UUID})
	//}

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
		//revel.TRACE.Println("R:", v.Result, v.Game)

		if _, ok := retval[v.Game]; !ok {
			retval[v.Game] = make(map[string]int)
		}

		if _, ok := retval[v.Game][v.Result]; !ok {

			retval[v.Game][v.Result] = 1
			//map[string]int{v.Result: 1}

		} else {

			retval[v.Game][v.Result]++

		}
		//revel.TRACE.Println("RETVAL", retval)

	}

	//for k, v := range retval {
	//revel.TRACE.Println("RETVAL V", v)
	//revel.TRACE.Println("RETVAL K", k)
	//}

	return retval

}
