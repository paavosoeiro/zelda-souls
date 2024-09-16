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
	normVec := e.Direction.Normalize()

	movementVec := normVec.Scale(deltaTime)

	e.Position = e.Position.Sum(Vec2{e.Speed.X * movementVec.X, e.Speed.Y * movementVec.Y})
}
