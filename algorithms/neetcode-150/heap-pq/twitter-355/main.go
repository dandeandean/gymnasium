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
	following := this.IsFollowing[userId]
	h := &MaxHeap{}
	for fID, v := range following {
		if v {
			theirTweets := this.Tweets[fID]
			lastTweet := theirTweets[len(theirTweets)-1]
			heap.Push(h, lastTweet)
		}
	}
	out := make([]int, 0)
	for h.Len() > 0 && len(out) < 10 {
		tw := heap.Pop(h).(Tweet)
		out = append(out, tw.ID)
		// go back into their old Tweets
		if tw.TwInd > 1 {
			tw2 := this.Tweets[tw.UserID][tw.TwInd-1]
			heap.Push(h, tw2)
		}
	}
	return out
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	this.IsFollowing[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.IsFollowing[followerId][followeeId] = false // just in case... I am a very superstitious fellow
	delete(
		this.IsFollowing[followerId],
		followeeId,
	)
}

func main() {
	obj := Constructor()
	obj.PostTweet(5, 0)
	obj.PostTweet(5, 1)
	obj.PostTweet(5, 2)
	param_2 := obj.GetNewsFeed(5)
	fmt.Println(obj, param_2)

}
