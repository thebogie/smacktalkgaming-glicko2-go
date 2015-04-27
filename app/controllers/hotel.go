package controllers

import (
	"fmt"
	"github.com/revel/revel"
	//"golang.org/x/crypto/bcrypt"
	//"mitchgottlieb.com/smacktalkgaming/app/models"
	//"mitchgottlieb.com/smacktalkgaming/app/routes"
	//"strings"
)

type Hotels struct {
	Application
}

func (c Hotels) checkUser() revel.Result {
	/*
		if user := c.connected(); user == nil {
			c.Flash.Error("Please log in first")
			return c.Redirect(routes.Application.Index())
		}
	*/
	return nil
}

func (c Hotels) Index() revel.Result {

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
	return c.Render("FISH")
}
