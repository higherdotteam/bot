package main

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"

func main() {
	fmt.Println("start...")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/assets", "./assets")

	r.POST("/", func(c *gin.Context) {
fmt.Println("111111")
		data := gin.H{"nav_dash": true}
		c.HTML(200, "index.tmpl", data)
	})
	r.GET("/list", func(c *gin.Context) {
		data := gin.H{"nav_dash": true}
		c.HTML(200, "list.tmpl", data)
	})
	r.POST("/order", func(c *gin.Context) {
		//_ := c.PostForm("stripeToken")
		c.Redirect(http.StatusMovedPermanently, "http://cordbouquet.tumblr.com/post/143915124906/thank-you")
	})

	r.RunTLS(":443", "/etc/letsencrypt/live/higher.team/fullchain.pem", "/etc/letsencrypt/live/higher.team/privkey.pem")
}
