package main

import "fmt"

func main() {
	for i:= 1; i<= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i,"Even")
		}else{
			fmt.Println(i,"Odd")
		}
		
	}

   for i:=1;i<=100;i++ {
	if i % 3 == 0 {
		fmt.Println(i)
	}
   }

   for i:=1;i<=100;i++ {
	 if i % 5 == 0 && i % 3 == 0 {
		fmt.Println(i ,"FizzBuzz")
	}else if i % 3 == 0 {
		fmt.Println(i, "Fizz")
	}else if i % 5 == 0 {
		fmt.Println(i, "Buzz")
	}
   }
}