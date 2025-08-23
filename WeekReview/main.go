package main

import (
	"fmt"
)

func BoubleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func ReverseString(s string)string{
	runes:=[]rune(s)
	n:=len(runes)
	for i:=0;i<n/2;i++{
		runes[i],runes[n-i-1]=runes[n-i-1],runes[i]
	}
	return string(runes)
}

func PalandromCheck(s string)bool{
	runes:=[]rune(s)
	n:=len(runes)
	for i:=0;i<n/2;i++{
		if runes[i]!=runes[n-i-1]{
			return false
		}
	}
	return true
}

func VovelsCheck(s string)[]string{
	Vovels:=[]rune("aeiouAEIOU")
	runes:=[]rune(s)
	
	var found []string
	
	for i:=0;i<len(runes);i++{
		for j:=0;j<len(Vovels);j++{
			if runes[i]==Vovels[j]{
				found = append(found, string(runes[i]))
			}
		}
	}
	return found
}

func main() {
	arr := []int{50, 90, 60, 10, 40, 30}
	BoubleSort(arr)

	fmt.Println("Bouble Sort array :",arr)

	text:="Roshith"
	fmt.Println("Reversed string :",ReverseString(text))


	text2:="ABA"
	fmt.Println("Palandrom check :",PalandromCheck(text2))

	text3:="Roshith"
	fmt.Println("Palandrom check :",VovelsCheck(text3))
}