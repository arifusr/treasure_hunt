package main

type Hunter struct {
	game          *GameTreasureHunter
	coordinate    [2]int
	oldmove       string
	foundTreasure bool
}

func NewHunter(game *GameTreasureHunter) *Hunter {
	return &Hunter{game: game}
}

func (hunter *Hunter) setcoordinate(coordinate [2]int) {
	hunter.coordinate = coordinate
}

func (hunter *Hunter) move() []*Hunter {
	allhunter := make([]*Hunter, 0)
	if hunter.oldmove == "" {
		h := hunter.moveUp()
		if h != nil {
			allhunter = append(allhunter, h)
		}
	} else if hunter.oldmove == "up" {
		h := hunter.moveUp()
		if h != nil {
			allhunter = append(allhunter, h)
		}
		hr := hunter.moveRight()
		if hr != nil {
			allhunter = append(allhunter, hr)
		}
	} else if hunter.oldmove == "right" {
		hr := hunter.moveRight()
		if hr != nil {
			allhunter = append(allhunter, hr)
		}
		hd := hunter.moveDown()
		if hd != nil {
			allhunter = append(allhunter, hd)
		}
	} else if hunter.oldmove == "down" {
		hunter.game.FoundTreasure(hunter.coordinate)
		hd := hunter.moveDown()
		if hd != nil {
			allhunter = append(allhunter, hd)
		}
	}
	return allhunter

}

func (hunter *Hunter) moveUp() *Hunter {
	up := hunter.coordinate[0] - 1
	if hunter.game.field[up][hunter.coordinate[1]] == "." {
		hunter.game.field[hunter.coordinate[0]][hunter.coordinate[1]] = "."
		hunter.game.field[up][hunter.coordinate[1]] = "X"
		hunter.game.rerender()
		newhunter := NewHunter(hunter.game)
		newhunter.setcoordinate([2]int{up, hunter.coordinate[1]})
		newhunter.oldmove = "up"
		return newhunter
	}
	return nil
}

func (hunter *Hunter) moveRight() *Hunter {
	right := hunter.coordinate[1] + 1
	if hunter.game.field[hunter.coordinate[0]][right] == "." {
		hunter.game.field[hunter.coordinate[0]][hunter.coordinate[1]] = "."
		hunter.game.field[hunter.coordinate[0]][right] = "X"
		hunter.game.rerender()
		newhunter := NewHunter(hunter.game)
		newhunter.setcoordinate([2]int{hunter.coordinate[0], right})
		newhunter.oldmove = "right"
		return newhunter
	}
	return nil
}

func (hunter *Hunter) moveDown() *Hunter {
	down := hunter.coordinate[0] + 1
	if hunter.game.field[down][hunter.coordinate[1]] == "." {
		hunter.foundTreasure = true
		if hunter.oldmove == "down" {
			hunter.game.field[hunter.coordinate[0]][hunter.coordinate[1]] = "$"
		} else {
			hunter.game.field[hunter.coordinate[0]][hunter.coordinate[1]] = "."
		}
		hunter.game.field[down][hunter.coordinate[1]] = "X"
		hunter.game.rerender()
		newhunter := NewHunter(hunter.game)
		newhunter.setcoordinate([2]int{down, hunter.coordinate[1]})
		newhunter.oldmove = "down"
		return newhunter
	} else if hunter.oldmove == "down" {
		hunter.game.field[hunter.coordinate[0]][hunter.coordinate[1]] = "$"
	}
	return nil
}
