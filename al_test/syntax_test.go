package al_test

import (
	"fmt"
	"testing"
)

func TestSyntax(t *testing.T) {
	arr := make([]int, 3, 3)
	fmt.Println(arr, len(arr), cap(arr))
	//arr = append(arr, 1)
	//fmt.Println(arr, len(arr), cap(arr))
	arr2 := []int{1, 2, 3}
	fmt.Println(arr2, len(arr2), cap(arr2))
	copy(arr, arr2)
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Println(arr2, len(arr2), cap(arr2))
}
