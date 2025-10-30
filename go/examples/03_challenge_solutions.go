package main

import (
	"fmt"
	"unicode"
)

// main is the entry point that calls the solution functions.
func main() {
	fmt.Println("--- Challenge 1: FizzBuzz ---")
	fizzBuzz()

	fmt.Println("\n--- Challenge 2: Password Checker ---")
	passwordChecker("Password123") // A valid password
	passwordChecker("short")      // Invalid: too short
	passwordChecker("longpassword") // Invalid: no uppercase or number
	passwordChecker("LongPassword") // Invalid: no number
	passwordChecker("password123")  // Invalid: no uppercase

	fmt.Println("\n--- Challenge 3: Sum of Numbers ---")
	sumOfNumbers(100)
	sumOfNumbers(10)
}

// --- Challenge 1: FizzBuzz ---
// Goal: Write a program that prints the numbers from 1 to 100.
// - For multiples of three, print "Fizz".
// - For multiples of five, print "Buzz".
// - For multiples of both, print "FizzBuzz".

func fizzBuzz() {
	// A classic 'for' loop is perfect for iterating from 1 to 100.
	for i := 1; i <= 100; i++ {
		// The modulo operator (%) gives the remainder of a division.
		// If 'i % 3 == 0', it means 'i' is perfectly divisible by 3.

		// We must check for the "FizzBuzz" case first (multiples of both 3 and 5),
		// because if we checked for "Fizz" first, it would print "Fizz" for a number
		// like 15 and move on, never reaching the "FizzBuzz" check.
		// A number is a multiple of both 3 and 5 if it's a multiple of 15.
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			// If none of the above conditions are met, just print the number.
			fmt.Println(i)
		}
	}
}

// --- Challenge 2: Password Checker ---
// Goal: Check if a password meets the following criteria:
// - At least 8 characters long.
// - Contains at least one uppercase letter.
// - Contains at least one number.

func passwordChecker(password string) {
	fmt.Printf("Checking password: '%s'\n", password)
	// We'll use boolean flags to keep track of whether each condition is met.
	var hasMinLength bool
	var hasUppercase bool
	var hasNumber bool

	// 1. Check for minimum length.
	if len(password) >= 8 {
		hasMinLength = true
	}

	// 2. Iterate over the string to check for uppercase letters and numbers.
	// A 'for...range' loop on a string iterates over its runes (characters).
	for _, char := range password {
		// The 'unicode' package has helpful functions for character classification.
		if unicode.IsUpper(char) {
			hasUppercase = true
		}
		if unicode.IsDigit(char) {
			hasNumber = true
		}
	}

	// 3. Provide feedback based on the flags.
	if hasMinLength && hasUppercase && hasNumber {
		fmt.Println("  -> Password is valid.")
		return // Exit the function early if everything is okay.
	}

	// If we're still here, the password is not valid. Let's give detailed feedback.
	fmt.Println("  -> Password is NOT valid. Issues found:")
	if !hasMinLength {
		fmt.Println("    - Must be at least 8 characters long.")
	}
	if !hasUppercase {
		fmt.Println("    - Must contain at least one uppercase letter.")
	}
	if !hasNumber {
		fmt.Println("    - Must contain at least one number.")
	}
}

// --- Challenge 3: Sum of Numbers ---
// Goal: Calculate and print the sum of all numbers from 1 to a given number N.

func sumOfNumbers(n int) {
	// Initialize a variable to hold the sum.
	sum := 0

	// Loop from 1 up to and including N.
	for i := 1; i <= n; i++ {
		// On each iteration, add the current number 'i' to the sum.
		sum += i // This is shorthand for 'sum = sum + i'
	}

	fmt.Printf("The sum of numbers from 1 to %d is: %d\n", n, sum)
}
