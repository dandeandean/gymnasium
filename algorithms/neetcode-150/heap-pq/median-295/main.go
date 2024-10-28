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
	if this.BigPile.Len() > 0 && this.BigPile.Peek().(int) < this.SmallPile.Peek().(int) {
		out := heap.Pop(this.SmallPile)
		heap.Push(this.BigPile, out)
	}
	if this.BigPile.Len() < this.SmallPile.Len()+1 {
		out := heap.Pop(this.SmallPile)
		heap.Push(this.BigPile, out)
	}
	if this.SmallPile.Len() > this.BigPile.Len()+1 {
		out := heap.Pop(this.BigPile)
		heap.Push(this.SmallPile, out)
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.diffPiles() == 0 {
		small, big := float64(this.SmallPile.Peek().(int)), float64(this.BigPile.Peek().(int))
		return (small + big) / 2.0
	}
	correctPile := this.SmallPile
	if this.BigPile.Len() > this.SmallPile.Len() {
		correctPile = this.BigPile
	}
	return float64(correctPile.Peek().(int))
}

func main() {
	median := Constructor()
	median.AddNum(1)
	fmt.Println(median.SmallPile, median.BigPile)
	median.AddNum(2)
	fmt.Println(median.SmallPile, median.BigPile)
	median.AddNum(3)
	fmt.Println(median.SmallPile, median.BigPile)
	median.AddNum(7)
	fmt.Println(median.SmallPile, median.BigPile)
}
