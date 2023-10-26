package components

import (
	"go-ecs-game-engine/ecs"
	"go-ecs-game-engine/id"
)

type Position struct {
	id.HasId
	X float32
	Y float32
}

func (p Position) Initialise() ecs.Component {
	p.Id = id.GetNew()
	return p
}

func (p Position) GetId() uint64 {
	return p.Id
}
