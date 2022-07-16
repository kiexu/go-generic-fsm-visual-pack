package fsmv

import (
	"github.com/gin-gonic/gin"
	fsm "github.com/kiexu/go-generic-fsm"
	"net/http"
)

const (
	mermaidRemoteSrc = `https://cdn.jsdelivr.net/npm/mermaid/dist/mermaid.min.js`
	mermaidNativeSrc = `/static/js/mermaid.min.js`
)

type (
	WebPageController struct {
	}
	VisualizeReq struct {
		Token string `form:"token"`
	}
)

func (wp *WebPageController) Visualize(c *gin.Context) {

	token := c.Param("token")
	if token == "" {
		c.JSON(http.StatusOK, &gin.H{
			"errMsg": "illegal token",
		})
		return
	}

	if fsm.GetVisualizer() == nil {
		c.JSON(http.StatusOK, &gin.H{
			"errMsg": "visualizer not init",
		})
		return
	}

	webPageVisualizer, ok := fsm.GetVisualizer().(*WebPageVisualizer)
	if !ok {
		c.JSON(http.StatusOK, &gin.H{
			"errMsg": "incompatible visualizer type",
		})
		return
	}

	output, err := webPageVisualizer.exportMermaid(token)
	if err != nil {
		c.JSON(http.StatusOK, &gin.H{
			"errMsg": "visualizer internal error: " + err.Error(),
		})
		return
	}

	var src string
	if GetGlobalConfig().NativeScript {
		src = mermaidNativeSrc
	} else {
		src = mermaidRemoteSrc
	}

	c.HTML(http.StatusOK, "visualization.tmpl", gin.H{
		"output": output,
		"src":    src,
	})

}
