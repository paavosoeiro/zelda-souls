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

type Player struct {
	*Entity
}

func (e *Entity) Update(deltaTime float64) {
	normVec := e.Direction.Normalize()

	movementVec := normVec.Scale(deltaTime)

	e.Position = e.Position.Sum(Vec2{e.Speed.X * movementVec.X, e.Speed.Y * movementVec.Y})
}
