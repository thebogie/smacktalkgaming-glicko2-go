package controllers

import (
	"fmt"
	"github.com/revel/revel"
	//"golang.org/x/crypto/bcrypt"
	"mitchgottlieb.com/smacktalkgaming/app/models"
	//"mitchgottlieb.com/smacktalkgaming/app/routes"
	//"strings"
)

type Players struct {
	Application
}

func (c Players) List() revel.Result {
	qobj := new(models.QueryObj)
	players := qobj.GetAllPlayers()
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
	fmt.Println("HERE")
	return c.Render(players)
}
