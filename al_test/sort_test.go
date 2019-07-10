package al_test

import (
	"testing"
)

func TestSort(t *testing.T) {
	list := [][]int{
		{1, 3, 5, 6, 2, 5, 4, 8, 7},
		{9, 3, 5, 7, 12, 1, 4, 2, 8},
		{4, 5, 2, 6, 7, 3, 1, 9},
		{10, 3, 5, 6, 2, 7, 5, 4, 7, 88},
	}
	//sortFns := []func([]int)([]int){insertionSort, mergeSort, quickSort}
	for index, arr := range list {
		//nums := insertionSort(arr) // 插入排序
		//nums := mergeSort(arr) // 归并排序
		nums := quickSort(arr) // 快速排序
		//nums := heapSort(arr) //堆排序

		if len(nums) != len(arr) {
			t.Error("array length error", index, nums)
			continue
		}
		t.Log(nums)
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[i-1] {
				t.Error(index, nums)
				break
			}
		}
	}
}

// 堆排序
//func heapSort(arr []int) []int {
//
//	return nil
//}
//
//func buildHeap(arr []int) {
//	length := len(arr)
//	for i := (length - 1) / 2; i >= 0; i-- {
//		adjustHeap(arr, i)
//	}
//}
//
//func adjustHeap(arr []int, pos int) {
//	length := len(arr)
//	for child := pos*2 + 1; child < length; child = pos*2 + 1 {
//	}
//}

// 快速排序
func quickSort(arr []int) []int {
	length := len(arr)
	nums := make([]int, length, length)
	copy(nums, arr)
	if length >= 2 {
		quickSortKernel(nums, 0, length-1)
	}

	return nums
}

func quickSortKernel(arr []int, left, right int) {
	if left >= right {
		return
	}
	key := arr[left]
	l, r := left, right

	for l < r {
		for l < r && key <= arr[r] { // 从右向左扫描，发现有小于key的数，就把它移到左边
			r--
		}
		arr[l] = arr[r]
		for l < r && key >= arr[l] { // 从左向右扫描，发现有大于key的数，就把它移到右边
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = key
	quickSortKernel(arr, left, l-1)
	quickSortKernel(arr, l+1, right)
}

// 归并排序
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

// 插入排序
func insertionSort(arr []int) []int {
	length := len(arr)
	nums := make([]int, length, length)
	copy(nums, arr)
	for i := 1; i < length; i++ {
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
	nums1 := []int{1, 2, 4, 5, 6, 7, 8, 11, 13}
	nums2 := []int{1, 2, 4, 5, 6, 7, 8, 11, 13, 14}
	list := []struct {
		nums   []int
		target int
		result bool
	}{
		{nums1, 0, false},
		{nums1, 2, true},
		{nums1, 1, true},
		{nums1, 3, false},
		{nums1, 4, true},
		{nums1, 7, true},
		{nums1, 11, true},
		{nums1, 13, true},
		{nums1, 15, false},
		{nums2, 3, false},
		{nums2, 2, true},
		{nums2, 11, true},
		{nums2, 8, true},
	}
	for _, item := range list {
		if binarySearch(item.nums, item.target) != item.result {
			t.Error(item.nums, item.target)
		}
	}
}

// 二分查找
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
