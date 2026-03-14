![logo](./media/readme-logo.jpg)

<p align="center" class="toc">
   <strong><a href="#what-is-it">What is it?</a></strong>
   |
  <strong><a href="#things-i-like">Things I like</a></strong>
   |
   <strong><a href="#things-im-not-a-huge-fan-of">Things I'm not a huge fan of</a></strong>
   |
   <strong><a href="#learning-problems-completed">Learning Problems Completed</a></strong>
   |
   <strong><a href="#take-aways">Take Aways</a></strong>
</p>

## What is it?
It's a small set of hacks to learn [Go](https://golang.org/).

Go short for Golang is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. Go was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and large codebases.

Go is common in many [Cloud Native](https://www.cncf.io/) solutions such as Kubernetes, Terraform, Prometheus, and many more. It was the tenth most popular language in 2019 on GitHub

## Go 101 - [Go: The Complete Developer's Guide](https://www.udemy.com/course/go-the-complete-developers-guide/)
- [cards](./cards) - A simple app that builds a deck of playing cards with functions for deal, shuffle, print, etc
- [even-and-odds](./even-and-odds/main.go) - A simple app to parse a collection of numbers and print if they are even or odd
- [structs](./structs/main.go) - A simple app to show the use of structs, pointers, and memory addresses
- [maps](./maps/main.go) - A simple app showing of the basics of Go maps
- [interfaces](./interfaces/main.go) - A simple app showing of basics of Go interfaces
- [http](./http/main.go) - A simple app to hit google.com and write out the contents of the html to the console
- [shapes](./shapes/main.go) - A simple app to demonstrate multiple shapes implementing area differently and the interface they respect
- [print-file](./print-file/) - A simple app to read a file specified from the command line and print out the contents on the console
- [channels] - A simple app to hit popular websites and report if they are up with concurrency via [Go Routines](#go-routines) and [Go Channels](#channels)

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
The purpose of a package in Go is to group together code with a similar purpose. The Go packages are documented [here](https://golang.org/pkg/)

#### Package Types
- `Executable` - Generates a file that we can run
- `Reusable` - Code used as "helpers", good place to put reusable logic

You must use `main` as the package name in order to make a `Executable` package type. Any other name besides `main` will cause the package type to be `Reusable`. Anytime you make a `Executable` package, you must include a function inside of it called `main`.

### Pass by Value Language
Strictly speaking, there is only one way to pass parameters in Go - by value. Every time a variable is passed as parameter, a new copy of the variable is created and passed to called function or method. The copy is allocated at a different memory address. The [structs](./structs/main.go) code shows a good example use.

Whenever you pass an integer, float, string, or strut into a function it creates a copy of each argument, and these copies are used inside of the function.

Some syntax associated with this:
- `&foo` - Gives the memory address of the foo value this variable is pointing at
- `*foo` - Gives the value of this memory address is  pointing at

## Go Routines
A Go routine is a function that is capable of running concurrently with other functions. To create a Go routine we use the keyword `go` followed by a function invocation. The Go Scheduler handles executing code for a CPU core. The initial running code is referred to as the Main routine and the others (when needed) are referred to as Child Routines.

## Channels
Channels provide a way for two Go routines to communicate with one another and synchronize their execution.

## Gotcha's in Go

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
