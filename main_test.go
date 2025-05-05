package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_alpha_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime numnber!"},
		{"not prime", 8, false, "8 is not a prime number because it is devisible by 2"},
		{"zero", 0, false, "0 is not prime, by definition"},
		{"one", 1, false, "1 is not prime, by definition"},
		{"negative number", -7, false, "negative numbers are not prime by definition"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)

		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_alpha_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it before
	os.Stdout = oldOut

	// read the output of our prompt() from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it before
	os.Stdout = oldOut

	// read the output of our prompt() from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "enter a whole number") {
		t.Errorf("incorrect intro: expected -> but got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "7 is prime", input: "7", expected: "7 is a prime numnber!"},
		{name: "q to quit", input: "q", expected: ""},
		{name: "check converted", input: "hello world", expected: "please enter a whole number"},
		{name: "empty string", input: "", expected: "please enter a whole number"},
	}

	for _, tc := range testCases {
		input := strings.NewReader(tc.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, tc.expected) {
			t.Errorf("%s: expected %s, but got %s", tc.name, tc.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// need a channel, and an instance of io.Reader
	doneChan := make(chan bool)

	// create a reference to bytes.Buffer
	var stdin bytes.Buffer

	// stimulate user's pressing 1 then enter
	// pressing q then enter
	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
