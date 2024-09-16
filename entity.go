package main

type Direction struct {
	Vec2
}

type Entity struct {
	Position  Vec2
	Direction Direction
	Speed     Vec2
	Sprite    rune
}

func (e *Entity) update(deltaTime float64) {
	//move := e.Speed * deltaTime
	//e.Position.Y += move * e.Direction.Y
	//
	//if e.Position.Y >= float64(height-1) || e.Position.Y == 0 {
	//	e.Direction.Y *= -1
	//}
	//
	//if e.Position.Y < 0 {
	//	e.Position.Y = 0
	//} else if e.Position.Y >= float64(height) {
	//	e.Position.Y = float64(height - 1)
	//}

	normVec := e.Direction.Normalize()

	movementVec := normVec.Scale(deltaTime)

	e.Position = e.Position.Sum(Vec2{e.Speed.X * movementVec.X, e.Speed.Y * movementVec.Y})
}
