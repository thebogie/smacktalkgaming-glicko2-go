package models

import (
	//"github.com/jmcvetta/neoism"
	"github.com/revel/revel"
	//"log"
	"encoding/json"
	"fmt"
	"net/http"
)

const TimezoneKey = "AIzaSyCXSL3n9tI-VTgRJOhXqJJJ42D1FO1EGBE"
const GeocodeKey = "AIzaSyBvMnC_gxM_viymDC-Et4Jfr9UEMO9l-Hg"

type GoogleAPI struct {
}

//returns timezoneoffset
func (gapi *GoogleAPI) GetTimeZone(lat string, lng string) (retval string) {

	type cargoObj struct {
		DstOffset    int    `json:"dstOffset"`
		RawOffset    int    `json:"rawOffset"`
		Status       string `json:"status"`
		TimeZoneId   string `json:"timeZoneId"`
		TimeZoneName string `json:"timeZoneName"`
	}

	revel.TRACE.Println("Google -> timezone")

	mapstring := fmt.Sprint("https://maps.googleapis.com/maps/api/timezone/json?location=", lat, ",", lng, "&timestamp=0", "&sensor=false&key=", TimezoneKey)

	revel.TRACE.Println("MAPSTRING:", mapstring)

	var retcargo cargoObj

	r, err := http.Get(mapstring)
	revel.TRACE.Println("r", r)
	if err != nil {
		//return err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&retcargo)

	revel.TRACE.Println("DECODED:", retcargo, err)
	return retcargo.TimeZoneId
}
