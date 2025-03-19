package main

import (
	"fmt"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// 연속 메모리 [한번에 메모리 할당과 해제를 수행]
// Random Access에 강력
// 삽입 및 삭제에 비효율적 [배열의 시작과 끝에 데이터를 추가 혹은 삭제는 양호]
// Cache Miss가 발생 확률이 낮음
// 요소가 사라질 때마다 GC 되지 않는다
////////////////////////////////////////////////////////////////////////////////

var start = time.Now()

func staticArray() {
	var arr_type_a = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Nomal array type a: %d, %d %v\n", len(arr_type_a), cap(arr_type_a), arr_type_a)

	var arr_type_b = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Nomal array type b: %d, %d %v\n", len(arr_type_b), cap(arr_type_b), arr_type_b)

	var arr_type_c = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Nomal array type c: %d, %d %v\n", len(arr_type_c), cap(arr_type_c), arr_type_c)
}

func emptyArray() {
	var empty_arr = []int{} // Non initialize array length 0
	fmt.Printf("Empty array: %d, %d %v\n", len(empty_arr), cap(empty_arr), empty_arr)
}

func copyArray() {
	var base_arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var copy_arr [10]int

	copy(copy_arr[0:], base_arr[0:5]) // Copy from Base array in the index copy 0 to 5
	copy_arr[5] = 100                 // Change copy array index length 5 to value 100
	copy(copy_arr[6:], base_arr[5:])  // Copy from Base array index length 5 to end

	fmt.Printf("Base array: %d, %d %v\n", len(base_arr), cap(base_arr), base_arr)
	fmt.Printf("Copy array: %d, %d %v\n", len(copy_arr), cap(copy_arr), copy_arr)
}

func dynamicArray() {
	var create_arr = make([]int, 10, 15) // Create array length 10, capacity 15
	fmt.Printf("Create array: %d, %d %v\n", len(create_arr), cap(create_arr), create_arr)

	var add_arr = append(create_arr, 1, 2, 3, 4, 5, 6) // Add array index length
	fmt.Printf("Add array: %d, %d %v\n", len(add_arr), cap(add_arr), add_arr)
}

func main() {
	staticArray()
	emptyArray()
	copyArray()
	dynamicArray()

	compile_time := time.Since(start)
	fmt.Printf("Compile time: %s \n", compile_time)
}
