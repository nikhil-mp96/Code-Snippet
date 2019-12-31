package main

import "fmt"

func main() {
	defer fmt.Println("First Defer Statement in main()")

	fmt.Println("Program started execution on parent thread")
	defer fmt.Println("Second Defer Statement in main()")

	subFunctionWithoutPanic()
	defer fmt.Println("Third Defer Statement in main()")

	subFunctionWithPanic()
	defer fmt.Println("Third Defer Statement in main()")
}

func subFunctionWithoutPanic() {
	defer fmt.Println("First Defer Statement in subFunction() Executed")
	defer fmt.Println("Second Defer Statement in subFunction() Executed")
	defer fmt.Println("Third Defer Statement in subFunction() Executed")
}

func subFunctionWithPanic() {
	defer fmt.Println("First Defer Statement in subFunction() Executed")
	defer fmt.Println("Second Defer Statement in subFunction() Executed")
	panic("Panic !!!")
	defer fmt.Println("Third Defer Statement in subFunction() Executed")
}


// Output
/*
Program started execution on parent thread
Third Defer Statement in subFunction() Executed
Second Defer Statement in subFunction() Executed
First Defer Statement in subFunction() Executed
Second Defer Statement in subFunction() Executed
First Defer Statement in subFunction() Executed
Third Defer Statement in main()
Second Defer Statement in main()
First Defer Statement in main()
panic: Panic !!!

goroutine 1 [running]:
main.subFunctionWithPanic()
	/Users/nikhil/Let-Us-Go/Error Handling using Defer/understanding_defer_control_flow.go:27 +0x140
main.main()
	/Users/nikhil/Let-Us-Go/Error Handling using Defer/understanding_defer_control_flow.go:14 +0x22c
exit status 2
 */