package main

type Sprite struct {
	pixels [3]int
}

func NewSprite() *Sprite {
	return &Sprite{
		pixels: [3]int{0, 0, 0},
	}
}

func (s *Sprite) Move(x int) {
	s.pixels[0] = x
	s.pixels[1] = x + 1
	s.pixels[2] = x + 2
}
