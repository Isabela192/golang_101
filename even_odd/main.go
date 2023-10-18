package main

import "fmt"

func main() {
	mySlice := make([]int, 11)
	for i := range mySlice {
		if i%2 == 0 {
			fmt.Println(i, "is Even")
		}
		if i%2 != 0 {
			fmt.Println(i, "is Odd")
		}
	}
}
