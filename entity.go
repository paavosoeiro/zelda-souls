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

type UpdateMessage struct {
	DeltaTime  float64
	GridWidth  float64
	GridHeight float64
}

func (e *Entity) Update(deltaTime float64, gridWidth, gridHeight float64) {
	normVec := e.Direction.Normalize()

	movementVec := normVec.Scale(deltaTime)

	e.Position = e.Position.Sum(Vec2{e.Speed.X * movementVec.X, e.Speed.Y * movementVec.Y})

	if e.Position.X < 0 {
		e.Position.X = 0
	} else if e.Position.X >= gridWidth {
		e.Position.X = gridWidth - 1
		e.Direction.X *= -1
	}

	if e.Position.Y < 0 {
		e.Position.Y = 0
	} else if e.Position.Y >= gridHeight {
		e.Position.Y = gridHeight - 1
		e.Direction.Y *= -1
	}
}

func (e *Entity) StartUpdateChannel(updateChan <-chan UpdateMessage) {
	for msg := range updateChan {
		e.Update(msg.DeltaTime, msg.GridWidth, msg.GridHeight)
	}
}
