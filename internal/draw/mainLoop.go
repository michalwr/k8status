package draw

import (
	"log"
	"os"
	"github.com/gdamore/tcell/v2"
)
func MainLoop(){
    defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()
	quit := func() {
		s.Fini()
		os.Exit(0)
	}
	for {
		s.Show()
		ev := s.PollEvent()
		w:=0
		h:=0
		switch ev := ev.(type) {
		case *tcell.EventResize:
			w, h = ev.Size()
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			}
		}
		DrawHeader(s,0,0,w-1,2)
		DrawBody(s,0,4,w-1,h-1)
		DrawFooter()
	}
}