package main

import (
	"fmt"
	"math/rand"
	"runtime"
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
	govar = 0
	quickvar = 0
	// x := []int{1, 5, 3, 4, 2}
	// goquicksort(x)
	// fmt.Println(x)
	for i := range matrix {
		start = time.Now()
		goquicksort(matrix[i])
		elapsed = time.Since(start)
		fmt.Println("go: " + elapsed.String())
		sortcheck(matrix[i])
		govar += elapsed.Nanoseconds()
		start2 = time.Now()
		quicksort(matrix2[i])
		elapsed2 = time.Since(start2)
		fmt.Println("quicksort: " + elapsed2.String())
		quickvar += elapsed2.Nanoseconds()

	}
	fmt.Println("avg go:", float64(govar)/10000000, "ms")
	fmt.Println("avg quicksort:", float64(quickvar)/10000000, "ms")
	fmt.Println("percentage difference:", float32((float64(quickvar)/10000000)/(float64(govar)/10000000)),"%")
	fmt.Println("available cores: ", runtime.GOMAXPROCS(0))

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
	// shuffling list to remove worst case complexity
	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	goquicksorter(list, 0, len(list)-1)
	

}
func goquicksorter(list []int, start int, end int) {
	if start >= end || end <= 0 {
		return
	}
	pivot := gopartition(list, start, end)
	var waiter sync.WaitGroup
	if ((pivot - 1) - start) > 350 {
		waiter.Add(1)
		temp := list[start:pivot]
		go func() {
			goquicksorter(temp, 0, len(temp)-1)
			waiter.Done()
		}()
	} else {
		goquicksorter(list, start, pivot-1)
	}
	if (end - pivot + 1) > 350 {
		waiter.Add(1)
		temp := list[pivot+1 : end+1]
		go func() {
			goquicksorter(temp, 0, len(temp)-1)
			waiter.Done()
		}()
	} else {
		goquicksorter(list, pivot+1, end)
	}
	waiter.Wait()

}

func sortcheck(list []int) {
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println("sorting error")
			fmt.Println(i)
		}
	}

}
