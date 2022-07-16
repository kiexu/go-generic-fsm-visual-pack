package fsmv

import (
	"fmt"
	fsm "github.com/kiexu/go-generic-fsm"
	"strings"
)

const (
	header    = "graph RL"
	indent    = "    " // indent with 4 spaces
	stateFmt  = "stateFmt: [%v]"
	indexFmt  = "idx: [%v]"
	prevState = "[previous]"
	currState = "[current]"
	blockFmt  = `%v["%v"]:::%v`
	edgeFmt   = `%v%v -- "[%v]" --> %v`
	cssBlock  = "c"
	css       = "classDef " + cssBlock + " fill:#FFF8DC,stroke:#696969,stroke-width:2px"
)

type mermaidFlowFormatter struct {
	fsm *fsm.FSM[string, string, string, string]
}

func NewMermaidFlowFormatter(fsm *fsm.FSM[string, string, string, string]) *mermaidFlowFormatter {
	return &mermaidFlowFormatter{fsm: fsm}
}

func (f *mermaidFlowFormatter) exportMermaid() (string, error) {
	if f.fsm == nil {
		return "", &IllegalWrapperErr{}
	}
	lines := make([]string, 0)
	lines = append(lines, header)
	for _, ec := range f.fsm.G().Adj() {
		if ec == nil {
			continue
		}
		for _, e := range ec.EList() {
			if e == nil {
				continue
			}
			lines = append(lines, f.exportEdge(e))
		}
	}
	lines = append(lines, css)
	return strings.Join(lines, "\n"), nil
}

func (f *mermaidFlowFormatter) exportNode(v *fsm.Vertex[string, string]) string {
	if v == nil {
		return ""
	}
	lines := make([]string, 0)
	lines = append(lines, fmt.Sprintf(stateFmt, v.StateVal())) // State name
	lines = append(lines, fmt.Sprintf(indexFmt, v.Idx()))      // Node indexFmt
	if f.fsm.CurrState() == v.StateVal() {
		lines = append(lines, currState)
	}
	if f.fsm.PrevState() == v.StateVal() {
		lines = append(lines, prevState)
	}
	return fmt.Sprintf(blockFmt, v.Idx(), strings.Join(lines, "<br>"), cssBlock)
}

func (f *mermaidFlowFormatter) exportEdge(e *fsm.Edge[string, string, string, string]) string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf(edgeFmt, indent, f.exportNode(e.FromV()), e.EventVal(), f.exportNode(e.ToV()))
}
