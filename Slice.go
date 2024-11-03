package main

import "fmt"


func main() {

	// arr := [5]float64{1,2,3,4,5}
	// x := arr[2:3]
   
	// fmt.Println(x)
	// slice1 := []int{1,2,3}
	// slice2 := append(slice1, 4, 5)
	// fmt.Println(slice1, slice2)
	// fmt.Println(slice2)

	slice3 := []int{1,2,3,5}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}