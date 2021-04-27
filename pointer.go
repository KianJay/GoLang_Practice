package main

import "fmt"

func changeVal(str *string) {
	*str = "pointer worked"
}

func changeVal2(str string) {
	str = "pointer worked"
}

func main() {
	toChange := "Hello"
	fmt.Println(toChange)
	changeVal(&toChange)
	fmt.Println(toChange)
}
