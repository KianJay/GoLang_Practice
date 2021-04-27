package main

import "fmt"

/// find out the number that two times
func main() {
	var a []int = []int{1, 2, 4, 4, 15, 15, 7, 88, 88, 3}
	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}

	/*
			for i := 0; i < len(a); i++ {
				fmt.Println(a[i])
			}

		for _, element := range a {
			fmt.Printf("%d\n", element)
		}*/
}
