package main

import (
	"log"
	"net/http"
myPongoGinRender "github.com/stnc/myPongoGinRender/v5"

	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
)

func GetAllData(c *gin.Context) {
	posts := []string{
		"Larry Ellison",
		"Carlos Slim Helu",
		"Mark Zuckerberg",
		"Amancio Ortega ",
		"Jeff Bezos",
		" Warren Buffet ",
		"Bill Gates",
		"selman tunç",
	}
	// Call the HTML method of the Context to render a template
	c.HTML(http.StatusOK, "index.html",
		pongo2.Context{
			"title": "hello pongo",
			"posts": posts,
		},
	)
}
func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(gin.Recovery())
	r.HTMLRender = myPongoGinRender.TemplatePath("templates")
	r.GET("/", GetAllData)
	log.Fatal(r.Run(":8888"))
}
