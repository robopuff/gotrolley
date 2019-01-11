package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type output interface {
	booleanInput() bool
	print(message string)
	println(message string)
	printf(format string, a ...interface{})
	pullLever() bool
}

type terminal struct{}

func (t terminal) booleanInput() bool {
	reader := bufio.NewReader(os.Stdin)
	read, _ := reader.ReadString('\n')

	answers := map[string]bool{
		"y": true,
		"c": true,
		"t": true,
		"1": true,
		"n": false,
		"d": false,
		"f": false,
		"0": false,
		"2": false,
	}

	if read == "\n" {
		t.print("You HAVE to make a choice, what's its going to be? ")
		return t.booleanInput()
	}

	read = strings.ToLower(read)[0:1]
	res, ok := answers[read]
	if !ok {
		t.print("I didn't catch that...You HAVE to make a choice. ")
		return t.booleanInput()
	}

	return res
}

func (t terminal) println(message string) {
	t.printf("%s\n", message)
}

func (t terminal) printf(format string, a ...interface{}) {
	t.print(fmt.Sprintf(format, a...))
}

func (t terminal) print(message string) {
	bytes := []byte(message)
	for _, char := range bytes {
		fmt.Print(string(char))
		time.Sleep(20 * time.Millisecond)
	}
}

func (t terminal) pullLever() bool {
	t.print("Do you pull a lever? [Y/N] ")
	return t.booleanInput()
}

func newTerminal() *terminal {
	return &terminal{}
}
