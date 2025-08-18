package main

import "fmt"

func Selectionsort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minindex := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minindex] {
				minindex = j
			}
		}
		arr[i], arr[minindex] = arr[minindex], arr[i]
	}
}
func main() {
	arr := []int{80, 90, 60, 70, 50, 30, 40, 10}


	Selectionsort(arr)
	fmt.Println("Selection sort :",arr)
}
