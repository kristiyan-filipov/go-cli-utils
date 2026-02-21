package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func Clear() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}

func CapitalizeWord(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(str[:1]) + str[1:]
}

// User input

var reader = bufio.NewReader(os.Stdin)

func InputPrompt(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func Input() (string, error) {
	return InputPrompt("")
}

func ChooseNumberFromRange(numberRange int, prompt string) int {
	for {
		input, err := InputPrompt(prompt)
		if err != nil {
			log.Println(err)
			continue
		}
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > numberRange {
			log.Printf("Invalid input. Input should be a number between 1 and %d.", numberRange)
			continue
		}
		return choice
	}
}

// Options

type Option struct {
	Name           string
	OptionFunction func()
}

func ChooseOption(options []Option, prompt string) int {
	for index, option := range options {
		fmt.Println(index+1, ") ", option.Name)
	}
	fmt.Println()
	return ChooseNumberFromRange(len(options), prompt)
}

func ChooseAndRunOption(options []Option, prompt string, clear bool) {
	index := ChooseOption(options, prompt)
	if clear {
		Clear()
	}
	options[index-1].OptionFunction()
}

func Confirm(prompt string) bool {
	for {
		input, err := InputPrompt(prompt)
		if err != nil {
			log.Println(err)
			continue
		}

		confirmation := strings.ToLower(strings.TrimSpace(input))

		switch confirmation {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			log.Println("Invalid input. Input should be 'y', 'yes', 'n' or 'no'.")
		}
	}
}
