package main

import "fmt"


func main() {

	var x[4] float64
	x[0]=20
	x[1]=203
	x[2]=202
	x[3]=206
	fmt.Println(x)
    var total float64 = 0
	// for i:=0; i < len(x); i++ {
    //    total  += x[i] 
	// }
	// fmt.Println(total/ float64(len(x)))


	for _, value := range x {
	  total += value
	}
	fmt.Println(total / float64(len(x)))
}