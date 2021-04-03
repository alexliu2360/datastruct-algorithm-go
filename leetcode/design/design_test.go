package design

import (
	"fmt"
	"testing"
)

//["Twitter","postTweet","getNewsFeed","follow","postTweet","getNewsFeed","unfollow","getNewsFeed"]
//[[]        ,[1,5],     [1],          [1,2],   [2,6],      [1]           ,[1,2],    [1]]
//[null,     null,       [5],           null,     null,      [5],          null,     [5]]
//[null,     null,       [5],           null       ,null,    [6,5],        null,     [5]]

func TestDesignTwitter(t *testing.T) {
	tw := Constructor()
	tw.PostTweet(1, 5)
	fmt.Println(tw.GetNewsFeed(1))
	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	fmt.Println(tw.GetNewsFeed(1))
	tw.Unfollow(1, 2)
	fmt.Println(tw.GetNewsFeed(1))
}
//["Twitter","postTweet","getNewsFeed","follow","getNewsFeed","unfollow","getNewsFeed"]
//[[],[1,1],[1],[2,1],[2],[2,1],[2]]
//[null,null,[1],null,[],null,[]]
//[null,null,[1],null,[1],null,[]]
func TestDesignTwitter2(t *testing.T) {
	tw := Constructor()
	tw.PostTweet(1, 1)
	fmt.Println(tw.GetNewsFeed(1))
	tw.Follow(2, 1)
	fmt.Println(tw.GetNewsFeed(2))
	tw.Unfollow(2, 1)
	fmt.Println(tw.GetNewsFeed(2))
}
//["Twitter","follow","getNewsFeed"]
//[[],[1,5],[1]]
func TestDesignTwitter3(t *testing.T) {
	tw := Constructor()
	tw.Follow(1, 5)
	fmt.Println(tw.GetNewsFeed(2))
}

//["Twitter","postTweet","follow","follow","getNewsFeed"]
//[[],[2,5],[1,2],[1,2],[1]]
//[null,null,null,null,[5,5]]
//[null,null,null,null,[5]]
func TestDesignTwitter4(t *testing.T) {
	tw := Constructor()
	tw.PostTweet(2, 5)
	tw.Follow(1, 2)
	tw.Follow(1, 2)
	fmt.Println(tw.GetNewsFeed(1))
}
