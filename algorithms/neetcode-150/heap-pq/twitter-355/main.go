package main

type Twitter struct {
}

func Constructor() Twitter {
	X := Twitter{}
	return X
}

func (this *Twitter) PostTweet(userId int, tweetId int) {

}

func (this *Twitter) GetNewsFeed(userId int) []int {
	return []int{69}
}

func (this *Twitter) Follow(followerId int, followeeId int) {

}

func (this *Twitter) Unfollow(followerId int, followeeId int) {

}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
