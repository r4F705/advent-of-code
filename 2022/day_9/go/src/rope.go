package main

type Rope struct {
	Head  *Vector2
	Tail  *Vector2
	Knots []Vector2
}

func NewRope(initPos Vector2, bodyLength int) *Rope {

	if bodyLength < 2 {
		panic("Rope must have at least 2 knots")
	}

	knots := make([]Vector2, bodyLength)
	for i := 0; i < bodyLength; i++ {
		knots[i] = initPos
	}
	return &Rope{
		Knots: knots,
		Head:  &knots[0],
		Tail:  &knots[len(knots)-1],
	}
}

func (r *Rope) Move(moveOrder *MoveOrder, update func()) {
	newHead := r.Knots[0]

	switch moveOrder.Direction {
	case "U":
		for i := 0; i < moveOrder.Amount; i++ {
			newHead.Y -= 1
			r.moveKnots(newHead)
			update()
		}
	case "R":
		for i := 0; i < moveOrder.Amount; i++ {
			newHead.X += 1
			r.moveKnots(newHead)
			update()
		}
	case "D":
		for i := 0; i < moveOrder.Amount; i++ {
			newHead.Y += 1
			r.moveKnots(newHead)
			update()
		}
	case "L":
		for i := 0; i < moveOrder.Amount; i++ {
			newHead.X -= 1
			r.moveKnots(newHead)
			update()
		}
	}
}

func (r *Rope) moveKnots(newPos Vector2) {
	for i, knot := range r.Knots {
		if i == 0 {
			r.Knots[i] = newPos
		} else if knot.Distance(r.Knots[i-1]) >= 2 {
			direction := r.Knots[i-1].Difference(knot).Normalize()
			r.Knots[i].X += direction.X
			r.Knots[i].Y += direction.Y
		}
	}
}
