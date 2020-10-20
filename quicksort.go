package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	// x := []int{1, 5, 3, 4, 2}
	// x := []int{5, 1, 2, 3, 4, 6}
	// goquicksort(x)
	// fmt.Println(x)
	matrix := [][]int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		newlist := []int{}
		for j := 0; j <= 100000; j++ {
			newlist = append(newlist, rand.Intn(10000))
		}
		matrix = append(matrix, newlist)
	}
	var matrix2 = make([][]int, 10)
	for i := range matrix {
		matrix2[i] = make([]int, len(matrix[i]))
		copy(matrix2[i], matrix[i])
	}
	var start time.Time
	var elapsed time.Duration
	var start2 time.Time
	var elapsed2 time.Duration
	var govar int64
	var quickvar int64
	govar=0
	quickvar=0
	// x := []int{1, 5, 3, 4, 2}
	// goquicksort(x)
	// fmt.Println(x)
	for i := range matrix {
		start = time.Now()
		goquicksort(matrix[i])
		elapsed = time.Since(start)
		fmt.Println("go: " + elapsed.String())
		govar+=elapsed.Nanoseconds()
		start2 = time.Now()
		quicksort(matrix2[i])
		elapsed2 = time.Since(start2)
		fmt.Println("quicksort: " + elapsed2.String())
		quickvar+=elapsed2.Nanoseconds()

	}
	fmt.Println("avg go:",float64(govar)/10000000,"ms")
	fmt.Println("avg quicksort:",float64(quickvar)/10000000,"ms")
		


}
func partition(list []int, left int, right int) int {
	pivot := left
	left++
	for left <= right {
		for left <= right && list[left] <= list[pivot] {
			left++
		}
		for right >= left && list[right] >= list[pivot] {
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
	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
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

//////////////////////////////////////////////////////////////////////////////

func gopartition(list []int, left int, right int) int {
	pivot := left
	left++
	

	for left <= right {
		for left <= right && list[left] <= list[pivot] {
			left++
		}
		for right >= left && list[right] >= list[pivot] {
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
func goquicksort(list []int) {
	//shuffling list to remove worst case complexity
	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	pivot1:=gopartition(list,0,len(list)-1)
	pivot2:=gopartition(list,0,pivot1-1)
	pivot3:=gopartition(list,pivot1+1,len(list)-1)
	q1:=list[0:pivot2]
	q2:=list[pivot2:pivot1]
	q3:=list[pivot1:pivot3]
	q4:=list[pivot3:len(list)]
	var waiter sync.WaitGroup
	waiter.Add(3)
	go func(){
		goquicksorter(q1,0,len(q1)-1)
		waiter.Done()
	}()
	go func(){
		goquicksorter(q2,0,len(q2)-1)
		waiter.Done()
	}()
	go func(){
		goquicksorter(q3,0,len(q3)-1)
		waiter.Done()
	}()
	goquicksorter(q4,0,len(q4)-1)
	waiter.Wait()
}
func goquicksorter(list []int, start int, end int) {
	if start >= end || end <= 0 {
		return
	}
	pivot := gopartition(list, start, end)
	goquicksorter(list, start, pivot-1)
	goquicksorter(list, pivot+1, end)
	

}
///alternate version(seems slower)
// if sentinal{
// 	var waiter sync.WaitGroup
// 	waiter.Add(1)
// 	go func(){
// 		goquicksorter(list, start, pivot-1,false)
// 		waiter.Done()
// 	}()
// 	goquicksorter(list, pivot+1, end,false)
// 	waiter.Wait()
// }else{

// goquicksorter(list, start, pivot-1,false)
// goquicksorter(list, pivot+1, end,false)}
//TODO: try parralel mergesort
