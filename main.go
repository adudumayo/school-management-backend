package main

import "fmt"

type learner struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Average string `json:"average"`
}

func main() {
	fmt.Println("vim-go")
}
