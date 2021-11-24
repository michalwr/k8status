package draw

import (
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/michalwr/k8status/internal/kubedata"
)

func MainLoop() {
	Kube:=kubedata.InitKubernetes()
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
	t := time.NewTicker(time.Second*10)
	w := 0
	h := 0	
	go func() {

	for {
		select {
		case <-t.C:

				s.Clear()
				Kube.RefreshData()
				DrawHeader(s, 0, 0, w-1, 3)
				DrawBody(s, 0, 4, w-1, h-1,Kube)
				DrawFooter()
			}
		}

	}()
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
			w, h = s.Size()
			s.Clear()
			DrawHeader(s, 0, 0, w-1, 3)
			DrawBody(s, 0, 4, w-1, h-1,Kube)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Key() == tcell.KeyCtrlR {

				s.Sync()
			}
		}

		//		DrawBody(s,0,4,w-1,h-1)
	}
}
