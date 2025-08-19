package main

import "fmt"

func Reverse(s string) string {

	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes)
}

func main() {
	s := "Roshith😊"
	fmt.Println("Reversed string :",Reverse(s))
}