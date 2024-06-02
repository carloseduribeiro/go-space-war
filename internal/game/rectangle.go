package game

type Rectangle struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewRectangle(x, y, width, height float64) Rectangle {
	return Rectangle{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (r Rectangle) maxX() float64 {
	return r.X + r.Width
}

func (r Rectangle) maxY() float64 {
	return r.Y + r.Height
}

func (r Rectangle) Intersects(other Rectangle) bool {
	return r.X <= other.maxX() &&
		other.X <= r.maxX() &&
		r.Y <= other.maxY() &&
		other.Y <= r.maxY()
}
