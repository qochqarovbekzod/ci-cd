package main

import "fmt"

func main() {
	Hello("salim")
}

func Hello(name string) string {
	return fmt.Sprintf("hello %s", name)
}
