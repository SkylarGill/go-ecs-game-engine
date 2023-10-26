package ecs

type Component interface {
	GetId() uint64
	Initialise() Component
}
