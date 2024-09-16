package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"time"
)

const (
	width  = 20
	height = 20
	fps    = 10
)

type GameState struct {
	playerX      int
	playerY      int
	mobDirection int
	lastUpdate   time.Time
	Mobs         []Entity
}

func main() {
	mainLoop()
}

func mainLoop() {
	state := GameState{
		playerX:      width / 2,
		playerY:      height / 2,
		mobDirection: 1,
		lastUpdate:   time.Now(),
		Mobs: []Entity{
			{Position: Vec2{X: 5, Y: 5}, Direction: Direction{Vec2{X: 1, Y: 1}}, Speed: Vec2{X: 0.5, Y: 0.5}, Sprite: 'E'},
		},
	}

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := keyboard.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	ticker := time.NewTicker(time.Second / fps)
	defer ticker.Stop()

	inputChannel := make(chan rune)
	processInput(&state, inputChannel)

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			deltaTime := now.Sub(state.lastUpdate).Seconds()
			state.lastUpdate = now

			update(&state, deltaTime)
			render(state)

		case _, ok := <-inputChannel:
			if !ok {
				fmt.Println("Saindo do jogo...")
				return
			}
		}

	}

}

func update(state *GameState, deltaTime float64) {

	for i := range state.Mobs {
		//move := state.Mobs[i].Speed * deltaTime
		//state.Mobs[i].Position.Y += move * state.Mobs[i].Direction.Y
		//
		//if state.Mobs[i].Position.Y >= float64(height-1) || state.Mobs[i].Position.Y == 0 {
		//	state.Mobs[i].Direction.Y *= -1
		//}
		//
		//if state.Mobs[i].Position.Y < 0 {
		//	state.Mobs[i].Position.Y = 0
		//} else if state.Mobs[i].Position.Y >= float64(height) {
		//	state.Mobs[i].Position.Y = float64(height - 1)
		//}

		state.Mobs[i].update(deltaTime)
	}

}

func processInput(state *GameState, inputChannel chan rune) {
	go func() {
		for {
			if char, key, err := keyboard.GetKey(); err == nil {
				inputChannel <- char
				if key == keyboard.KeyEsc || char == 'q' {
					close(inputChannel)
					return
				}
				switch char {
				case 'w':
					if state.playerY > 0 {
						state.playerY--
					}
				case 's':
					if state.playerY < height-1 {
						state.playerY++
					}
				case 'a':
					if state.playerX > 0 {
						state.playerX--
					}
				case 'd':
					if state.playerX < width-1 {
						state.playerX++
					}
				}
			}
		}
	}()

}

func render(state GameState) {
	var buffer []rune
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == state.playerX && y == state.playerY {
				buffer = append(buffer, '@')
			} else if len(state.Mobs) > 0 {
				for _, mob := range state.Mobs {
					if x == int(mob.Position.X) && y == int(mob.Position.Y) {
						buffer = append(buffer, mob.Sprite)
					} else {
						buffer = append(buffer, '.')
					}
				}
			} else {
				buffer = append(buffer, '.')
			}
		}
		buffer = append(buffer, '\n')
	}
	clearScreen()
	fmt.Print(string(buffer))
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
