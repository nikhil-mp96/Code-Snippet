package main

import (
	"fmt"
	"strconv"
)

/*
	For loop is only looping statement in golang

	Different variations of for loop in golang are explained below

	Program Execution output:
		#Traditional For Loop# 1
		#Traditional For Loop# 2
		#Traditional For Loop# 3
		#Traditional For Loop# 4
		#Traditional For Loop# 5
		#Traditional For Loop# 6
		#Traditional For Loop# 7
		#Traditional For Loop# 8
		#Traditional For Loop# 9
		#Traditional For Loop# 10
		#Foreach# index - 0 :: Value - 1
		#Foreach# index - 1 :: Value - 2
		#Foreach# index - 2 :: Value - 3
		#Foreach# index - 3 :: Value - 4
		#Foreach# index - 4 :: Value - 5
		#Foreach# index - 5 :: Value - 6
		#Foreach# index - 6 :: Value - 7
		#Foreach# index - 7 :: Value - 8
		#Foreach# index - 8 :: Value - 9
		#Foreach# index - 9 :: Value - 10
		#While# Loop 1
		#While# Loop 2
		#While# Loop 3
		#While# Loop 4
		#While# Loop 5
		#While# Loop 6
		#While# Loop 7
		#While# Loop 8
		#While# Loop 9
		#While# Loop 10
		#Infinite while loop# 1
		#Infinite while loop# 2
		#Infinite while loop# 3
		#Infinite while loop# 4
		#Infinite while loop# 5
		#Infinite while loop# 6
		#Infinite while loop# 7
		#Infinite while loop# 8
		#Infinite while loop# 9
		#Infinite while loop# 10
*/

func main() {
	/*
		Traditional form of For Loops
		for initialization; condition; incrementer {
			...
		}
	*/
	// Creating for loop to print 1-10
	for i := 1; i <= 10; i++ {
		fmt.Println("#Traditional For Loop# " + strconv.Itoa(i))
	}

	/*
		Similar to foreach loops when looping over arrays
		for key, value := range array {
			...
		}

		range is keyword used to iterate over elements array using for each
		in each iteration it produces a pair of key value, typically used with slices
	*/
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range arr {
		fmt.Println("#Foreach# index - " + strconv.Itoa(index) + " :: Value - " + strconv.Itoa(value))
	}

	/*
		Traditional While loop using for in golang
		for condition {
			...
		}
	*/
	i := 1
	for i <= 10 {
		fmt.Println("#While Loop# " + strconv.Itoa(i))
		i++
	}

	/*
		Traditional infinite while loop using for loops in golang
		for {
			...
		}
		can be interpreted as a for loop without any condition is always true
	*/
	i = 1
	for {
		fmt.Println("#Infinite while loop# " + strconv.Itoa(i))
		i++
		if i > 10 {
			break
		}
	}
}
