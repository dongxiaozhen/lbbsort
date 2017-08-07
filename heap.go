package lbbsort

func HeapSort(nums []int) {
	buildHeap(nums)
	for i := len(nums) - 1; i >= 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapAdjust(nums, 0, i-1)
	}
}

func buildHeap(nums []int) {
	len := len(nums) - 1
	for i := len / 2; i >= 0; i-- {
		heapAdjust(nums, i, len)
	}
}

// 找出三个节点里最大一个交换数据，然后递归
func heapAdjust(nums []int, top int, size int) {
	left := 2*top + 1
	right := 2*top + 2
	max := top
	if left <= size && nums[left] > nums[top] {
		max = left
	}
	if right <= size && nums[right] > nums[max] {
		max = right
	}
	if max != top {
		nums[top], nums[max] = nums[max], nums[top]
		heapAdjust(nums, max, size)
	}
}
