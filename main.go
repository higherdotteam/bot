package main

import "github.com/gin-gonic/gin"

//import "net/http"
import "fmt"

func main() {
	fmt.Println("start...")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/assets", "./assets")

	r.GET("/", func(c *gin.Context) {
		data := gin.H{"nav_dash": true}
		c.HTML(200, "index.tmpl", data)
	})
	r.POST("/slack", func(c *gin.Context) {
		//data := gin.H{"nav_dash": true}
		c.String(200, "wefwef")
	})

	//r.RunTLS(":443", "/etc/letsencrypt/live/higher.team/fullchain.pem", "/etc/letsencrypt/live/higher.team/privkey.pem")
	r.Run()
}
