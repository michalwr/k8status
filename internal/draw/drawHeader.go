package draw

import "github.com/gdamore/tcell/v2"

func DrawHeader(s tcell.Screen,x1 int, y1 int, x2 int, y2 int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorBlack)
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}
	label:="K8s Status"
	drawText(s, x2/2, y1+1, x2, y2+1, style, label)
}
