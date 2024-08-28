package main

type Rope struct {
	Head Vector2
	Body []Vector2
	Tail Vector2
}

func NewRope(initPos Vector2, bodyLength int) *Rope {
	body := make([]Vector2, bodyLength)
	for i := 0; i < bodyLength; i++ {
		body[i] = initPos
	}
	return &Rope{
		Head: initPos,
		Body: body,
		Tail: initPos,
	}
}

func (r *Rope) Move(moveOrder *MoveOrder, update func()) {
	newHead := r.Head

	switch moveOrder.Direction {
	case "U":
		for i := 0; i < moveOrder.Amount; i++ {
			newHead.Y -= 1
			if len(r.Body) > 0 {
				r.moveTail(r.Body[len(r.Body)-1])
			} else {
				r.moveTail(newHead)
			}
			r.moveBody(newHead)
			r.moveHead(newHead)
			update()
		}
	case "R":
		for i := 0; i < moveOrder.Amount; i++ {
			newHead.X += 1
			if len(r.Body) > 0 {
				r.moveTail(r.Body[len(r.Body)-1])
			} else {
				r.moveTail(newHead)
			}
			r.moveBody(newHead)
			r.moveHead(newHead)
			update()
		}
	case "D":
		for i := 0; i < moveOrder.Amount; i++ {
			newHead.Y += 1
			if len(r.Body) > 0 {
				r.moveTail(r.Body[len(r.Body)-1])
			} else {
				r.moveTail(newHead)
			}
			r.moveBody(newHead)
			r.moveHead(newHead)
			update()
		}
	case "L":
		for i := 0; i < moveOrder.Amount; i++ {
			newHead.X -= 1
			if len(r.Body) > 0 {
				r.moveTail(r.Body[len(r.Body)-1])
			} else {
				r.moveTail(newHead)
			}
			r.moveBody(newHead)
			r.moveHead(newHead)
			update()
		}
	}
}

func (r *Rope) moveHead(newPos Vector2) {
	r.Head = newPos
}

func (r *Rope) moveBody(newPos Vector2) {
	if len(r.Body) > 0 {
		var moveMade Vector2
		for i := 0; i < len(r.Body); i++ {
			part := &r.Body[i]
			if i == 0 && newPos.Distance(*part) >= 2 {
				moveMade = newPos.Difference(*part)
				part.X = r.Head.X
				part.Y = r.Head.Y
				newPos = *part
			} else if i > 0 && newPos.Distance(*part) >= 2 {
				part.X += moveMade.X
				part.Y += moveMade.Y
				newPos = *part
			}
		}
	}

}

func (r *Rope) moveTail(newPos Vector2) {

	// if len(r.Body) > 0 {
	// 	if newPos.Distance(r.Tail) >= 2 {
	// 		r.Tail = newPos
	// 	}
	// } else {
	// 	if newPos.Distance(r.Tail) >= 2 {
	// 		r.Tail = r.Head
	// 	}
	// }

}
