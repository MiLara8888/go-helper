package test

import (
	"testing"

	listnode "github.com/MiLara8888/go-helper/list_node"
)

func TestInsertion(t *testing.T) {
    l := listnode.LinkedList{}
	l.Fill(10)
	l.InsertionSort()
    l.Head.Console()
}


func TestBubble(t *testing.T) {
    l := listnode.LinkedList{}
	l.Fill(10)
	l.BubbleSort(10)
    l.Head.Console()
}

func BenchmarkBubbleLNode(b *testing.B) {
    for i := 0; i < b.N; i++ {
    l := listnode.LinkedList{}
	l.Fill(10_000)
	l.BubbleSort(10_000)

    }
}
func BenchmarkBubbleLNode2(b *testing.B) {
    for i := 0; i < b.N; i++ {
    l := listnode.LinkedList{}
	l.Fill(100_000)
	l.BubbleSort(100_000)

    }
}
func BenchmarkBubbleLNode3(b *testing.B) {
    for i := 0; i < b.N; i++ {
    l := listnode.LinkedList{}
	l.Fill(1_000_000)
	l.BubbleSort(1_000_000)
    }
}

func BenchmarkInsert1(b *testing.B) {
    for i := 0; i < b.N; i++ {
    l := listnode.LinkedList{}
	l.Fill(10_000)
	l.InsertionSort()

    }
}
func BenchmarkInsert2(b *testing.B) {
    for i := 0; i < b.N; i++ {
    l := listnode.LinkedList{}
	l.Fill(100_000)
	l.InsertionSort()

    }
}
func BenchmarkInsert3(b *testing.B) {
    for i := 0; i < b.N; i++ {
    l := listnode.LinkedList{}
	l.Fill(1_000_000)
	l.InsertionSort()
    }
}