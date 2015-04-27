package controllers

import (
	//"fmt"
	"github.com/revel/revel"
	//"mitchgottlieb.com/smacktalkgaming/app/models"
	//"strconv"
)

/*
func connected(rc *revel.Controller) *models.User {
	return rc.RenderArgs["user"].(*models.User)
}


func setuser(rc *revel.Controller) revel.Result {

	var user *models.User
	if _, ok := rc.Session["uid"]; ok {
		uid, _ := strconv.ParseInt(rc.Session["uid"], 10, 0)
		user = models.GetUser(int(uid))
		fmt.Println("-----------------------USER THERE", user)
	}

	if user == nil {
		user = models.NewUser()
		rc.Session["uid"] = fmt.Sprintf("%d", user.Uid)
		fmt.Println("-----------------------USER NEW", user)
	}
	rc.RenderArgs["user"] = user

	return nil
}

func checkUser(rc *revel.Controller) revel.Result {

	fmt.Println("CHECKING IF FACEBOOKED IN")
	user := connected(rc)
	if user == nil || len(user.AccessToken) == 0 {
		rc.Flash.Error("Please log in first")
		return rc.Redirect(Application.Index)
	}
	return nil
}
*/
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
