package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// main is the entry point for our program.
func main() {
	// --- Step 1: Generate a random secret number ---

	// The default random number generator is deterministic, meaning it will
	// produce the same sequence of numbers every time. To make it feel truly
	// random, we need to give it a unique starting point (a "seed").
	// Using the current time in nanoseconds is a great way to ensure the
	// seed is different on each run.
	rand.Seed(time.Now().UnixNano())

	// rand.Intn(100) generates a random integer between 0 and 99.
	// We add 1 to get a number in the range [1, 100].
	secretNumber := rand.Intn(100) + 1

	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it? You have 10 tries.")

	// --- Step 2: Set up a loop for 10 guesses ---

	// We use a classic 'for' loop that will run exactly 10 times.
	// The 'guesses' variable will keep track of the current attempt number.
	for guesses := 1; guesses <= 10; guesses++ {
		// Provide feedback to the player on their progress.
		fmt.Printf("Guess #%d (You have %d guesses left): ", guesses, 11-guesses)

		// --- Step 3: Read the player's input from the console ---

		// bufio.NewReader(os.Stdin) creates a new reader that is linked to
		// the standard input (the keyboard). This is more efficient for
		// reading text than reading one byte at a time.
		reader := bufio.NewReader(os.Stdin)

		// ReadString reads input until it encounters the specified delimiter ('\n' for the Enter key).
		// It returns the input string and any error that occurred.
		input, err := reader.ReadString('\n')
		if err != nil {
			// If an error occurs (e.g., end-of-file), we can't proceed.
			// log.Fatal will print the error message and exit the program.
			log.Fatal(err)
		}

		// The input string includes the newline character ('\n').
		// strings.TrimSpace removes any leading/trailing whitespace, including the newline.
		input = strings.TrimSpace(input)

		// The input is a string, but we need an integer to compare it.
		// strconv.Atoi (ASCII to Integer) converts the string to an integer.
		// This can fail if the user types non-numeric text.
		guess, err := strconv.Atoi(input)
		if err != nil {
			// If the conversion fails, inform the user and skip to the next loop iteration.
			// 'continue' is perfect for this.
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// --- Step 4: Compare the guess to the secret number ---

		// Here's the core game logic, using an if-else if-else chain.
		if guess < secretNumber {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > secretNumber {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			// The guess is correct!
			fmt.Println("Good job! You guessed it!")
			// 'return' exits the main function, immediately ending the program.
			// This is a clean way to stop since the game is won.
			return
		}
	}

	// --- Step 5: Handle the case where the player runs out of guesses ---

	// If the 'for' loop completes without the player guessing the number,
	// this line of code will be executed, revealing the secret number.
	fmt.Println("Sorry, you didn't guess my number. It was:", secretNumber)
}
