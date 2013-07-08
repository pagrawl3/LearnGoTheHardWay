package main 

import "fmt"

type Position struct {
	x int
	y int
}

func main() {
	pos := Position{1,1}
	fmt.Printf("%+v\n",pos) //should print out {x:1 y:1}
}