package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type myLogWritter struct{}

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Long hand option to print the response
	// bs := make([]byte, 99999)
	// response.Body.Read(bs)
	// fmt.Println(string(bs))

	// Short option to print the response
	// io.Copy(os.Stdout, response.Body)

	mlw := myLogWritter{}
	io.Copy(mlw, response.Body)
}

// Custom implementation of Write interface
func (myLogWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote tjhis many bytes:", len(bs))
	return len(bs), nil
}
