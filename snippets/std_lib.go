package main

import "fmt"

func main() {
	elms := []string{"aap", "noot", "mies"}
	elms = append(elms[:1], elms[2:]...)
	fmt.Println(elms)
}
