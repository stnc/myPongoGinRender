my Pongo Gin Render

=========

Package myPongoGinRender is a template renderer that can be used with the Gin web
framework https://github.com/gin-gonic/gin it uses the Pongo2 template library
https://github.com/flosch/pongo2

Compatible with pongo version 4

pongo2 is a Django-syntax like templating-language (official website).

<strong>Orginal Code By Andrejs Cainikovs - Rob van der Linde (https://gitlab.com/go-box/pongo2gin) </strong>

## Here is Compatible with pongo version 4 &  version 5
 [myPongoGinRender](https://github.com/stnc/myPongoGinRender/) - pongo2 gin minimal framework stability renderer / Compatible with pongo version 2 
 [myPongoGinRender](https://github.com/stnc/myPongoGinRender/tree/main/v4) -  Compatible with pongo version 4
 [myPongoGinRender](https://github.com/stnc/myPongoGinRender/tree/main/v5) -  Compatible with pongo version 5

# please don't forget to give stars :)

## Installation  

`go get "github.com/stnc/myPongoGinRender"`

Requirements
------------

Requires Gin 1.14 or higher and Pongo2.

Usage
-----

To use myPongoGinRender you need to set your router.HTMLRenderer to a new renderer
instance, this is done after creating the Gin router when the Gin application
starts up. This assumes templates will be located in the "templates"
directory, or you can use myPongoGinRender.TemplatePath("templates") to specify a custom location.

To render templates from a route, call c.HTML just as you would with
regular Gin templates, the only difference is that you pass template
data as a pongo2.Context instead of gin.H type.


![Screen](https://raw.githubusercontent.com/stnc/myPongoGinRender/master/example/ginScreen.png)

Basic Example
-------------

```go

package main

import (
	"log"
	"net/http"
	"github.com/stnc/myPongoGinRender"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

//GetAllData all list
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

```

HTML 
----------------


```html

<h1> {{ title }}</h1>

{% for post in posts%}
<ul>
  <li>{{post}}</li>
</ul>
{% endfor %}

```

Template Caching
----------------

Templates will be cached if the current Gin Mode is set to anything but "debug",
this means the first time a template is used it will still load from disk, but
after that the cached template will be used from memory instead.

If he Gin Mode is set to "debug" then templates will be loaded from disk on
each request.

Caching is implemented by the Pongo2 library itself.

GoDoc
-----


Specials Thanks
-----
https://gitlab.com/go-box/pongo2gin

https://zhuanlan.zhihu.com/p/105956154  


