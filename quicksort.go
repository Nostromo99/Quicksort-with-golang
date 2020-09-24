package main

import "fmt"

func main() {
	x := []int{9, 8, 7, 6, 5, 4}
	quicksorter(x, 0, len(x)-1)
	fmt.Println(x)

}
func partition(list []int, left int, right int) int {
	pivot := left
	temp := left
	left++
	for left <= right {
		for list[left] < list[pivot] && left < right {
			left++
		}
		for list[right] > list[pivot] && right > pivot {
			right--
		}
		if left <= right {
			list[left], list[right] = list[right], list[left]
			temp = left
			left++
			right--
		}
	}
	list[temp], list[pivot] = list[pivot], list[temp]
	return temp
}
func quicksort() {

}
func quicksorter(list []int, start int, end int) {
	if start >= end || end <= 0 {
		return
	}
	pivot := partition(list, start, end)
	quicksorter(list, start, pivot-1)
	quicksorter(list, pivot+1, end)

}
