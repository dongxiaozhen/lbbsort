package main

import (
	"fmt"

	"github.com/dongxiaozhen/lbbsort"
)

func main() {
	fmt.Println("vim-go")
	a := []int{90, 10, 50, 40, 60, 20, 70, 80, 30}
	fmt.Println(a)
	lbbsort.BubbleSort(a)
	fmt.Println(a)
	a = []int{90, 10, 50, 40, 60, 20, 70, 80, 30}
	lbbsort.QuickSort(a)
	fmt.Println(a)
	a = []int{90, 10, 50, 40, 60, 20, 70, 80, 30}
	lbbsort.ShellSort(a)
	fmt.Println(a)
	a = []int{90, 10, 50, 40, 60, 20, 70, 80, 30}
	lbbsort.HeapSort(a)
	fmt.Println(a)
	// Sort.Sort()
}
