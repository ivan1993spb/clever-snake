package game

import (
	"github.com/ivan1993spb/snake-server/engine"
	"github.com/ivan1993spb/snake-server/playground"
)

type Game struct {
	world *World
}

type ErrCreateGame struct {
	Err error
}

func (e *ErrCreateGame) Error() string {
	return "cannot create game: " + e.Err.Error()
}

func NewGame(width, height uint8) (*Game, error) {
	area, err := engine.NewArea(width, height)
	if err != nil {
		return nil, &ErrCreateGame{Err: err}
	}

	scene, err := engine.NewScene(area)
	if err != nil {
		return nil, &ErrCreateGame{Err: err}
	}

	pg := playground.NewPlayground(scene)
	world := NewWorld(pg)

	return &Game{
		world: world,
	}, nil
}

func (g *Game) Start() {
	g.world.start()
}

func (g *Game) Stop() {
	g.world.stop()
}

func (g *Game) World() *World {
	return g.world
}

func (g *Game) Events(stop <-chan struct{}, buffer uint) <-chan Event {
	return g.world.Events(stop, buffer)
}
