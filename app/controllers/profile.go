// profile
package controllers

import (
	"github.com/revel/revel"
	"log"
	"mitchgottlieb.com/smacktalkgaming/app/models"
)

type Profile struct {
	Application
}

func getProfile(UUID string) map[string]map[string]string {

	qobj := new(models.QueryObj)
	//log.Println("RETURN", qobj.TotalNumberOfGamesPlayed())
	//retval := map[string]interface{}{}
	retval := map[string]map[string]string{
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

func getOverall(UUID string) map[string]map[string]int {

	qobj := new(models.QueryObj)
	return qobj.OverallGameRecord(UUID)
}

func (c Profile) Show(uuid string) revel.Result {

	retval := getProfile(uuid)
	retval2 := getOverall(uuid)

	log.Println("RETVAL2: ", retval2)

	return c.Render(retval, retval2)
}

func (c Profile) Index() revel.Result {

	user := c.Connected()
	log.Println("WHICH USER", user.PlayerUUID)
	retval := getProfile(user.PlayerUUID)

	//qobj := new(models.QueryObj)
	//retval["NUMBEROFGAMES"] = qobj.QueryTotalNumberOfGamesPlayed()
	//retval["FISH"] = "FISH"

	return c.Render(retval)
}

func (c Profile) checkUser() revel.Result {

	log.Println("CHECKING IF FACEBOOKED IN")
	user := c.Connected()
	if user == nil || len(user.AccessToken) == 0 {
		c.Flash.Error("Please log in first")
		return c.Redirect(Application.Index)
	}
	return nil
}
