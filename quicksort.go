package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// x := []int{82, 9, 90, 27, 37, 40, 24, 23, 96, 76, 88, 38, 28, 64, 57, 85, 7, 83, 95, 58, 39, 79, 6, 43, 92, 46, 20, 91, 51, 35, 68, 94, 52, 32, 93, 53, 98, 45, 84, 21, 14, 97, 74, 75, 2, 99, 49, 77, 60, 16}
	y := [][]int{
		{380, 366, 239, 123, 59, 204, 192, 50, 148, 429, 472, 316, 49, 445, 124, 321, 326, 60, 257, 390, 278, 119, 74, 51, 100, 15,
			473, 365, 263, 466, 155, 425, 469, 258, 393, 500, 484, 65, 77, 8, 367, 101, 31, 310, 12, 28, 249, 256, 428, 483, 9, 173, 443,
			194, 451, 363, 127, 61, 207, 157, 205, 37, 329, 314, 296, 55, 109, 159, 92, 184, 64, 46, 116, 305, 376, 343, 306, 171, 84, 358,
			353, 464, 114, 324, 230, 69, 217, 80, 115, 374, 420, 209, 497, 262, 401, 379, 457, 11, 480, 437},
		{96, 308, 170, 411, 217, 150, 390, 469, 131, 101, 113, 75, 56, 446, 281, 185, 421, 53, 137, 360, 115, 245, 466, 223, 379, 160, 471,
			425, 378, 65, 193, 43, 457, 361, 235, 312, 110, 341, 477, 479, 397, 177, 64, 298,
			178, 181, 73, 12, 179, 144, 487, 127, 280, 36, 74, 494, 145, 45, 324, 440, 407,
			270, 384, 255, 265, 331, 426, 409, 261, 355, 329,
			23, 68, 435, 500, 37, 462, 182, 136, 284, 79, 277, 496, 240, 263, 130, 278, 333, 314, 475, 201, 323, 304, 391, 171, 172, 11, 395, 52, 212},
		{117, 406, 58, 304, 381, 374, 35, 214, 261, 354, 419, 23, 467, 180, 44, 209, 31, 277, 334, 496, 55, 494, 228, 167, 367,
			397, 431, 102, 176, 28, 199, 417, 482, 1, 80, 244, 37, 357, 46, 84, 202, 493, 250, 459, 193, 264, 11, 289, 408, 211, 427, 48,
			410, 116, 182, 153, 356, 13, 95, 474, 393, 27, 382, 378, 384, 324, 258, 267, 179, 36, 383, 479, 130,
			463, 454, 17, 164, 126, 456, 318, 51, 425, 247, 422, 307, 252, 82, 94, 306, 174, 2, 12, 489, 100, 41, 272, 296, 305, 155, 248},
		{387, 315, 173, 201, 454, 397, 82, 224, 398, 371, 251, 370, 356, 298, 245, 19, 481, 193,
			21, 213, 69, 129, 122, 162, 212, 287, 241, 77, 95, 47, 222, 87, 231, 462, 460, 39, 299, 259, 357, 217, 341, 174, 256, 362, 123, 177,
			452, 214, 10, 283, 189, 255, 20, 53, 347, 98, 37, 342, 496, 488, 409, 445, 500, 99,
			484, 92, 2, 467, 366, 343, 149, 233, 422, 84, 25, 52, 336, 304, 270, 302, 145, 183, 216, 413, 13, 360, 325, 166, 404, 289, 411, 28, 274, 326, 3, 197, 261, 161, 291, 16},
		{445, 95, 26, 431, 495, 430, 215, 313, 352, 414, 118, 451, 208, 69, 449, 136, 24, 84, 365, 466, 112, 331,
			10, 72, 40, 198, 226, 452, 232, 162, 267, 216, 122, 78, 397, 66, 80, 70, 441, 17, 85, 436, 388, 107, 130, 121, 291, 288, 139, 494, 152, 218, 389, 304, 104, 160, 138, 374, 28, 234,
			257, 79, 205, 62, 423, 134, 274, 273, 213, 311, 376, 56, 435, 98, 457, 192, 187, 488, 396, 175,
			161, 281, 110, 323, 27, 191, 393, 128, 19, 474, 478, 336, 3, 460, 158, 305, 262, 500, 32, 421},
	}
	for _, value := range y {
		quicksort(value)
		fmt.Println(value)
	}
	// x := []int{1, 5, 3, 4, 2}
	// x := []int{5, 1, 2, 3, 4, 6}
	// goquicksort(x)
	// fmt.Println(x)

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

//////////////////////////////////////////////////////////////////////////////

func gopartition(list []int, left int, right int) int {

	pivot := left
	left++
	for left <= right {
		var waiter sync.WaitGroup
		waiter.Add(1)
		l := make(chan int)
		go func(list []int, left int, right int, pivot int, l chan int) {
			for left <= right && list[left] < list[pivot] {
				left++
			}
			waiter.Done()
			l <- left
		}(list, left, right, pivot, l)
		for right >= left && list[right] > list[pivot] {
			right--
		}
		waiter.Wait()
		left = <-l
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
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	goquicksorter(list, 0, len(list)-1)
}
func goquicksorter(list []int, start int, end int) {
	if start >= end || end <= 0 {
		return
	}
	pivot := gopartition(list, start, end)
	goquicksorter(list, start, pivot-1)
	goquicksorter(list, pivot+1, end)

}
