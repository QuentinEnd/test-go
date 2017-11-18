package main

import (
	tl "github.com/JoelOtter/termloop"
)

type Destination struct {
	*tl.Entity
}

type Player struct {
	*tl.Entity
}

func (destination *Destination) Tick(keyboard tl.Event) {
	if keyboard.Type == tl.EventKey {
		x, y := destination.Position()
		switch keyboard.Key {
		case tl.KeyArrowUp:
			destination.SetPosition(x, y-1)
		case tl.KeyArrowDown:
			destination.SetPosition(x, y+1)
		case tl.KeyArrowLeft:
			destination.SetPosition(x-1, y)
		case tl.KeyArrowRight:
			destination.SetPosition(x+1, y)
		}
	}
}

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
	})
	destination := Destination{tl.NewEntity(2, 2, 1, 1)}
	destination.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlue, Ch: '옷'})
	player := Player{tl.NewEntity(10, 10, 1, 1)}
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorYellow, Ch: '옷'})
	level.AddEntity(&destination)
	level.AddEntity(&player)
	game.Screen().SetLevel(level)
	game.Start()
}
