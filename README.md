# Work in Progress (WIP)

## What is it?
It's a small set of hacks to learn [Go](https://golang.org/). I'm using `Golang go1.13.6 darwin/amd64` for this adventure.

Go short for Golang is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. Go was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and large codebases.

Go is common in many [Cloud Native](https://www.cncf.io/) solutions such as Kubernetes, Terraform, Prometheus, and many more. It was the tenth most popular language in 2019 on GitHub

## Learning Problems Completed
- [cards](./cards) - A simple app that builds a deck of playing cards with functions for deal, shuffle, print, etc.
- [even-and-odds](./even-and-odds) - A simple kata to parse a collection of numbers and print if they are even or odd.
- [structs] - A simple bit of code to show the use of structs, pointers, and memory addresses.

## Take Aways

### Go CLI
Go has a [Command Line Interface (CLI)](https://golang.org/cmd/go/). Some common uses of Go's CLI:

| CLI Command  | Use                                                        |
|--------------|------------------------------------------------------------|
| `go build`   | Compiles a bunch of go source code files                   |
| `go run`     | Compiles and executes one or two files                     |
| `go fmt`     | Formats all the code in each file in the current directory |
| `go install` | Compiles and installs a package                            |
| `go get`     | Downloads the raw source code of someone else's package    |
| `go test`    | Runs any tests associated with the current project         |

### Go Packages
The purpose of a package in Go is to group together code with a similar purpose.

The Go packages are documented [here](https://golang.org/pkg/)

#### Package Types
- `Executable` - Generates a file that we can run
- `Reusable` - Code used as "helpers", good place to put reusable logic

You must use `main` as the package name in order to make a `Executable` package type. Any other name besides `main` will cause the package type to be `Reusable`.

Anytime you make a `Executable` package, you must include a function inside of it called `main`.

### Import
In Go you use import statements to give your package access to code written in another package.

Example:
```go
import "fmt"
```

### Variables
You can use the longer or shorthand manner of declaring variables. They achieve the same outcome. The latter relies on the Go compiler to determine the type.
```go
...
	var card string = "Ace of Spaces"
```
```go
...
	card := "Ace of Spades"
```

After the declaration the values are updated as such:
```go
...
func main() {
	var card string = "Ace of Spaces"
	card = "Five of Diamonds"
	fmt.Println(card) // Five of Diamonds
}
```
```go
...
func main() {
	card := "Ace of Spades"
	card = "Five of Diamonds"
	fmt.Println(card) // Five of Diamonds
}
```

### Functions
Use `func` to declare functions in Go. Go does not support function overloading.

Functions that return must declare the return type.
```go
...
func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
```

#### Receiver Functions vs Passing in arguments to a Function
Receiver Functions execute on the current object in memory where as a normal function would take in the object via arguments. Receiver Functions is usually by reference and passing in arguments to a function is by value.

### Multiple return values
```go
...
func main() {
    color1, color2, color3 := colors()

    fmt.Println(color1, color2, color3)
}

func colors() (string, string, string) {
    return "red", "yellow", "blue"
}
```

### Arrays & Slices
- Array in Go is a fixed length list of things
- Slice in Go is an array that can grow or shrink
- Arrays and Slices in Go must define a singular type
- Go is zero based
- Slice range sytnax is startIndexIncluding to upToNotIncluding
  ```go
	fruits := []string{"apple", "banana", "grape", "orange"}
	fruits[0:2] // "apple" "banana"
	fruits[:2] // "apple" "banana"
	```

### Struct
A struct is a collection of fields. There are three ways of declaring struts.

- **Value Type**
- All keys must be the same type
- Values can be of different type
- Keys don't support indexing
- You need to know all the different fields at compile time
- Use to represent a "thing" with a lot of different properties

```go
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2}) // {1 2}
}
```

### Maps
A map maps keys to values. The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added.

- **Reference Type**
- All keys must be the same type
- All values must be the same type
- Keys are indexed, meaning we can iterate over them
- Use to represent a collection of related properties
- Don't need to know all the keys at compile time

```go
	colors := map[string]string{
		"red":    "#f44336",
		"pink":   "#e91e63",
		"purple": "#9c27b0",
	}

	fmt.Println(colors) // map[pink:#e91e63 purple:#9c27b0 red:#f44336]
```

or

```go
	moreColors := make(map[string]string)
	moreColors["blue"] = "#03a9f4"
	moreColors["cyan"] = "#00bcd4"
	moreColors["teal"] = "#009688"

	fmt.Println(moreColors) // map[blue:#03a9f4 cyan:#00bcd4 teal:#009688]
```

### Loops
```go
...
cards := []string{"Ace of Spades", newCard()}

for index, card := range cards {
	fmt.Println(index, card)
}
```

### Type Conversion
For example:
```go
[]byte("Hello world.")
```
- `[]byte` - The type we want
- `("Hello world.")` - The type we have

### Zero Values
Variables declared without an explicit initial value are given their zero value. The zero value is:

- `0` for numeric types
- `false` for the boolean type
- `""` (the empty string) for strings

### Pass by Value Language
Strictly speaking, there is only one way to pass parameters in Go - by value. Every time a variable is passed as parameter, a new copy of the variable is created and passed to called function or method. The copy is allocated at a different memory address. The [structs](./structs/main.go) code shows a good example use.

Whenever you pass an integer, float, string, or strut into a function it creates a copy of each argument, and these copies are used inside of the function.

Some syntax associated with this:
- `&foo` - Gives the memory address of the foo value this variable is pointing at
- `*foo` - Gives the value of this memory address is  pointing at

![pointers-and-address-info](https://user-images.githubusercontent.com/16245634/74107990-ed524680-4b3a-11ea-945c-e9154ce7bb04.png)

![pointer-use-in-receiver-vs-in-use](https://user-images.githubusercontent.com/16245634/74108076-f4c61f80-4b3b-11ea-9a7c-02f03fd8fa3e.png)

## Interfaces
Interfaces are named collections of method signatures.

Example:
```go
package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
```

#### Gotcha's in Go
![some-types-behave-different](https://user-images.githubusercontent.com/16245634/74108425-f0036a80-4b3f-11ea-95a1-5bd334beef49.png)

- Everything in Go is pass by value
- When we create a slice in Go it will automatically create an array and a structure that records the length of the slice, the capacity of the slice, and a reference to the underlying array
- With "value types" in Go, we have to worry about pointers if we want to pass a value to a function and modify the orginal value inside the function
- With "reference types" we do not need to worry about pointers if we want to pass a value to a function and modify the orginal valye inside the function
- A slice is a "reference type"

### Testing in Go
Go testing is not RSpec, mocha, jasmine, selenium, etc. To make a test, create a new file ending in `_test.go`. To run all tests in a package run the `go test` command.

Useful test command variants:
- `go test -coverprofile=coverage.out` - Get a report of run tests
- `go tool cover -html=coverage.out` - Get a nice HTML report of run tests
- `go test -v` - A more detailed test report of what tests pass and what failed

## Things I like
- Types in Go feel simple :+1:
- Short hand variables option
- Go's compliler checking import references on save
- The [VS Code go plugin](https://github.com/microsoft/vscode-go) written by Microsoft is pretty great

## Things I'm not a huge fan of
- It's really (really) opinionated
- Mandated formatting choices
  - Tabs for indents because it's [the Go "way"](https://github.com/golang/go/issues/16256#issuecomment-230173434). [Go reasons these benifits for this choice](https://golang.org/doc/effective_go.html#formatting). [This Chrome plugin](https://github.com/sindresorhus/tab-size-on-github) by the amazing [@sindresorhus](https://github.com/sindresorhus) helps my happiness with readable code.
  - Comma dangle
- Go's convention of using single letter variables feels bad
- Go's [struct](#structs) allowed construction of using ordered values feels prone to mistakes in the future when the order of the struct changes. I'm glad they offer a named attribute option.