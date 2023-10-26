package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"go-ecs-game-engine/components"
	"go-ecs-game-engine/ecs"
	"image/color"
	"log"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 30, 255})

	vector.DrawFilledCircle(screen, 50, 50, 25, color.RGBA{100, 255, 100, 255}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Go Pool Game")

	entity := ecs.NewEntity()
	entity = ecs.NewEntity()
	entity = ecs.NewEntity()
	comp := components.Position{}
	entity.AddComponent(comp, false)

	entity.RemoveComponentById(3)

	log.Print(len(entity.Components))
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
