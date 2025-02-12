package main

import "fmt"

type learner struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Grade   int     `json:"class"`
	Average float64 `json:"average"`
}

func main() {
	fmt.Println("vim-go")
}
