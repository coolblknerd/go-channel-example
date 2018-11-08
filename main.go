package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// This is creating the channel(unbuffered)
	court := make(chan int)

	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	// This starts the game
	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// Simulates waiting for the ball to be hit back to us
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// This dicatates if the ball is missed
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Closes the channel to signal we lost
			close(court)
			return
		}

		// Display and then increment the hit count by one
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hits ball back to opposing player
		court <- ball
	}
}
