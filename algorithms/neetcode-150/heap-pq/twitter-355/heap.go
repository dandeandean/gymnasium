package main

type Set map[int]bool

type Tweet struct {
	Time   int
	ID     int
	UserID int
	TwInd  int
}

type Twitter struct {
	Time        int
	Users       Set             // Set of Users
	Tweets      map[int][]Tweet // UserID -> Tweets they've made
	IsFollowing map[int]Set     // UserID -> set of Users they follow ; call with map[follower][followee] ?
}

type MaxHeap []Tweet

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].Time > h[j].Time }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Tweet)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
