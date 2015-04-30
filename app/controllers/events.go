package controllers

import (
	"fmt"
	"github.com/revel/revel"
	//"golang.org/x/crypto/bcrypt"
	"mitchgottlieb.com/smacktalkgaming/app/models"
	//"mitchgottlieb.com/smacktalkgaming/app/routes"
	//"strings"
)

type Events struct {
	Application
}

func (c Events) List() revel.Result {
	qobj := new(models.QueryObj)
	events := qobj.GetAllEvents()
	/*
		results, err := c.Txn.Select(models.Booking{},
			`select * from Booking where UserId = ?`, c.connected().UserId)
		if err != nil {
			panic(err)
		}

		var bookings []*models.Booking
		for _, r := range results {
			b := r.(*models.Booking)
			bookings = append(bookings, b)
		}
	*/
	fmt.Println("EVENTS", events)
	return c.Render(events)
}
