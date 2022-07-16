package fsmv

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

const (
	visualizePath     = "/fsm/visualize/"
	visualizeToken    = ":token"
	visualizeFullPath = visualizePath + visualizeToken
)

//go:embed template/* js/*
var f embed.FS

func InitRouter(router *gin.Engine) {

	tpl := template.Must(template.New("").ParseFS(f, "template/*.tmpl"))
	router.SetHTMLTemplate(tpl)
	router.StaticFS("/static", http.FS(f))

	c := &WebPageController{}
	router.GET(visualizeFullPath, c.Visualize)
}
