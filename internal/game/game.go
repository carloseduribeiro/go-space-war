package game

import (
	"fmt"
	"image/color"
	"log/slog"

	"github.com/carloseduribeiro/go-space-war/internal/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	player           *Player
	lasers           []*Laser
	meteors          []*Meteor
	meteorSpawnTimer *Timer
	score            int
}

func NewGame() *Game {
	lasers := make([]*Laser, 0)
	g := &Game{
		lasers:           lasers,
		meteorSpawnTimer: NewTimer(24),
	}
	player := NewPlayer(g)
	g.player = player
	return g
}

func (g *Game) Update() error {
	g.player.Update()
	for _, l := range g.lasers {
		l.Update()
	}
	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()
		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}
	for _, m := range g.meteors {
		m.Update()
	}
	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			slog.Info("Perdeu! Game Reset.")
			g.Reset()
		}
	}
	for i, m := range g.meteors {
		for j, l := range g.lasers {
			if m.Collider().Intersects(l.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score += 1
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	for _, l := range g.lasers {
		l.Draw(screen)
	}
	for _, m := range g.meteors {
		m.Draw(screen)
	}
	drawOp := text.DrawOptions{}
	drawOp.GeoM.Translate(10, 10)
	drawOp.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, fmt.Sprintf("Points: %d", g.score), assets.FontUi, &drawOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeigth
}

func (g *Game) AddLaser(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
	g.meteorSpawnTimer.Reset()
}
