package connections

import (
	"errors"
	"sync"

	"github.com/ivan1993spb/snake-server/game"
)

type ConnectionGroup struct {
	limit   int
	counter int
	mutex   *sync.RWMutex
	game    *game.Game
}

func NewConnectionGroup(connectionLimit int, g *game.Game) (*ConnectionGroup, error) {
	if connectionLimit > 0 {
		return &ConnectionGroup{
			limit: connectionLimit,
			mutex: &sync.RWMutex{},
			game:  g,
		}, nil
	}

	return nil, errors.New("cannot create connection group: invalid connection limit")
}

func (cg *ConnectionGroup) GetLimit() int {
	return cg.limit
}

func (cg *ConnectionGroup) GetCount() int {
	cg.mutex.RLock()
	defer cg.mutex.RUnlock()
	return cg.counter
}

// unsafeIsFull returns true if group is full
func (cg *ConnectionGroup) unsafeIsFull() bool {
	return cg.counter == cg.limit
}

func (cg *ConnectionGroup) IsFull() bool {
	cg.mutex.RLock()
	defer cg.mutex.RUnlock()
	return cg.unsafeIsFull()
}

// unsafeIsEmpty returns true if group is empty
func (cg *ConnectionGroup) unsafeIsEmpty() bool {
	return cg.counter == 0
}

func (cg *ConnectionGroup) IsEmpty() bool {
	cg.mutex.RLock()
	defer cg.mutex.RUnlock()
	return cg.unsafeIsEmpty()
}

func (cg *ConnectionGroup) Run(f func(game *game.Game)) error {
	cg.mutex.Lock()
	if cg.unsafeIsFull() {
		cg.mutex.Unlock()
		return errors.New("add connection to group: group is full")
	}
	cg.counter += 1
	cg.mutex.Unlock()

	f(cg.game)

	cg.mutex.Lock()
	cg.counter -= 1
	cg.mutex.Unlock()
	return nil
}
