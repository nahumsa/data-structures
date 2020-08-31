package btree

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	bt := BinarySearchTree{}
	bt.InsertElement(5, 5)
	bt.InsertElement(7, 1)
	want := []bool{true, false, true, false}
	numbers := []int{5, 1, 7, 2}
	for i := 0; i < len(numbers); i++ {
		if bt.SearchNode(numbers[i]) != want[i] {
			t.Errorf("Expected was %v, we got %v", want[i], numbers[i])
		}
	}
}

func TestBinaryTreeMinMax(t *testing.T) {

	bt := BinarySearchTree{}
	list := rand.Perm(100)
	for _, val := range list {
		bt.InsertElement(val, val)
	}
	min := 0
	max := 99

	if *bt.MinNode() != min {
		t.Errorf("Min: Expected was %v, we got %v", min, *bt.MinNode())
	}

	if *bt.MaxNode() != max {
		t.Errorf("Max: Expected was %v, we got %v", max, *bt.MaxNode())
	}

}
func BenchmarkSearch(b *testing.B) {
	searchBenchMark(b)
}

func searchBenchMark(b *testing.B) {
	for _, size := range []int{
		100, 200, 400, 800, 1600, 3200,
	} {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				b.StopTimer()
				list := rand.Perm(size)
				bt := BinarySearchTree{}
				for _, val := range list {
					bt.InsertElement(val, val)
				}
				find := rand.Intn(size)
				b.StartTimer()
				bt.SearchNode(find)
			}
		})
	}
}
