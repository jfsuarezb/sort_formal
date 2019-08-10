package main

import (
	"fmt"
	"math/rand"
	"time"
)

func remove(slice []int, s int) []int {
	if len(slice)-1 == s {
		return slice[:s]
	}
	x := append(slice[:s], slice[s+1:]...)
	return x
}

func randomArray() [5]int {
	var randomArray [5]int
	n := 0
	for n < 5 {
		rand.Seed(time.Now().UnixNano())
		randomArray[n] = rand.Intn(10)
		n++
	}
	return randomArray
}

func heapsort(array [5]int) []int {
	var sorted []int
	slice := array[:]
	//fmt.Println(slice)
	var heap []int
	var element int
	fmt.Println(heap)
	for i := 0; i < len(array); i++ {
		if i == 0 {
			heap = maximise(slice)
		} else {
			heap = maximise(heap)
		}
		//fmt.Println(heap, i)
		if len(heap) == 1 {
			sorted = append(sorted, heap[0])
			break
		}
		heap, element = swap(heap)
		//fmt.Println(heap, element)
		sorted = append(sorted, element)
		//fmt.Println(heap, i, "lkfjldajf")
		//fmt.Println("....")
	}
	return sorted
}

func maximise(slice []int) []int {
	usable := slice
	var returning []int
	var element, index int
	for i := 0; i < len(slice); i++ {
		if len(usable) == 1 {
			returning = append(returning, usable[0])
			return returning
		}
		element, index = searchmax(usable)
		usable = remove(usable, index)
		returning = append(returning, element)
	}
	return returning
}

func swap(slice []int) ([]int, int) {
	maxel, maxin := searchmax(slice)
	//fmt.Println(maxel)
	minel, minin := searchmin(slice)
	//fmt.Println(minel)
	usable := slice
	//fmt.Println(slice)
	usable = remove(usable, minin)
	//fmt.Println(usable)
	usable[maxin] = minel
	//fmt.Println(usable)
	//fmt.Println(".......")
	return usable, maxel
}

func searchmax(slice []int) (int, int) {
	x := 0
	in := 0
	for index, element := range slice {
		if element > x {
			x = element
			in = index
		}
	}
	return x, in
}

func searchmin(slice []int) (int, int) {
	x := slice[0]
	in := 0
	for index, element := range slice {
		if element < x {
			x = element
			in = index
		}
	}
	return x, in
}

func main() {
	x := randomArray()
	fmt.Println(x)
	//fmt.Println(searchmin(x[:]))
	sorted := heapsort(x)
	fmt.Println(sorted)
}
