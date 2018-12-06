package main

import (
	"fmt"
)

func main() {
	s := "test"
	for _, c := range s {
		fmt.Println(string(c))
	}
}
