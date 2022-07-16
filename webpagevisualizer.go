package fsmv

import (
	"github.com/google/uuid"
	fsm "github.com/kiexu/go-generic-fsm"
	"sync"
)

const (
	retryNumCeil = 5 // Chance to conflict
)

var _ fsm.Visualizer = new(WebPageVisualizer)

type WebPageVisualizer struct {
	tokenMap map[string]fsm.VisualGenerator
	mutex    sync.Mutex
}

// Open add generator to tokenMap
func (w *WebPageVisualizer) Open(wrapper *fsm.VisualizeStartWrapper) error {

	w.mutex.Lock()
	defer w.mutex.Unlock()

	if wrapper == nil || wrapper.VisualGen == nil {
		return &IllegalWrapperErr{}
	}

	var token string
	for i := 0; i < retryNumCeil; i += 1 {
		token = uuid.NewString()
		if _, ok := w.tokenMap[token]; !ok {
			w.tokenMap[token] = wrapper.VisualGen
			path := visualizePath + token
			wrapper.Path = &path
			wrapper.Token = &token
			return nil
		}
	}

	return &UnknownErr{}
}

// Close delete generator from tokenMap
func (w *WebPageVisualizer) Close(wrapper *fsm.VisualizeStopWrapper) error {

	w.mutex.Lock()
	defer w.mutex.Unlock()

	if wrapper.Token == nil {
		return &IllegalWrapperErr{}
	}

	if _, ok := w.tokenMap[*wrapper.Token]; !ok {
		return &IllegalWrapperErr{}
	}

	delete(w.tokenMap, *wrapper.Token)

	return nil
}

// exportMermaid 导出mermaid格式
func (w *WebPageVisualizer) exportMermaid(token string) (code string, err error) {

	w.mutex.Lock()
	defer w.mutex.Unlock()

	f, ok := w.tokenMap[token]
	if !ok {
		return "", &IllegalWrapperErr{}
	}

	code, err = NewMermaidFlowFormatter(f()).exportMermaid()
	return
}
