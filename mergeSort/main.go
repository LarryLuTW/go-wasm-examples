package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func jsArrayToGoArray(jsArr js.Value) []int {
	n := jsArr.Length()
	goArr := make([]int, n)
	for i := 0; i < n; i++ {
		goArr[i] = jsArr.Index(i).Int()
	}
	return goArr
}

func goArrayToJsArray(goArr []int) []interface{} {
	jsArr := make([]interface{}, len(goArr))
	for i, v := range goArr {
		jsArr[i] = v
	}
	return jsArr
}

func _merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func _mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	mid := len(slice) / 2
	return _merge(_mergeSort(slice[:mid]), _mergeSort(slice[mid:]))
}

func mergeSort(this js.Value, inputs []js.Value) interface{} {
	goArr := jsArrayToGoArray(inputs[0])

	start := time.Now()

	goArr = _mergeSort(goArr)

	elapsed := time.Since(start)
	fmt.Printf("MergeSort in Go: %s\n", elapsed)

	return goArrayToJsArray(goArr)
}

func main() {
	js.Global().Set("mergeSortGo", js.FuncOf(mergeSort))
	select {} // block the main thread forever
}
