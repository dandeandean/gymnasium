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

func (this *MedianFinder) diffPiles() int {
	raw := this.SmallPile.Len() - this.BigPile.Len()
	if raw < 0 {
		return -raw
	}
	return raw
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.SmallPile, num)
	for this.diffPiles() > 1 && this.SmallPile.Len() > 0 {
		out := heap.Pop(this.SmallPile)
		heap.Push(this.BigPile, out)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.diffPiles() == 0 {
		small, big := this.SmallPile.Peek().(float64), this.BigPile.Peek().(float64)
		return (small + big) / 2.0
	}
	return float64(this.SmallPile.data[0])
}

func main() {
	median := Constructor()
	median.AddNum(1)
	median.AddNum(2)
	fmt.Println(median.FindMedian())
	median.AddNum(3)
	fmt.Println(median.FindMedian())
}
