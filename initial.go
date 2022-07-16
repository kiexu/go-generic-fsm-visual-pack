package fsmv

import (
	"fmt"
	"github.com/gin-gonic/gin"
	fsm "github.com/kiexu/go-generic-fsm"
	"io"
	"sync"
)

// InitFSMVisualPack init Visual Pack
func InitFSMVisualPack(conf *Config) (err error) {

	defer func() {
		if err != nil {
			// Release resources
			fsm.SetVisualizer(nil)
		}
	}()

	if conf == nil || conf.Port <= 0 {
		return &IllegalConfigErr{}
	}

	// Disable all log
	gin.DefaultWriter = io.Discard

	if conf.Mode != nil {
		gin.SetMode(*conf.Mode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// store native config
	SetGlobalConfig(conf)

	fsm.SetVisualizer(&WebPageVisualizer{
		tokenMap: make(map[string]fsm.VisualGenerator),
		mutex:    sync.Mutex{},
	})

	r := gin.New()
	InitRouter(r)
	go func() {
		_ = r.Run(fmt.Sprintf(":%v", conf.Port))
	}()

	return nil
}
