package game

import (
	"github.com/carloseduribeiro/go-space-war/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(spawnPos Vector) *Laser {
	image := assets.LaserSprite
	bounds := image.Bounds()
	halfWidth := float64(bounds.Dx()) / 2
	halfHeigth := float64(bounds.Dy()) / 2
	spawnPos.X -= halfWidth
	spawnPos.Y -= halfHeigth
	return &Laser{
		image:    image,
		position: spawnPos,
	}
}

func (l *Laser) Update() {
	speed := 7.0
	l.position.Y += -speed
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(l.position.X, l.position.Y)
	screen.DrawImage(l.image, op)
}

func (l *Laser) Collider() Rectangle {
	bounds := l.image.Bounds()
	return NewRectangle(l.position.X, l.position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
