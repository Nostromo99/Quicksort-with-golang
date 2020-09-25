package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	x := []int{82, 9, 90, 27, 37, 40, 24, 23, 96, 76, 88, 38, 28, 64, 57, 85, 7, 83, 95, 58, 39, 79, 6, 43, 92, 46, 20, 91, 51, 35, 68, 94, 52, 32, 93, 53, 98, 45, 84, 21, 14, 97, 74, 75, 2, 99, 49, 77, 60, 16}
	// x := []int{1, 5, 3, 4, 2}
	quicksort(x)
	fmt.Println(x)

}
func partition(list []int, left int, right int) int {
	pivot := left
	left++
	for left <= right {
		for left <= right && list[left] < list[pivot] {
			left++
		}
		for right >= left && list[right] > list[pivot] {
			right--
		}
		if left < right {
			list[left], list[right] = list[right], list[left]

			left++
			right--
		}
	}
	list[pivot], list[left-1] = list[left-1], list[pivot]
	return right
}
func quicksort(list []int) {
	//shuffling list to remove worst case complexity
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	quicksorter(list, 0, len(list)-1)
}
func quicksorter(list []int, start int, end int) {
	if start >= end || end <= 0 {
		return
	}
	pivot := partition(list, start, end)
	quicksorter(list, start, pivot-1)
	quicksorter(list, pivot+1, end)

}
