package draw

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/michalwr/k8status/internal/kubedata"
)

func drawPods(s tcell.Screen, x1 int, y1 int, x2 int, y2 int,Kube kubedata.KubernetesCluster ) {
	printNodeLine := 1
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}
	style = tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorBlack)
	drawText(s, x1+1, y1+printNodeLine, x1+50, y2+printNodeLine, style, "Pod Name")
	drawText(s, x1+50, y1+printNodeLine, x1+80, y2+printNodeLine, style, "NameSpace")
	drawText(s, x1+80, y1+printNodeLine, x1+120, y2+printNodeLine, style, "Container")
	drawText(s, x1+120, y1+printNodeLine, x1+140, y2+printNodeLine, style, "Exit Code")
	drawText(s, x1+140, y1+printNodeLine, x1+160, y2+printNodeLine, style, "Restarts")
	printNodeLine++

	for _, v := range Kube.PodProps {
		if v.Status == "Running" && printNodeLine<y2-5{
			for _,vd := range v.ContainerStatuses{
			if !vd.Status {
			style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
			drawText(s, x1+1, y1+printNodeLine, x1+60, y2+printNodeLine, style, v.Name)
			style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
			drawText(s, x1+50, y1+printNodeLine, x1+80, y2+printNodeLine, style, v.Namespace)
			style = tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
			drawText(s, x1+80, y1+printNodeLine, x1+120, y2+printNodeLine, style, vd.Name)
			style = tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
			drawText(s, x1+120, y1+printNodeLine, x1+140, y2+printNodeLine, style, strconv.Itoa(int(vd.ExitCode)))
			style = tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
			drawText(s, x1+140, y1+printNodeLine, x1+160, y2+printNodeLine, style, strconv.Itoa(vd.RestartCount))
			printNodeLine++
			}
			}
		}
	}
	s.Show()
}
func drawNodes(s tcell.Screen, x1 int, y1 int, x2 int, y2 int,Kube  kubedata.KubernetesCluster) {
	printNodeLine := 1
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}
	for _, v := range Kube.NodeProps {
		style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
		drawText(s, x1+1, y1+printNodeLine, x1+30, y2+printNodeLine, style, v.Name)
		if v.Status == "True" {
			style = tcell.StyleDefault.Foreground(tcell.ColorDarkGreen).Background(tcell.ColorBlack)
			drawText(s, x1+30, y1+printNodeLine, x1+40, y2+printNodeLine, style, "Ready")
		} else {
			style = tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
			drawText(s, x1+30, y1+printNodeLine, x1+40, y2+printNodeLine, style, "NotReady")
		}
		printNodeLine++
	}
	s.Show()
}

func DrawBody(s tcell.Screen, x1 int, y1 int, x2 int, y2 int,Kube   kubedata.KubernetesCluster) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}
	s.Show()
	drawNodes(s, 2, 5, 50, y2-1,Kube)
	drawPods(s, 52, 5, x2-1, y2-1,Kube)
}
