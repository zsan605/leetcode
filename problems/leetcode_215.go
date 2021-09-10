package problems

func findKthLargest(nums []int, k int) int {
	// 构建大根堆
	headSize := len(nums)
	buildMaxHeap(nums, headSize)
	for i := 1; i < k; i++ {
		nums[0], nums[headSize-i] = nums[headSize-i], nums[0]
		maxHeap(nums, 0, headSize-i)
	}

	return nums[0]
}

func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeap(nums, i, heapSize)
	}
}

func maxHeap(nums []int, i, heapSize int) {
	l, r, largest := 2*i+1, 2*i+2, i

	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}

	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}

	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeap(nums, largest, heapSize)
	}
}
