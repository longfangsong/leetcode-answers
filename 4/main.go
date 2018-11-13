package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	left := 0
	right := len(nums1)
	halfLength := (len(nums1) + len(nums2) + 1) / 2
	for left <= right {
		partitionNums1At := (left + right) / 2
		partitionNums2At := halfLength - partitionNums1At
		if partitionNums1At < right && nums2[partitionNums2At-1] > nums1[partitionNums1At] {
			left = partitionNums1At + 1
		} else if partitionNums1At > left && nums1[partitionNums1At-1] > nums2[partitionNums2At] {
			right = partitionNums1At - 1
		} else {
			var maxOfLeftside int
			if partitionNums1At == 0 {
				maxOfLeftside = nums2[partitionNums2At-1]
			} else if partitionNums2At == 0 {
				maxOfLeftside = nums1[partitionNums1At-1]
			} else {
				maxOfLeftside = max(nums1[partitionNums1At-1], nums2[partitionNums2At-1])
			}
			if (len(nums1)+len(nums2))%2 == 1 {
				return float64(maxOfLeftside)
			}
			var minOfRightside int
			if partitionNums1At == len(nums1) {
				minOfRightside = nums2[partitionNums2At]
			} else if partitionNums2At == len(nums2) {
				minOfRightside = nums1[partitionNums1At]
			} else {
				minOfRightside = min(nums1[partitionNums1At], nums2[partitionNums2At])
			}
			return float64(maxOfLeftside+minOfRightside) / 2.0
		}
	}
	return -1
}

func main() {

}
