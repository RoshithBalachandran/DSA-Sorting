package main

import (
	"fmt"
	"time"
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

func InsertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func Seletction(arr []int) {

	n:=len(arr)
	for i:=0;i<n-1;i++{
		minindex:=i
		for j:=i+1;j<n;j++{
			if arr[j]<arr[minindex]{
				minindex=j
			}
		}
		arr[i],arr[minindex]=arr[minindex],arr[i]
	}
}

func main() {
	arr := []int{80, 50, 90, 70, 30, 40, 10}
	new := []int{5, 65, 95, 35, 45, 25}
	sel := []int{1, 6, 8, 9, 7, 4, 5, 3, 2}

	start := time.Now()
	BoubleSort(arr)
	InsertionSort(new)

	Seletction(sel)

	end := time.Since(start)

	fmt.Println("Bouble Sort :", arr)
	fmt.Println("Insertion Sort :", new)
	fmt.Println("Insertion Sort :", sel)

	fmt.Println("Time taken :", end)
}
