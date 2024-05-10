package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

type player struct {
	name          string
	missingChance uint8
	score         int
}

func newPlayer(n string) *player {
	return &player{
		name:          n,
		missingChance: uint8(rand.Intn(23) + 10),
	}
}

func main() {
	didThePoint := make(chan bool)
	wg.Add(2)
	setPoint := 11

	firstPlayer := newPlayer("John")
	fmt.Printf("Player %s has %d%% missing chance\n",
		firstPlayer.name, firstPlayer.missingChance)
	go game(firstPlayer, setPoint, didThePoint)

	secondPlayer := newPlayer("Tom")
	fmt.Printf("Player %s has %d%% missing chance\n",
		secondPlayer.name, secondPlayer.missingChance)
	go game(secondPlayer, setPoint, didThePoint)

	didThePoint <- false
	wg.Wait()

	fmt.Printf("Player %s Final Score: %d\n", firstPlayer.name, firstPlayer.score)
	fmt.Printf("Player %s Final Score: %d\n", secondPlayer.name, secondPlayer.score)
}

func game(p *player, setPoint int, didThePoint chan bool) {
	defer wg.Done()

	for {
		scored, ok := <-didThePoint

		if !ok {
			fmt.Println("Game ended")
			return
		}

		if scored {
			p.score++
			fmt.Printf("Player %s Scored: %d\n", p.name, p.score)

			if p.score == setPoint {
				fmt.Printf("Player %s Won.\n", p.name)
				close(didThePoint)
				return
			}
		}

		if rand.Intn(100) <= int(p.missingChance) {
			fmt.Printf("Player %s Missed\n", p.name)
			didThePoint <- true
		} else {
			fmt.Printf("Player %s Hit\n", p.name)
			didThePoint <- false
		}
	}
}
