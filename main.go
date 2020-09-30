package main

import "fmt"

func main()  {
	var base uint64
	var height uint64
	fmt.Scanln(&base)
	fmt.Scanln(&height)
	area := (base*height) /2
	fmt.Println(area)
}