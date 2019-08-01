package main

import (
	"fmt"
	"math/rand"
	"time"
)

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
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

func sortArray(array [5]int) [5]int {
	modArray := array[:]
	var returningArray [5]int
	for i := 0; i < 5; i++ {
		x := modArray[0]
		var in int
		for index, value := range modArray {
			if x > value {
				x = value
				in = index
			}
		}
		modArray = remove(modArray, in)
		returningArray[i] = x
	}
	return returningArray
}

func main() {
	x := randomArray()
	fmt.Println(x)
	fmt.Println(sortArray(x))
}
