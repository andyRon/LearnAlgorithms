package main

import "fmt"

func algorithm_A(n int) {
	fmt.Println(0)
}

func algorithm_B(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(0)
	}
}

func algorithm_C(n int) {
	for i := 0; i < 1000000; i++ {
		fmt.Println(0)
	}
}
