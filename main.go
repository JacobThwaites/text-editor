package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
    screen, err := tcell.NewScreen()
    if err != nil {
        log.Fatalf("Error creating screen: %v", err)
    }
    if err = screen.Init(); err != nil {
        log.Fatalf("Error initializing screen: %v", err)
    }
    defer screen.Fini()

    content := []rune{}
    cursorX := 0

    for {
        screen.Clear()
        for i, r := range content {
            screen.SetContent(i, 0, r, nil, tcell.StyleDefault)
        }
        screen.ShowCursor(cursorX, 0)
        screen.Show()

        ev := screen.PollEvent()
        switch ev := ev.(type) {
        case *tcell.EventKey:
            if ev.Key() == tcell.KeyCtrlC {
                return
            } else if ev.Key() == tcell.KeyBackspace || ev.Key() == tcell.KeyBackspace2 {
                if cursorX > 0 {
                    content = append(content[:cursorX-1], content[cursorX:]...)
                    cursorX--
                }
            } else if ev.Key() == tcell.KeyRune {
                content = append(content[:cursorX], append([]rune{ev.Rune()}, content[cursorX:]...)...)
                cursorX++
            }
        }
    }
}
