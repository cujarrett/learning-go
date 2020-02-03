# Work in Progress (WIP)

## What is it?
It's a small set of hacks to learn [Go](https://golang.org/). I'm using `Golang go1.13.6 darwin/amd64` for this adventure.

Go short for Golang is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. Go was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and large codebases.

Go is common in many [Cloud Native](https://www.cncf.io/) solutions such as Kubernetes, Terraform, Prometheus, and many more. It was the tenth most popular language in 2019 on GitHub

## Take Aways
### The Basics
main.go
```
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

#### Package Types
- `Executable` - Generates a file that we can run
- `Reusable` - Code used as "helpers", good place to put reusable logic

You must use `main` as the package name in order to make a `Executable` package type. Any other name besides `main` will cause the package type to be `Reusable`.

Anytime you make a `Executable` package, you must include a function inside of it called `main`.

### Import
In Go you use import statements to give your package access to code written in another package.

Example:
```
import "fmt"
```

### Variables
You can use the longer or shorthand manner of declaring variables. They achieve the same outcome. The latter relies on the Go compiler to determine the type.
```
	var card string = "Ace of Spaces"
```
```
	card := "Ace of Spades"
```

After the declaration the values are updated as such:
```
func main() {
	var card string = "Ace of Spaces"
	card = "Five of Diamonds"
	fmt.Println(card) // Five of Diamonds
}
```
```
func main() {
	card := "Ace of Spades"
	card = "Five of Diamonds"
	fmt.Println(card) // Five of Diamonds
}
```
### Functions
Use `func` to declare functions in Go.

### How is the main.go file organized?

## Things I like
- Types in Go feel simple :+1:

## Things I'm not a huge fan of
- It's really (really) opinionated. I was five minutes into Go and it was replacing my space indented code with tabs becuase it's [the Go "way"](https://github.com/golang/go/issues/16256#issuecomment-230173434). [Really?](https://media.giphy.com/media/CggoHW4h87Ktq/giphy.gif) [This Chrome plugin](https://github.com/sindresorhus/tab-size-on-github) by the amazing [@sindresorhus](https://github.com/sindresorhus) helps my happiness with readable code.
