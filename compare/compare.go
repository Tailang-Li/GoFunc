package main

import "fmt"

func main() {

	// compare between array
	fmt.Println("compare between array")
	arrayA := [2]int{1, 2}
	arrayB := [2]int{1, 2}
	fmt.Println(arrayA == arrayB)

	// compare between slice
	sliceA := []int{1, 2}
	sliceB := []int{1, 2}
	// slices cannot be compared with ==
	// fmt.Println(sliceA == sliceB) invalid operation

	// compare between map
	mapA := map[int]struct{}{
		1: {},
		2: {},
	}
	mapB := map[int]struct{}{
		1: {},
		2: {},
	}
	// maps cannot be compared with ==
	// fmt.Println(mapA == mapB) invalid operation

	// compare between channel
	fmt.Println("compare between channel")
	chanA := make(chan int)
	chanB := make(chan int)
	fmt.Println(chanA == chanB)

	// compare between struct with basic types
	fmt.Println("compare between struct with basic types")
	structA := struct {
		A int
		B string
	}{1, "1"}
	structB := struct {
		A int
		B string
	}{1, "1"}
	fmt.Println(structA == structB)

	// compare between struct with slice or map
	structC := struct {
		A []int
	}{sliceA}
	structD := struct {
		A []int
	}{sliceA}
	// structs cannot be compared if contain incomparable items
	// fmt.Println(structC == structD) invalid operation

	// compare between pointer
	fmt.Println("compare between pointer")
	fmt.Println(&arrayA == &arrayB)
	fmt.Println(&sliceA == &sliceB)
	fmt.Println(&mapA == &mapB)
	fmt.Println(&chanA == &chanB)
	fmt.Println(&structA == &structB)
	fmt.Println(&structC == &structD)
}
