package main

import "fmt"

type KthLargest struct {
	k    int
	heap []int
}

func heapify(nums []int, i int) {
	left, right := i*2+1, i*2+2
	smallest := i
	if left < len(nums) {
		if nums[left] < nums[smallest] {
			smallest = left
		}
	}
	if right < len(nums) {
		if nums[right] < nums[smallest] {
			smallest = right
		}
	}
	if smallest != i {
		nums[smallest], nums[i] = nums[i], nums[smallest]
		heapify(nums, smallest)
	}
}
func Constructor(k int, nums []int) KthLargest {
	for len(nums) > k {
		for i := len(nums)/2 + 1; i >= 0; i-- {
			heapify(nums, i)
		}
		nums = nums[1:]
	}
	return KthLargest{
		k:    k,
		heap: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	this.heap = append(this.heap, val)
	for i := len(this.heap) / 2; i >= 0; i-- {
		heapify(this.heap, i)
	}
	if len(this.heap) > this.k {
		this.heap = this.heap[1:]
	}
	for i := len(this.heap) / 2; i >= 0; i-- {
		heapify(this.heap, i)
	}
	return this.heap[0]
} /**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
func main() {
	arr := []int{4, 5, 8, 2}
	fmt.Println(arr)
	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, i)
	}
	fmt.Println(arr)
}
