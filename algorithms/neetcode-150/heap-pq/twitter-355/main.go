package main

import "fmt"

type Set map[int]bool

type Tweet struct {
	Time   int
	ID     int
	UserID int
}

type Twitter struct {
	Time        int
	Users       Set             // Set of Users
	Tweets      map[int][]Tweet // UserID -> Tweets they've made
	IsFollowing map[int]Set     // UserID -> set of Users they follow ; call with map[follower][followee] ?
}

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
	this.Tweets[userId] = append(
		this.Tweets[userId],
		Tweet{Time: this.Time, ID: tweetId, UserID: userId},
	)
	this.Time++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	following := this.IsFollowing[userId]
	allTweets := make([]Tweet, 0)
	out := make([]int, 0)
	for k, v := range following {
		if v {
			allTweets = append(allTweets, this.Tweets[k]...) // ... means slice to slice or something
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
	/**
	 * Your Twitter object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.PostTweet(userId,tweetId);
	 * param_2 := obj.GetNewsFeed(userId);
	 * obj.Follow(followerId,followeeId);
	 * obj.Unfollow(followerId,followeeId);
	 */
	obj := Constructor()
	obj.PostTweet(0, 0)
	param_2 := obj.GetNewsFeed(0)
	fmt.Println(obj, param_2)

}
