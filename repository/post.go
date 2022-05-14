package repository

import (
	"errors"
	"sync"
	"time"
)

type Post struct {
	Id         int    `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}
func (*PostDao) QueryPostsByParentId(parentId int64) []*Post {
	postRwLock.RLock()
	defer postRwLock.RUnlock()
	return postIndexMap[parentId]
}

func (*PostDao) QueryPostsIdLen(ParentId int64) int {
	return len(postIndexMap[ParentId])
}

func (*PostDao) CreateNewPost(ParentId int64, Id int, Content string, CreateTime int64) (Post, error) {
	post := Post{Id: Id, ParentId: ParentId, Content: Content, CreateTime: CreateTime}

	return post, nil
}

func (*PostDao) UpdatePostsByParentId(ParentId int64, NewPost *Post) error {
	postRwLock.Lock()
	defer postRwLock.Unlock()
	if NewPost.ParentId != ParentId {
		return errors.New("unmatch")
	}
	id := num_post + 1
	num_post = num_post + 1
	NewPost.Id = id
	NewPost.CreateTime = time.Now().Unix()
	posts, ok := postIndexMap[ParentId]
	if !ok {
		posts = []*Post{}
	}
	posts = append(posts, NewPost)
	postIndexMap[ParentId] = posts
	return nil
}
