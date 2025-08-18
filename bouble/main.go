package main

import (
	"fmt"
	"time"
)

func BpubleSort(arr []int) {

	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{90, 10, 80, 20, 50}
	start := time.Now()

	BpubleSort(arr)
	end:=time.Since(start)

	fmt.Println("Bouble sort :",arr)

	fmt.Println("Time taken :",end)
}