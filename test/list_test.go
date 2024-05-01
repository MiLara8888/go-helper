package test

import (
	"fmt"
	slicelist "github.com/MiLara8888/go-helper/listnode/slice_list"
	"testing"
)

func TestList(t *testing.T) {

    l := slicelist.List{}
	l.Fill(11)
	res := slicelist.MergeSort(l.List)
	for _, j := range res{
		fmt.Println(j)
	}

}

func BenchmarkBubbleList(b *testing.B) {
    for i := 0; i < b.N; i++ {
    l := slicelist.List{}
	l.Fill(10_000)
	l.Bublsort()
    }
}

func BenchmarkInsertionList(b *testing.B) {
    for i := 0; i < b.N; i++ {
    l := slicelist.List{}
	l.Fill(10_000)
	l.InsertionSort()
    }
}

func BenchmarkMergeList(b *testing.B) {
    for i := 0; i < b.N; i++ {
    l := slicelist.List{}
	l.Fill(10_000)
	_ = slicelist.MergeSort(l.List)
    }
}