package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func printMemStat(ms runtime.MemStats) {

	runtime.ReadMemStats(&ms)
	fmt.Println("--------------------------------------")
	fmt.Println("Memory Statistics Reporting time: ", time.Now())
	fmt.Println("--------------------------------------")
	fmt.Println("Bytes of allocated heap objects: ", ms.Alloc)
	fmt.Println("Total accumulative bytes of Heap object: ", ms.TotalAlloc)
	fmt.Println("Bytes of memory obtained from OS: ", ms.Sys)
	fmt.Println("Count of heap objects: ", ms.Mallocs)
	fmt.Println("Count of heap objects freed: ", ms.Frees)
	fmt.Println("Count of live heap objects", ms.Mallocs-ms.Frees)
	fmt.Println("Number of completed GC cycles: ", ms.NumGC)
	fmt.Println("--------------------------------------")
}

func allocateSlice() {

	intArr := make([]int, 900000)
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Int()
	}

}

func allocateArray() {

	intArr := new([900000]int)
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Int()
	}
}

func test_garbage_collection() {

	var ms runtime.MemStats

	printMemStat(ms)

	for i := 0; i < 100; i++ {
		fmt.Println()
		fmt.Println("Loop No.", i+1)

		allocateSlice()
		allocateArray()

		time.Sleep(1 * time.Second)

		printMemStat(ms)
	}

	fmt.Println()
	fmt.Println("This demonstrated that both new() and make() objects are collectable")

}

func test_array_pointer() {

	var a [10]int

	for i := 0; i < len(a); i++ {
		(&a)[i] = i
	}

	fmt.Println((&a)[4])

	a[4] = 1234

	fmt.Println((&a)[4])

	fmt.Println("It looks like access entries in an array through a pointer or the array have the same effect")

}

func test_slice_pointer() {

	var a []int

	for i := 0; i < 10; i++ {
		a = append(a, i)
	}

	fmt.Println(a[4])

	fmt.Println("Accessing element of a slice though a pointer to the slice it not allowed")
}

func main() {

	run_all := false
	if run_all {
		test_garbage_collection()
		test_array_pointer()
	}

	test_slice_pointer()

}
