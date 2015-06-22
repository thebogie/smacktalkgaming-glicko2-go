package controllers

import (
	//"fmt"
	"github.com/revel/revel"
	//"mitchgottlieb.com/smacktalkgaming/app/models"
	//"strconv"
)

func init() {

	//revel.OnAppStart(InitDB)
	revel.InterceptFunc(setuser, revel.BEFORE, &Application{})
	revel.InterceptMethod(Seed.checkUser, revel.BEFORE)
	revel.InterceptMethod(Profile.checkUser, revel.BEFORE)

	//revel.InterceptFunc(setuser, revel.BEFORE, &revel.Controller{})
	//revel.InterceptFunc(checkUser, revel.BEFORE, &revel.Controller{})
	//revel.InterceptMethod(Application.checkUser, revel.BEFORE)
	//

	//revel.InterceptFunc(checkUser, revel.BEFORE, &Seed{})
	//revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	//revel.InterceptMethod(Application.AddUser, revel.BEFORE)
	//revel.InterceptMethod(Application.checkUser, revel.BEFORE)
	//revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	//revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}
