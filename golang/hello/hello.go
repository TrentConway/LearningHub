// First golang project
// BASIC GO OPERATIONS 

<<<<<<< HEAD
// from tutorial: https://www.youtube.com/watch?v=C8LgvuEBraI
=======
>>>>>>> 7c6cc76bd7601ba70b991ee5da5ba6e09014f4e1

// executable commands must always include main package
package main

// format library -- control i/o
import (
	"fmt"
	"errors"
	"math"
)

func hello() {
	fmt.Println("Hello, world.")
}
func sum() {
	var x int = 7
	// refers the type 
	y := 5
	sum := x + y
	fmt.Println(sum)

	if sum < 10 {
		fmt.Println("smaller than 10")
	}else {
		fmt.Println("larger than 10")
	}
}

func list() {
	// zero indexed and initalised with zero value
	// array
	var a[5]int 
	a[2] = 3
	fmt.Println(a)
	// slice
	b := []int{5,4,3,2,1}
	// this does not modify a slice it returns a new one
	b = append(b,13)
	fmt.Println(b)

}

func dict () {
	v := make(map[string]int)
	v["triangle"] = 3
	v["square"] = 4
	v["pentagon"] = 5
	fmt.Println(v)

	delete(v, "square")
	fmt.Println(v)
}

func loop() {
	// for loop
	for i := 0; i<5; i++ {
		fmt.Println(i)
	}
	// iterate over an array or slice
	arr := []string{"a","b","c"}
	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}
}

// func can return 2 values 
func sqrt(x float64) (float64, error){
	if x < 0 {
		return 0, errors.New("undefined")
	}
	return math.Sqrt(x), nil
}

// struct - define custom types
type person struct {
	name string
	age int
}
<<<<<<< HEAD
=======

>>>>>>> 7c6cc76bd7601ba70b991ee5da5ba6e09014f4e1
func personFunc() {
	p := person{name:"trent", age: 24}
	fmt.Println(p.age)
}

// pointers
// normal functions perform a cp then modify
// you can use pointers to modify the value directly at the address
func point () {
	i := 7
	// pass in the pointer
	inc(&i)
	fmt.Println(i)

}
<<<<<<< HEAD
=======

>>>>>>> 7c6cc76bd7601ba70b991ee5da5ba6e09014f4e1
// pass in address
func inc(x *int) {
	// dereference address
	*x++
}

<<<<<<< HEAD

=======
>>>>>>> 7c6cc76bd7601ba70b991ee5da5ba6e09014f4e1
func main() {
	hello()
	sum()
	list()
	dict()
	loop()
	result, error := sqrt(16)
	fmt.Println(result, error)
	if error != nil {
		fmt.Println(error)
	} else{ 
		fmt.Println(result)
	}
	personFunc()
	point()
}
