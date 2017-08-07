package lbbsort

func QuickSort(nums []int) {
	qsort(nums, 0, len(nums)-1)
}
func qsort(array []int, low, high int) {
	if low < high {
		m := partition(array, low, high)
		// fmt.Println(m)
		qsort(array, low, m-1)
		qsort(array, m+1, high)
	}
}

// ------------key-----------
//     <=key              >key
func partition(array []int, low, high int) int {
	key := array[low]
	tmpLow := low
	tmpHigh := high
	for {
		//查找小于等于key的元素，该元素的位置一定是tmpLow到high之间，因为array[tmpLow]及左边元素小于等于key，不会越界
		for array[tmpHigh] > key {
			tmpHigh--
		}
		//找到大于key的元素，该元素的位置一定是low到tmpHigh+1之间。因为array[tmpHigh+1]必定大于key
		for array[tmpLow] <= key && tmpLow < tmpHigh {
			tmpLow++
		}

		if tmpLow >= tmpHigh {
			break
		}
		array[tmpLow], array[tmpHigh] = array[tmpHigh], array[tmpLow]
	}
	array[tmpLow], array[low] = array[low], array[tmpLow]
	return tmpLow
}
