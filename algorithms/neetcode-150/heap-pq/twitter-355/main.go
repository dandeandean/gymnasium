package main

type Set map[int]bool

type Tweet struct {
	Time int
	ID   int
}

type Twitter struct {
	Time        int
	Users       Set             // Set of Users
	Tweets      map[int][]Tweet // UserID -> Tweets they've made
	IsFollowing map[int]Set     // UserID -> set of Users they follow ; call with map[follower][followee] ?
}

func Constructor() Twitter {
	X := Twitter{}
	return X
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Tweets[userId] = append(
		this.Tweets[userId],
		Tweet{Time: this.Time, ID: tweetId},
	)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	return []int{69}
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

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
