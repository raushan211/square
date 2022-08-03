package main

import "fmt"

func main() {
	var x int
	fmt.Println("enter the number")
	fmt.Scanln(&x)
	square := x * x
	cube := x * x * x
	fmt.Println("Square = ", square)
	fmt.Println("Qube = ", cube)

}
