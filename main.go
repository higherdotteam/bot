package main

import "github.com/gin-gonic/gin"

import "net/http"
import "fmt"
import "io/ioutil"
import "encoding/json"

type SlackResponse struct {
	// "type": "url_verification", "token": "QTT8T3f8VFmZ2MvH4m3jLBCh", "challenge": "q2TjwXAtQrPYt9ts3pxAmg3ryaNV0Cgo3OS9xJFlkghzuRkawyIe"
	Type      string `json:"type"`
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
		// https://higher.team/slack?code=2305278787c0&state=
		/*		    url = "https://" + team + ".slack.com/api/oauth.access?client_id=" +
		ENV['SLACK_CID'] + "&client_secret=" + ENV['SLACK_SECRET'] +
		"&code=" + code
		response = Net::HTTP.get_response(URI.parse(url)) */

		c.Redirect(http.StatusMovedPermanently, "http://higher.team/")
	})
	r.POST("/slack", func(c *gin.Context) {
		body := c.Request.Body
		x, _ := ioutil.ReadAll(body)

		result := SlackResponse{}
		json.Unmarshal(x, &result)

		fmt.Printf("%v", result)

		c.String(200, result.Challenge)
	})

	//r.RunTLS(":443", "/etc/letsencrypt/live/higher.team/fullchain.pem", "/etc/letsencrypt/live/higher.team/privkey.pem")
	r.Run()
}
