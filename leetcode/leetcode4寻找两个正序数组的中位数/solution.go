/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-08 11:39:38
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-08 16:06:07
 */
package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	left, right := (m+n+1)/2, (m+n+2)/2
	return float64(findKth(nums1, 0, nums2, 0, left) + findKth(nums1, 0, nums2, 0, right))
}

func findKth(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}
	if k == 1 {
		return min(nums1[i], nums2[j])
	}
	var midVal1, midVal2 int
	if i+k/2-1 < len(nums1) {
		midVal1 = nums1[i+k/2-1]
	} else {
		midVal1 = math.MaxInt
	}
	if j+k/2-1 < len(nums2) {
		midVal2 = nums2[j+k/2-1]
	} else {
		midVal2 = math.MaxInt
	}

	if midVal1 < midVal2 {
		return findKth(nums1, i+k/2, nums2, j, k-k/2)
	}
	return findKth(nums1, i, nums2, j+k/2, k-k/2)

}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
