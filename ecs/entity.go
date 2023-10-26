package ecs

import (
	"go-ecs-game-engine/id"
	"log"
	"reflect"
)

type Entity struct {
	id.HasId
	Components map[reflect.Type]*Component
}

func (e *Entity) AddComponent(component Component, replaceComponentIfPresent bool) {
	compType := reflect.TypeOf(component)

	if !replaceComponentIfPresent && e.Components[compType] != nil {
		log.Printf("Component already exists for type %s", compType.String())
	}

	component = component.Initialise()
	e.Components[compType] = &component
}

func (e *Entity) RemoveComponentById(id uint64) {
	for compType, comp := range e.Components {
		compActual := *comp
		if compActual.GetId() == id {
			delete(e.Components, compType)
		}
	}
}

func NewEntity() Entity {
	return Entity{
		HasId: id.HasId{
			Id: id.GetNew(),
		},
		Components: make(map[reflect.Type]*Component),
	}
}
