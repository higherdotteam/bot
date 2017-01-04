package main

import "github.com/gin-gonic/gin"

import "net/http"
import "fmt"
import "io/ioutil"
import "encoding/json"
import "os"

type SlackResponse struct {
	Type      string `json:"type"`
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
}

type SlackAccess struct {
	Type      string `json:"access_token"`
	Type      string `json:"user_id"`
	Type      string `json:"team_name"`
	Type      string `json:"team_id"`
	Type      string `json:""`
	Type      string `json:"access_token"`
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
}

func main() {
	fmt.Println("start...")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/assets", "./assets")

	r.GET("/", func(c *gin.Context) {
		data := gin.H{"nav_dash": true}
		c.HTML(200, "index.tmpl", data)
	})
	r.GET("/slack", func(c *gin.Context) {
		r, _ := http.Get("https://slack.com/api/oauth.access?client_id=" + os.Getenv("HDT_CID") + "&client_secret=" +
			os.Getenv("HDT_SECRET") + "&code=" + c.Query("code"))
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))

		/* {"ok":true,"access_token":"6f","scope":"identify,bot","user_id":"U1","team_name":"higher.team","team_id":"TL","bot":{"bot_user_id":"UP","bot_access_token":"zg"}} */

		c.Redirect(http.StatusMovedPermanently, "http://higher.team/")
	})
	r.POST("/slack", func(c *gin.Context) {

		/* {"token":"h","team_id":"T035N23CL","api_app_id":"A3M072NLA",
		"event":{"type":"message","user":"U035LF6C1","text":"fwefef",
		         "ts":"1483496552.000007","channel":"D3MN1KCHM",
						 "event_ts":"1483496552.000007"},
		"type":"event_callback","authed_users":["U3LV4ST2P"]} */

		body := c.Request.Body
		x, _ := ioutil.ReadAll(body)

		result := SlackResponse{}
		json.Unmarshal(x, &result)

		fmt.Printf("%v", result)

		c.String(200, result.Challenge)
	})
	r.POST("/slack/challenge", func(c *gin.Context) {
		body := c.Request.Body
		x, _ := ioutil.ReadAll(body)

		result := SlackResponse{}
		json.Unmarshal(x, &result)

		fmt.Printf("%v", result)

		c.String(200, result.Challenge)
	})

	r.RunTLS(":443", "/etc/letsencrypt/live/higher.team/fullchain.pem", "/etc/letsencrypt/live/higher.team/privkey.pem")
	//r.Run()
}
