package main

import (
	"log"
	"github.com/learnAlgorithm/util/helper"
)

func selectionSort(arr []int, n int) {

	for i := 0; i < n-1; i++ {
		//find the min index
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			helper.Swap(arr, i, minIndex)
		}

	}

}



func main() {
	arr := []int{9, 7, 10, 88, 72, 777, 21, 6, 5, 4, 3}
	selectionSort(arr, len(arr))
	log.Println(arr)
}
