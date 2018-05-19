package game

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/ivan1993spb/snake-server/objects/apple"
	"github.com/ivan1993spb/snake-server/world"
)

type Game struct {
	world  *world.World
	logger logrus.FieldLogger
}

type ErrCreateGame struct {
	Err error
}

func (e *ErrCreateGame) Error() string {
	return "cannot create game: " + e.Err.Error()
}

func NewGame(logger logrus.FieldLogger, width, height uint8) (*Game, error) {
	w, err := world.NewWorld(width, height)
	if err != nil {
		return nil, fmt.Errorf("cannot create game: %s", err)
	}

	return &Game{
		world:  w,
		logger: logger,
	}, nil
}

func (g *Game) Start(stop <-chan struct{}) {
	g.world.Start(stop)

	go func() {
		for event := range g.world.Events(stop, 32) {
			g.logger.Debugln("game event", event)
		}
	}()

	go func() {
		apple.NewApple(g.world)
		for event := range g.world.Events(stop, 32) {
			if event.Type == world.EventTypeObjectDelete {
				switch event.Payload.(type) {
				case *apple.Apple:
					apple.NewApple(g.world)
				}
			}
		}
	}()

	// TODO: Start observers.
}

func (g *Game) World() *world.World {
	return g.world
}

func (g *Game) Events(stop <-chan struct{}, buffer uint) <-chan world.Event {
	return g.world.Events(stop, buffer)
}
