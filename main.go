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
	mobX         int
	mobY         int
	mobDirection int
}

func main() {
	mainLoop()
}

func mainLoop() {
	state := GameState{
		playerX:      width / 2,
		playerY:      height / 2,
		mobX:         0,
		mobY:         height / 2,
		mobDirection: 1,
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
			update(&state)
			render(state)

		case _, ok := <-inputChannel:
			if !ok {
				fmt.Println("Saindo do jogo...")
				return
			}
			//switch char {
			//case 'w':
			//	if state.playerY > 0 {
			//		state.playerY--
			//	}
			//case 's':
			//	if state.playerY < height-1 {
			//		state.playerY++
			//	}
			//case 'a':
			//	if state.playerX > 0 {
			//		state.playerX--
			//	}
			//case 'd':
			//	if state.playerX < width-1 {
			//		state.playerX++
			//	}
			//}
		}

	}

}

func update(state *GameState) {
	state.mobX += state.mobDirection
	if state.mobX == width-1 || state.mobX == 0 {
		state.mobDirection *= -1
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

	//for char := range inputChannel {
	//	switch char {
	//	case 'w':
	//		if state.playerY > 0 {
	//			state.playerY--
	//		}
	//	case 's':
	//		if state.playerY < height-1 {
	//			state.playerY++
	//		}
	//	case 'a':
	//		if state.playerX > 0 {
	//			state.playerX--
	//		}
	//	case 'd':
	//		if state.playerX < width-1 {
	//			state.playerX++
	//		}
	//	}
	//}

	//if char, key, err := keyboard.GetKey(); err == nil {
	//	switch key {
	//	case keyboard.KeyArrowUp:
	//		if state.playerY > 0 {
	//			state.playerY--
	//		}
	//	case keyboard.KeyArrowDown:
	//		if state.playerY < height-1 {
	//			state.playerY++
	//		}
	//	case keyboard.KeyArrowLeft:
	//		if state.playerX > 0 {
	//			state.playerX--
	//		}
	//	case keyboard.KeyArrowRight:
	//		if state.playerX < width-1 {
	//			state.playerX++
	//		}
	//	default:
	//
	//	}
	//	if char == 'q' {
	//		fmt.Println("Saindo do jogo...")
	//	}
	//}
}

func render(state GameState) {
	clearScreen()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == state.playerX && y == state.playerY {
				fmt.Print("@")
			} else if x == state.mobX && y == state.mobY {
				fmt.Print("M")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
