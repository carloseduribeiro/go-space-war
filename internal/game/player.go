package game

import (
	"github.com/carloseduribeiro/go-space-war/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image             *ebiten.Image
	position          Vector
	game              *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite
	bounds := image.Bounds()
	halfWidth := float64(bounds.Dx()) / 2
	position := Vector{
		X: (screenWidth / 2) - halfWidth,
		Y: 500,
	}
	return &Player{
		image:             image,
		position:          position,
		game:              game,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() {
	speed := 6.0
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}
	p.laserLoadingTimer.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.IsReady() {
		p.laserLoadingTimer.Reset()
		bounds := p.image.Bounds()
		halfWidth := float64(bounds.Dx()) / 2
		halfHeigth := float64(bounds.Dy()) / 2
		laserSpawnPos := Vector{
			p.position.X + halfWidth,
			p.position.Y + halfHeigth,
		}
		laser := NewLaser(laserSpawnPos)
		p.game.AddLaser(laser)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.image, op)
}
