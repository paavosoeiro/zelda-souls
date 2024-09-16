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
	Player     Player
	lastUpdate time.Time
	Mobs       []Entity
}

func main() {
	mainLoop()
}

func mainLoop() {
	state := GameState{
		Player: Player{&Entity{Position: Vec2{
			X: width / 2,
			Y: height / 2,
		}}},
		lastUpdate: time.Now(),
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
	processInput(state.Player, inputChannel)

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
		state.Mobs[i].Update(deltaTime)
	}

}

func processInput(player Player, inputChannel chan rune) {
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
					if player.Position.Y > 0 {
						player.Position.Y--
					}
				case 's':
					if player.Position.Y < height-1 {
						player.Position.Y++
					}
				case 'a':
					if player.Position.X > 0 {
						player.Position.X--
					}
				case 'd':
					if player.Position.X < width-1 {
						player.Position.X++
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
			if x == int(state.Player.Position.X) && y == int(state.Player.Position.Y) {
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
