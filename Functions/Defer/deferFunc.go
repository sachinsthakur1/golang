package main

import "fmt"

func cleanup() {
	fmt.Println(("I am cleaning up "))
}

func main() {
	fmt.Println(("I am entering main function"))
	defer cleanup()
	fmt.Println(("I am exiting main function"))
}
