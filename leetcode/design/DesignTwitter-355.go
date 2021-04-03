package design

import (
	"container/heap"
)

// 解题思路
// 用list作为set，加入唯一性判断，用作followed
// 用container包中heap作为基础，加入符合题目要求的最大堆MaxHeap
// 用MaxHeap作为基础，构建优先队列priQ，保存按照时间顺序存储的tweet，只需要单个用户的表头，而单个用户的表头为最新的tweet（即时间戳最大）
// 用链表作为基础，构建单个用户的tweet按时间排序的链表
// 由于用自带的time包，time.Now().Unix(),产生的时间过快，导致可能存在相同时间，因此改用自定义的timestamp int64 来作为时间大小
// 哈希表 + 链表 + 优先队列
//========================MaxHeap====================
type MaxHeap []*Tweet

func (f MaxHeap) Len() int {
	return len(f)
}

func (f MaxHeap) Less(i, j int) bool {
	return f[i].time > f[j].time
}

func (f MaxHeap) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *MaxHeap) Push(x interface{}) {
	e := x.(*Tweet)
	*f = append(*f, e)
	heap.Fix(f, f.Len()-1)
}

func (f *MaxHeap) Pop() interface{} {
	e := (*f)[0]
	*f = (*f)[1:len(*f)]
	heap.Fix(f, 0)
	return e
}

//========================User====================
type User struct {
	id       int
	followed []int // 关注人列表
	head     *Tweet
}

func NewUser(userId int) *User {
	u := User{
		id: userId,
	}
	u.follow(userId)
	return &u
}

// 关注某人
func (u *User) follow(userId int) {
	for _, u := range u.followed {
		if u == userId {
			return
		}
	}
	u.followed = append(u.followed, userId)
}

// 取关某人
func (u *User) unfollow(userId int) {
	if userId == u.id {
		return
	}
	for i, f := range u.followed {
		if f == userId {
			//删除id
			u.followed = append(u.followed[:i], u.followed[i+1:]...)
			break
		}
	}
}

// 发文
func (u *User) post(tweetID int) {
	t := NewTweet(tweetID)
	// 将新文加入表头
	t.next = u.head
	u.head = t
}
//========================Tweet====================
type Tweet struct {
	id   int
	time int64
	next *Tweet
}

var timestamp int64 = 0

func NewTweet(tweetID int) *Tweet {
	timestamp++
	return &Tweet{
		id:   tweetID,
		time: timestamp,
	}
}
//========================Twitter====================
type Twitter struct {
	userMap map[int]*User
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		userMap: make(map[int]*User),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if user, ok := this.userMap[userId]; !ok {
		u := NewUser(userId)
		u.post(tweetId)
		this.userMap[userId] = u
		return
	} else {
		user.post(tweetId)
	}
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	var res []int
	priQ := MaxHeap{}
	if user, ok := this.userMap[userId]; !ok {
		return nil
	} else {
		followers := user.followed
		// 利用链表和优先队列按照时间大小进行多路归并
		for _, f := range followers {
			if followerUser, ok := this.userMap[f]; !ok {
				//还没有创建f用户
				continue
			} else {
				tw := followerUser.head
				if tw == nil {
					continue
				}
				priQ.Push(tw)
			}
		}

		// 将发布的tweet推入优先队列
		for i := 0; i < 10 && priQ.Len() != 0; i++ {
			t := priQ.Pop().(*Tweet)
			res = append(res, t.id)
			if t.next != nil {
				priQ.Push(t.next)
			}
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if user, ok := this.userMap[followerId]; !ok {
		this.userMap[followerId] = NewUser(followerId)
		this.userMap[followerId].follow(followeeId)
		return
	} else {
		user.follow(followeeId)
	}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if user, ok := this.userMap[followerId]; !ok {
		return
	} else {
		user.unfollow(followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
