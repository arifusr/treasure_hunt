package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type GameTreasureHunter struct {
	message   chan string
	quit      chan string
	field     [6][8]string
	uimessage string
	state     string
	hunter    []*Hunter
	treasure  [][2]int
}

func NewGameTreasureHunter(message chan string, quit chan string) *GameTreasureHunter {
	game := &GameTreasureHunter{message: message, quit: quit}
	hunter := NewHunter(game)
	game.hunter = []*Hunter{hunter}
	return game

}

func (game *GameTreasureHunter) FoundTreasure(treasure [2]int) {
	game.treasure = append(game.treasure, treasure)
}

func (game *GameTreasureHunter) Process(m string) {
	switch m {
	case "rerender":
		game.rerender()
	case "pengembaramove":
		game.pengembaramove()
	}

}

func (game *GameTreasureHunter) listenKeyboard() {
	for {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		if input.Text() == "quit" {
			game.quit <- "quit"
		}
		switch game.state {
		case "start":
			game.state = "prolog"
			game.uimessage = "seorang pengembara masuk ke dalam gua"
			game.message <- "rerender"
		case "prolog":
			game.message <- "pengembaramove"
		}

	}

}

func (game *GameTreasureHunter) printTreasure() {
	for _, x := range game.treasure {
		game.field[x[0]][x[1]] = "$"
	}
	game.uimessage = "disinilah emas itu ditemukan"
	game.rerender()
}

func (game *GameTreasureHunter) pengembaramove() {
	//pengambara hanya boleh move up 1 or more, move right 1 or more move down 1 or more
	allhunter := make([]*Hunter, 0)
	for _, h := range game.hunter {
		hunters := h.move()
		allhunter = append(allhunter, hunters...)
	}
	time.Sleep(1 * time.Second)
	if len(allhunter) > 0 {
		game.hunter = allhunter
		game.pengembaramove()
	} else {
		game.printTreasure()
	}

}

func (game *GameTreasureHunter) rerender() {
	fmt.Print("\033[H\033[2J")

	for _, x := range game.field {
		for _, y := range x {
			if y == "" {
				fmt.Print(" ")
			} else {
				fmt.Print(y)
			}
		}
		fmt.Println()
	}
	fmt.Println(game.uimessage)

}

func (game *GameTreasureHunter) Start() {
	go game.InitField()
	go game.listenKeyboard()

}

func (game *GameTreasureHunter) InitField() {
	for xi, x := range game.field {
		for yi := range x {
			game.field[xi][yi] = "."
			if xi == 0 || yi == 0 || yi == 7 || xi == 5 {
				game.field[xi][yi] = "#"
			}
			if xi == 2 && (yi == 2 || yi == 3 || yi == 4) {
				game.field[xi][yi] = "#"

			}
			if xi == 3 && (yi == 4 || yi == 6) {
				game.field[xi][yi] = "#"

			}
			if xi == 4 && yi == 2 {
				game.field[xi][yi] = "#"

			}
			if xi == 4 && yi == 1 {
				game.field[xi][yi] = "X"
				game.hunter[0].setcoordinate([2]int{xi, yi})
			}

		}
	}
	game.state = "start"
	game.uimessage = "disinilah pengembaraan itu dimulai"
	game.message <- "rerender"
}
