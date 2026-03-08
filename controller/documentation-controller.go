package controller

import (
	"github.com/gin-gonic/gin"
)

type DocumentationController struct {
	title       string
	initialized bool
}

func (dc DocumentationController) Init(title string) DocumentationController {
	if dc.initialized {
		return dc
	}

	dc.title = title
	dc.initialized = true
	return dc
}

func (dc DocumentationController) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"title": dc.title})
	}
}
