package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// Task #1. Rock-paper-scissors app. You play with the app.

const (
	ROCK     = iota // 00
	PAPER           // 01
	SCISSORS        // 10
)

func main() {
	fmt.Println("Let's play a round of rock paper scissors!")
	fmt.Printf("Input your choice \n[0] Rock\n[1] Paper\n[2] Scissors: ")

	in := readUserInput()
	rnd := rand.Intn(2)

	computeWinner(in, rnd)
}

func readUserInput() int {
	reader := bufio.NewReader(os.Stdin)
	rawIn, err := reader.ReadBytes(byte('\n'))
	if err != nil {
		log.Fatalf("Failed to read user input: %v", err)
	}

	// We validate the user input based on the ascii code
	if rawIn[0] < 48 || rawIn[0] > 50 {
		fmt.Print("Input value must be 0, 1 or 2. Please try again: ")
		return readUserInput()
	}

	// Convert byte to int by subtracting ASCII value of '0'
	return int(rawIn[0] - '0')
}

func computeWinner(player int, cpu int) {
	fmt.Printf("Your input: \t\t%d\n", player)
	fmt.Printf("Computer input: \t%d\n", cpu)

	switch {
	case player == cpu:
		fmt.Println("Tie! Try again next round.")
	case (player == ROCK && cpu == SCISSORS) || (player == SCISSORS && cpu == PAPER) || (player == PAPER && cpu == ROCK):
		fmt.Println("Nice! You won.")
	default:
		fmt.Println("The computer wins! Try again next round.")
	}
}
