package main

import (
	"fmt"
	"github.com/mgll/golab/concurrency/runner"
	"log"
	"math"
	"os"
	"time"
)

const timeout = 4 * time.Second

func main() {
	log.Println("Starting work.")
	r := runner.New(timeout)

	r.Add(primeGenerator(10), wait(1), wait(2), wait(1), wait(2))

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to runner timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

func wait(n uint) func() {
	return func() {
		log.Printf("Waiting for %d second\n", n)
		time.Sleep(time.Duration(n) * time.Second)
	}
}

func primeGenerator(n uint) func() {
	return func() {
		fmt.Printf("Start searching %v primes...\n", n)
		var primes []int
		i := 2
		for {
			if uint(len(primes)) == n {
				break
			}
			if isPrime(i) {
				primes = append(primes, i)
			}
			i++
		}
		fmt.Printf("Primes: %v\n", primes)
	}
}

func isPrime(n int) bool {
	sqRoot := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqRoot; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
