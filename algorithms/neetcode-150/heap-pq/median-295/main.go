package main

import (
	"container/heap"
	"fmt"
)

type MedianFinder struct {
	SmallPile *MayHeap
	BigPile   *MayHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		SmallPile: MaxHeap(),
		BigPile:   MinHeap(),
	}
}

func (this *MedianFinder) diffPiles() (int, *MayHeap, *MayHeap) {
	/// returns diff, larger, smaller
	raw := this.SmallPile.Len() - this.BigPile.Len()
	if raw < 0 {
		return -raw, this.BigPile, this.SmallPile
	}
	return raw, this.SmallPile, this.BigPile
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.SmallPile, num)
	if this.BigPile.Len() > 0 {
		if this.BigPile.Peek().(int) < this.SmallPile.Peek().(int) {
			out := heap.Pop(this.SmallPile)
			heap.Push(this.BigPile, out)
		}
	}
	nDiff, larger, smaller := this.diffPiles()
	if nDiff > 1 {
		heap.Push(smaller, heap.Pop(larger))
	}

}

func (this *MedianFinder) FindMedian() float64 {
	nDiff, _, _ := this.diffPiles()
	if nDiff == 0 {
		small, big := float64(this.SmallPile.Peek().(int)), float64(this.BigPile.Peek().(int))
		return (small + big) / 2.0
	}
	if this.SmallPile.Len() > 0 {
		return float64(this.SmallPile.Peek().(int))
	}
	return float64(this.BigPile.Peek().(int))
}

func main() {
	median := Constructor()
	median.AddNum(-1)
	median.AddNum(-2)
	median.AddNum(-3)
	median.AddNum(-4)
	median.AddNum(-5)
	fmt.Println(median.SmallPile,
		median.BigPile)
	fmt.Println(median.FindMedian())
}
