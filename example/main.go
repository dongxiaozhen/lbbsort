package main

import (
	"fmt"

	"github.com/dongxiaozhen/sort"
)

func main() {
	fmt.Println("vim-go")
	a := []int{90, 10, 50, 40, 60, 20, 70, 80, 30}
	fmt.Println(a)
	sort.BubbleSort(a)
	fmt.Println(a)
	a = []int{90, 10, 50, 40, 60, 20, 70, 80, 30}
	sort.QuickSort(a)
	fmt.Println(a)
	a = []int{90, 10, 50, 40, 60, 20, 70, 80, 30}
	sort.ShellSort(a)
	fmt.Println(a)
	a = []int{90, 10, 50, 40, 60, 20, 70, 80, 30}
	sort.HeapSort(a)
	fmt.Println(a)
	// sort.Sort()
}
