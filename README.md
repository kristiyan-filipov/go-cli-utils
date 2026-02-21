# Go CLI Utils

A simple Go package providing reusable utilities for building interactive CLI applications. Includes input prompts, menu handling, confirmation dialogs, and basic terminal utilities.

## Installation

```bash
go get github.com/kristiyan-filipov/go-cli-utils
```

## Usage

```go
package main

import (
	"fmt"

	cliutils "github.com/kristiyan-filipov/go-cli-utils"
)

func main() {
	if cliutils.Confirm("Do you want to continue? (y/n): ") {
		fmt.Println("Continuing...")
	} else {
		fmt.Println("Cancelled.")
	}
}
```

## Functions

### `Clear()`

Clears the terminal screen. Works on Unix and Windows systems.

### `CapitalizeWord(str string) string`

Capitalizes the first letter of a string.

### `InputPrompt(prompt string) (string, error)`

Displays a prompt and returns user input as a string.

### `Input() (string, error)`

Reads input from the user without a prompt.

### `ChooseNumberFromRange(numberRange int, prompt string) int`

Prompts the user to choose a number within a given range. Loops until valid input is entered.

### `Option` struct

```go
type Option struct {
	Name           string
	OptionFunction func()
}
```

Represents an option with a name and an associated function to execute, most common use is menu navigation.

### `ChooseOption(options []Option, prompt string) int`

Displays a numbered list of `Option`s and prompts the user to select one. Returns the index of the selected option.

### `ChooseAndRunOption(options []Option, prompt string, clear bool)`

Displays options, optionally clears the screen, and executes the selected option's function.

### `Confirm(prompt string) bool`

Prompts the user with a yes/no question. Loops until a valid response is entered and returns `true` for yes, `false` for no.

## Example

```go
package main

import (
	"fmt"

	cliutils "github.com/kristiyan-filipov/go-cli-utils"
)

func main() {
	cliutils.Clear()

	options := []cliutils.Option{
		{
			Name: "Say Hello",
			OptionFunction: func() {
				fmt.Println("Hello!")
			},
		},
		{
			Name: "Exit",
			OptionFunction: func() {
				fmt.Println("Goodbye!")
			},
		},
	}

	cliutils.ChooseAndRunOption(options, "Choose an action: ", true)
}
```
