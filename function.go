package main

import "fmt"


func main() {

	xs := []float64{98,93,77,82,83}
    x:= 5
	f(x)
	fmt.Println(f1())
	z,y := fn()
	fmt.Println(z , y)
   fmt.Println("The average is ", average(xs))
//    Variadic Functions
   fmt.Println(add(10,20,30))
   //find the greatest number
   fmt.Println("The greatest values is ",findGreatest(10,20,30,3,5,21))
   fmt.Println(multply(10,20,30))
   xsa := []int{1,2,3}
   fmt.Println(add(xsa...))


//    7.4 Closure

   add := func(x, y int) int {
	return x + y
	}
	fmt.Println(add(12,12))

	// 7.5 Recursion

	fmt.Println(factorial(7))

	fmt.Println(Fibonacci(10))



	
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4


	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) // 0
	fmt.Println(nextOdd()) // 2
	fmt.Println(nextOdd()) // 4
	// 7.6 Defer, Panic & Recover

	defer second()
	first()

	// Panic & Recover
	defer func() {
		str := recover()
		fmt.Println(str)
		}()
	panic("PANIC")

	
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
	total += v
	}
	return total / float64(len(xs))
}

func f(x int) {
	fmt.Println(x)
}

func f1() int {
	return f2()
}	
func f2() int {
    return 1
}
func fn() (int, int) {
	return 2 , 3
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
	total += v
	}
	return total
}

func multply(args ...int) int {
	total := 1
	for _, v := range args {
	total *= v
	}
	return total
}

func findGreatest(args ...int) int {
   greatest := args[0]
   for _, v:= range args {
	if greatest < v{
		greatest = v
	}
   }
   return greatest
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
	ret = i
	i += 2
	return
	}
}


func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
	ret = i
	i += 2
	return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	} 
	return x * factorial(x-1)
}

func Fibonacci(x uint) uint {
	if x == 0 {
		return 0
	} else if x == 1{
		return 1
	}
	return Fibonacci(x-1) + Fibonacci(x-2)
}

func first() {
	fmt.Println("1st")
}
func second() {
  fmt.Println("2nd")
}