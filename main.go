package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Print a welcome message
	intro()

	// create a channel to indicate when the user wants to quite
	doneChan := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(os.Stdin, doneChan)

	// block until the doneChan gets a value
	<-doneChan

	// close te channel
	close(doneChan)

	// say goodbye
	fmt.Println("goodbye")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check to see if the user wants to quite
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert what the user typed into an int
	numbToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "please enter a whole number", false
	}

	_, msg := isPrime(numbToCheck)

	return msg, false
}

func intro() {
	fmt.Println("is it prime?")
	fmt.Println("------------")
	fmt.Println("enter a whole number, and we'll tell you it is a prime number or not. Enter q to quit")

	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "negative numbers are not prime by definition"
	}

	//------------------------------
	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is devisible by %d", n, i)
		}

	}

	return true, fmt.Sprintf("%d is a prime numnber!", n)
}
