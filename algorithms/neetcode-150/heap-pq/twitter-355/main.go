package main

import (
	"container/heap"
	"fmt"
)

func Constructor() Twitter {
	X := Twitter{
		Time:        0,
		Users:       make(Set),
		Tweets:      make(map[int][]Tweet),
		IsFollowing: make(map[int]Set),
	}
	return X
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	n := len(this.Tweets[userId])
	this.Tweets[userId] = append(
		this.Tweets[userId],
		Tweet{
			Time:   this.Time,
			ID:     tweetId,
			UserID: userId,
			TwInd:  n,
		},
	)
	this.Time++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	if this.IsFollowing[userId] == nil {
		this.IsFollowing[userId] = make(Set)
	}
	out := make([]int, 0)
	h := &MaxHeap{}
	this.IsFollowing[userId][userId] = true
	following := this.IsFollowing[userId]
	for fID, v := range following {
		if v {
			theirTweets := this.Tweets[fID]
			if len(theirTweets) > 0 {
				lastTweet := theirTweets[len(theirTweets)-1]
				heap.Push(h, lastTweet)
			}
		}
	}
	for h.Len() > 0 && len(out) < 10 {
		tw := heap.Pop(h).(Tweet)
		out = append(out, tw.ID)
		if tw.TwInd > 0 { // go back into their old Tweets
			tw2 := this.Tweets[tw.UserID][tw.TwInd-1]
			heap.Push(h, tw2)
		}
	}
	return out
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.IsFollowing[followerId] == nil {
		this.IsFollowing[followerId] = make(Set)
	}
	this.IsFollowing[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.IsFollowing[followerId] == nil {
		return
	}
	this.IsFollowing[followerId][followeeId] = false // just in case... I am a very superstitious fellow
	delete(
		this.IsFollowing[followerId],
		followeeId,
	)
}

func main() {
	//[null,null,null,[3,5]]
	obj := Constructor()
	obj.PostTweet(1, 5)
	obj.PostTweet(1, 3)
	param_2 := obj.GetNewsFeed(1)
	fmt.Println(param_2)

}
