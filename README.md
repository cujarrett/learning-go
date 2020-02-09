# Work in Progress (WIP)

## What is it?
It's a small set of hacks to learn [Go](https://golang.org/). I'm using `Golang go1.13.6 darwin/amd64` for this adventure.

Go short for Golang is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. Go was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and large codebases.

Go is common in many [Cloud Native](https://www.cncf.io/) solutions such as Kubernetes, Terraform, Prometheus, and many more. It was the tenth most popular language in 2019 on GitHub

## Take Aways
### The Basics
main.go
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world.")
}
```

### [Go CLI](https://golang.org/cmd/go/)
Some common uses:

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
Use `func` to declare functions in Go.

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

### Receiver Functions vs Passing in arguments to a Function
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

### Examples
#### Function with a receiver, parameters, and a return
'describe' is a function with a receiver of type 'color' that requires an argument of type 'string', then returns a value of type 'string'
```go
...
func (c color) describe(description string) (string) {
   return string(c) + " " + description
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

### Loops
```go
...
cards := []string{"Ace of Spades", newCard()}

for index, card := range cards {
	fmt.Println(index, card)
}
```

### Testing in Go
Go testing is not RSpec, mocha, jasmine, selenium, etc. To make a test, create a new file ending in `_test.go`. To run all tests in a package run the `go test` command.



### Type Conversion
```go
[]byte("Hello world.")
```
- `[]byte` - The type we want
- `("Hello world.")` - The type we have

## Things I like
- Types in Go feel simple :+1:
- Short hand variables option
- Go's compliler checking import references on save
- The [VS Code go plugin](https://github.com/microsoft/vscode-go) written by Microsoft is pretty great

## Things I'm not a huge fan of
- It's really (really) opinionated
- Mandated formatting with tabs becuase it's [the Go "way"](https://github.com/golang/go/issues/16256#issuecomment-230173434). [Go reasons these benifits for this choice](https://golang.org/doc/effective_go.html#formatting). [This Chrome plugin](https://github.com/sindresorhus/tab-size-on-github) by the amazing [@sindresorhus](https://github.com/sindresorhus) helps my happiness with readable code.
- Go's convention of using single letter variables feels bad
- Go's test framework doesn't provide any info about how many tests ran or how many tests passed or failed