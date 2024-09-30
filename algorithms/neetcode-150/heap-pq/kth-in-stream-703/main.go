package main

type KthLargest struct {
	k    int
	heap []int
}

func heapify(nums []int) {

}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{
		k:    k,
		heap: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	return 69
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
