package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":    "#f44336",
		"pink":   "#e91e63",
		"purple": "#9c27b0",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// Hex code for red is #f44336
// Hex code for pink is #e91e63
// Hex code for purple is #9c27b0
