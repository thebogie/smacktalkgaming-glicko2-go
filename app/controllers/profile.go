// profile
package controllers

import (
	"github.com/revel/revel"
	//"log"
	"mitchgottlieb.com/smacktalkgaming/app/models"
)

type Profile struct {
	Application
}

func getProfile(UUID string) map[string]map[string]string {

	qobj := new(models.QueryObj)
	//revel.TRACE.Println("RETURN", qobj.TotalNumberOfGamesPlayed())
	//retval := map[string]interface{}{}
	retval := map[string]map[string]string{
		"PLAYERPROFILEUUID": map[string]string{
			"tag":   "",
			"value": UUID},
		"NUMBEROFGAMES": map[string]string{
			"tag":   "Total Games played",
			"value": qobj.TotalNumberOfGamesPlayed(UUID)},
		"NUMBEROFGAMESWON": map[string]string{
			"tag":   "Total Games Won or Demolish",
			"value": qobj.TotalGamesWon(UUID)},
		"TOTALGAMESLOST": map[string]string{
			"tag":   "Total Games Lost",
			"value": qobj.TotalGamesLost(UUID)},
	}

	return retval
}

func getOverall(playerUUID string) map[string]map[string]int {

	qobj := new(models.QueryObj)
	return qobj.OverallGameRecord(playerUUID)
}

func (c Profile) Show(uuid string) revel.Result {
	permission := make(map[string]string)
	permission["readonly"] = "false"
	if uuid != c.Session["playerUUID"] {
		permission["readonly"] = "true"
	}

	playerinfo := new(models.QueryObj).GetPlayer(uuid)

	events, playedins, games := new(models.QueryObj).GetOverallStats(uuid)
	//retval2 := getOverall(uuid)

	revel.TRACE.Println("playerinfo: ", playerinfo, events)

	return c.Render(playerinfo, events, playedins, games, permission)
}

func (c Profile) Index() revel.Result {

	user := c.Connected()
	revel.TRACE.Println("WHICH USER", user.PlayerUUID)
	retval := getProfile(user.PlayerUUID)

	//qobj := new(models.QueryObj)
	//retval["NUMBEROFGAMES"] = qobj.QueryTotalNumberOfGamesPlayed()
	//retval["FISH"] = "FISH"

	return c.Render(retval)
}

func (c Profile) checkUser() revel.Result {

	revel.TRACE.Println("CHECKING IF FACEBOOKED IN")
	user := c.Connected()
	if user == nil || len(user.AccessToken) == 0 {
		c.Flash.Error("Please log in first")
		return c.Redirect(Application.Index)
	}
	return nil
}
