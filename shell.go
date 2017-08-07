package lbbsort

// 0-----2----4---6---8
// 8和6比较，6和4比较，4和2比较，，
func ShellSort(nums []int) {
	len := len(nums)
	terval := len / 2

	for terval > 0 {
		for i := terval; i < len; i++ {
			for j := i - terval; j >= 0; j -= terval {
				if nums[j] > nums[j+terval] {
					nums[j], nums[j+terval] = nums[j+terval], nums[j]
				}
			}
		}
		terval /= 2
	}
}
