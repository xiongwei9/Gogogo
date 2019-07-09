package al_test

import (
	"testing"
)

func TestSort(t *testing.T) {
	list := [][]int{
		{1, 3, 5, 6, 2, 5, 4, 8, 7},
		{9, 3, 5, 7, 12, 1, 4, 2, 8},
	}
	for index, arr := range list {
		//nums := insertionSort(arr) // 插入排序
		nums := mergeSort(arr) // 归并排序
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[i-1] {
				t.Error(index, nums)
			}
		}
	}
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	mid := length / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	lLen, rLen := len(left), len(right)
	l, r := 0, 0
	result := make([]int, 0, lLen+rLen)
	for l < lLen && r < rLen {
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

func insertionSort(arr []int) []int {
	nums := make([]int, 0, len(arr))
	copy(nums, arr)
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		var j int
		for j = i - 1; j >= 0 && nums[j] > key; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = key
	}
	return nums
}

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 4, 5, 6, 7, 8, 11, 13}
	list := []struct {
		nums   []int
		target int
		result bool
	}{
		{nums, 2, true},
		{nums, 1, true},
		{nums, 3, false},
		{nums, 4, true},
		{nums, 7, true},
		{nums, 11, true},
		{nums, 13, true},
		{nums, 15, false},
	}
	for _, item := range list {
		if binarySearch(item.nums, item.target) != item.result {
			t.Error(item.nums, item.target)
		}
	}
}

func binarySearch(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	if target < nums[l] || target > nums[r] {
		return false
	}
	for l <= r {
		m := (l + r) / 2
		if target == nums[m] {
			return true
		} else if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}
