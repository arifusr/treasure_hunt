package main

func main() {
	m := make(chan string)
	q := make(chan string)
	game := NewGameTreasureHunter(m, q)

	game.Start()

	for {
		select {
		case msg := <-m:
			game.Process(msg)
		case <-q:
			return
		}
	}

}
