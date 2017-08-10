package lbbsort

import "sort"

type Uint32Slice []uint32

func (p Uint32Slice) Len() int           { return len(p) }
func (p Uint32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Uint32Slice) Sort()              { sort.Sort(p) }
func (p Uint32Slice) Search(x uint32) int {
	return sort.Search(len(p), func(i int) bool { return p[i] >= x })
}

func Uint32s(a []uint32) { sort.Sort(Uint32Slice(a)) }
