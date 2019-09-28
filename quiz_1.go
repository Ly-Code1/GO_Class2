package main

import (
	"fmt"
)



func main() {
	fruits := []string{"Apple","Orage","Plum","Pomegranate","Grape"}
	fmt.Println("o fruits : ", fruits)

	//Problem
	//some := fruits[1:3]

	//ex1 : solution
	some := make([]string,len(fruits))
	copy(some, fruits[1:3])


	//ex2 : solution
	//some :=fruits[1:3:3]


	some = append(some, "Banana")

	fmt.Println("f", len(fruits), cap(fruits))
	fmt.Println("s", len(some), cap(some))

	fmt.Println("n fruits: ", fruits)
	fmt.Println("some : ", some)
}
