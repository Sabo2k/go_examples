package main

import (
"fmt"
"errors"
)

// struct ---
type person struct {
	name string
	age int
}

// Funktionen ---
func multiply(x int, y int) {
	
	return x * y
}

// main fct. ---
func main() {
	
// basics ---
	const helloworld = "Hello World"
	x := 5 // implicit integer variable x
	y := 7
	sum := x + y
	
	if x > 6 {
		fmt.Println("x is bigger than 6")
	}
	
	fmt.Println(sum)
	
	p := person{name: "Jenny", age: 23}
	
// arrays ---
	a := [5]int{5, 4, 3, 2, 1}
	a = append(a, 10) // a = {5, 4, 3, 2, 1, 10}
	b := []int{6, 7, 8, 9}
	
	b[0] = 7
	
// maps ---
	mymap := make(map[string]int)
	mymap["Dreieck"] = 2
	mymap["Viereck"] = 3
	mymap["Kreis"] = 4
	
	delete(mymap, "Dreieck")
 	
// for loop ---
	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
	
	for {
		fmt.Println("Endlosschleife\n")
	}
	
// pointer ---
	num := 7
	increment(&num)
	
}

func main increment(x *int) {
	*x++
}
