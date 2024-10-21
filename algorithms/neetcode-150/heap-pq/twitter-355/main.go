package main

type Set map[int]bool

type Tweets []int

type Twitter struct {
	Users       Set            // Set of Users
	Tweets      map[int]Tweets // UserID -> Tweets they've made
	IsFollowing Set            // UserID -> set of Users they follow
}

func Constructor() Twitter {
	X := Twitter{}
	return X
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Tweets[userId] = append(this.Tweets[userId], tweetId)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	return []int{69}
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	this.IsFollowing[followerId] = true

}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.IsFollowing[followerId] = false
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
