package draw

import (

	"github.com/gdamore/tcell/v2"
	"github.com/michalwr/k8status/internal/kubedata"
)
func drawPods(s tcell.Screen,x1 int, y1 int, x2 int, y2 int) {
	KubeNodes:=kubedata.GetNodeStatus()
	printNodeLine:=1
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}
	for _,v := range KubeNodes{
		style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
		drawText(s, x1+1, y1+printNodeLine, x1+30, y2+printNodeLine, style, v.Name)
		if v.Status == "True" {
			style = tcell.StyleDefault.Foreground(tcell.ColorDarkGreen).Background(tcell.ColorBlack)
			drawText(s, x1+30, y1+printNodeLine, x1+40, y2+printNodeLine, style, "Ready")
		}else{
			style = tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
			drawText(s, x1+30, y1+printNodeLine, x1+40, y2+printNodeLine, style, "NotReady")			
		}
printNodeLine++
	}
}
func drawNodes(s tcell.Screen,x1 int, y1 int, x2 int, y2 int) {
	KubeNodes:=kubedata.GetNodeStatus()
	printNodeLine:=1
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}
	for _,v := range KubeNodes{
		style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
		drawText(s, x1+1, y1+printNodeLine, x1+30, y2+printNodeLine, style, v.Name)
		if v.Status == "True" {
			style = tcell.StyleDefault.Foreground(tcell.ColorDarkGreen).Background(tcell.ColorBlack)
			drawText(s, x1+30, y1+printNodeLine, x1+40, y2+printNodeLine, style, "Ready")
		}else{
			style = tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
			drawText(s, x1+30, y1+printNodeLine, x1+40, y2+printNodeLine, style, "NotReady")			
		}
printNodeLine++
	}
}

func DrawBody(s tcell.Screen,x1 int, y1 int, x2 int, y2 int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}
	drawNodes(s,2,5,50,y2-1)
	drawPods(s,52,5,x2-1,y2-1)

}

