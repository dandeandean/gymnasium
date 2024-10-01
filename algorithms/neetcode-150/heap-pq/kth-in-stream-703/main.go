package main

type KthLargest struct {
	k    int
	heap []int
}

func heapify(nums []int, i int) {
	left, right := i*2+1, i*2+2
	if right >= len(nums) {
		return
	}
	smallest := 0
	if nums[left] < nums[smallest] {
		smallest = left
	}
	if nums[right] < nums[smallest] {
		smallest = right
	}
	if smallest != 0 {
		nums[i], nums[smallest] = nums[smallest], nums[i]
		heapify(nums, smallest)
	}

}

func Constructor(k int, nums []int) KthLargest {
	for len(nums) > k {
		nums = nums[1:]
		heapify(nums, 0)
	}
	return KthLargest{
		k:    k,
		heap: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	this.heap = append(this.heap, val)
	if len(this.heap) > this.k {
		heapify(this.heap, 0)
		this.heap = this.heap[1:]

	}
	heapify(this.heap, 0)
	return this.heap[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
