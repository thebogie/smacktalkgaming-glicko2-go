package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	//"github.com/revel/revel/cache"
	"golang.org/x/oauth2"
	"mitchgottlieb.com/smacktalkgaming/app/models"
	//"mitchgottlieb.com/smacktalkgaming/app/routes"
	"net/http"
	"net/url"
	"strconv"
	//"time"
)

type Application struct {
	*revel.Controller
}

// The following keys correspond to a test application
// registered on Facebook, and associated with the loisant.org domain.
// You need to bind loisant.org to your machine with /etc/hosts to
// test the application locally.

var FACEBOOK = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	RedirectURL:  "",
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://graph.facebook.com/oauth/authorize",
		TokenURL: "https://graph.facebook.com/oauth/access_token",
	},
	//Scopes: []string{"user_first_name", "user_last_name"},
}

func (c Application) Index() revel.Result {

	//testthis, _ := revel.Config.StringDefault()

	FACEBOOK.ClientID, _ = revel.Config.String("appvars.clientid")
	FACEBOOK.ClientSecret, _ = revel.Config.String("appvars.clientsecret")
	FACEBOOK.RedirectURL, _ = revel.Config.String("appvars.redirecturl")

	//fmt.Println("WEB", testthis)

	u := c.Connected()
	me := map[string]interface{}{}

	//fmt.Println("U:", u.AccessToken)

	if u != nil && u.AccessToken != "" {

		resp, _ := http.Get("https://graph.facebook.com/me?access_token=" +
			url.QueryEscape(u.AccessToken))
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&me); err != nil {
			revel.ERROR.Println(err)
		}

		//logged in! we have their name. Find in the database their UUID for easy access to their player
		neo := new(models.Neo4jObj)
		u.PlayerUUID = neo.Create(&models.Player{Firstname: me["first_name"].(string), Surname: me["last_name"].(string)})
		me["playerUUID"] = u.PlayerUUID
		c.Session["playerUUID"] = u.PlayerUUID

		revel.INFO.Println("HERE", u.AccessToken)
		revel.INFO.Println("HERE", u.PlayerUUID)
		revel.INFO.Println("HERE", u.Uid)
	}
	authUrl := FACEBOOK.AuthCodeURL("foo")
	return c.Render(me, authUrl)
}

func (c Application) Auth(code string) revel.Result {

	tok, err := FACEBOOK.Exchange(oauth2.NoContext, code)
	if err != nil {
		revel.ERROR.Println(err)
		return c.Redirect(Application.Index)
	}
	fmt.Println("TOKEN", tok)
	user := c.Connected()
	user.AccessToken = tok.AccessToken

	return c.Redirect(Application.Index)
}

func (c Application) Logout() revel.Result {

	fmt.Println("SESSION", c.Session.Cookie().Raw)
	for k := range c.Session {
		delete(c.Session, k)
	}
	fmt.Println("SESSION AFTER", c.Session.Cookie().Raw)
	return c.Redirect(Application.Index)

}

func setuser(c *revel.Controller) revel.Result {
	var user *models.User
	if _, ok := c.Session["uid"]; ok {
		uid, _ := strconv.ParseInt(c.Session["uid"], 10, 0)
		user = models.GetUser(int(uid))
		fmt.Println("-----------------------USER EXIST", user)
	}
	if user == nil {
		user = models.NewUser()
		c.Session["uid"] = fmt.Sprintf("%d", user.Uid)
		fmt.Println("-----------------------USER NEW", user)
	}
	c.RenderArgs["user"] = user
	return nil
}

func (c Application) Connected() *models.User {
	return c.RenderArgs["user"].(*models.User)
}
