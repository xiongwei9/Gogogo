package al

import "testing"

func TestSubArray(t *testing.T) {
	arr := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	left, right, sum := maxSubarray(arr)
	if left != 7 || right != 10 || sum != 43 {
		t.Error("error", left, right, sum)
	}
}

// 求解最大子数组
func maxSubarray(arr []int) (int, int, int) {
	size := len(arr)
	return divideMaxSubarray(arr, 0, size-1)
}

// 使用分治法求解
func divideMaxSubarray(arr []int, left, right int) (int, int, int) {
	if left == right {
		return left, right, arr[left]
	}

	mid := (left + right) / 2
	leftLow, leftHigh, leftSum := divideMaxSubarray(arr, left, mid)
	rightLow, rightHigh, rightSum := divideMaxSubarray(arr, mid+1, right)
	crossLow, crossHigh, crossSum := maxCrossingArray(arr, left, right)

	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}

func maxCrossingArray(arr []int, left, right int) (int, int, int) {
	var sum int
	mid := (left + right) / 2

	sum = 0
	leftSum := arr[mid]
	var leftMax int
	for i := mid; i >= 0; i-- { // 从中间向左遍历，查找最大点
		sum += arr[i]
		if sum >= leftSum {
			leftSum = sum
			leftMax = i
		}
	}

	if mid+1 > right {
		return leftMax, mid, leftSum
	}

	sum = 0
	rightSum := arr[mid+1]
	var rightMax int
	for i := mid + 1; i <= right; i++ { // 从中间向右遍历，查找最大点
		sum += arr[i]
		if sum >= rightSum {
			rightSum = sum
			rightMax = i
		}
	}
	// 返回跨越中点的最大子序列
	return leftMax, rightMax, leftSum + rightSum
}
