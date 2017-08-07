package main

import (
	"fmt"

	"github.com/dongxiaozhen/lbbsort"
)

func main() {
	// var ids []uint32
	var ids lbbsort.Uint32Slice
	ids = append(ids, 4)
	ids = append(ids, 3)
	ids = append(ids, 2)
	// lbbsort.Uint32s(ids)
	ids.Sort()
	fmt.Println(ids)
}
