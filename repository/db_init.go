package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

//给topic和post两个索引加锁保证并发安全
var (
	topicIndexMap map[int64]*Topic
	topicRwLock   sync.RWMutex
	postIndexMap  map[int64][]*Post
	postRwLock    sync.RWMutex
	num_topic     int
	num_post      int
)

func Init(filePath string) error {
	if err := initTopicIndexMap(filePath); err != nil {
		return err
	}
	if err := initPostIndexMap(filePath); err != nil {
		return err
	}
	return nil
}

func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	cnt := 0
	for scanner.Scan() {
		text := scanner.Text()
		cnt = cnt + 1
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	num_topic = cnt
	return nil
}

func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	cnt := 0
	for scanner.Scan() {
		cnt = cnt + 1
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		posts, ok := postTmpMap[post.ParentId]
		if !ok {
			postTmpMap[post.ParentId] = []*Post{&post}
			continue
		}
		posts = append(posts, &post)
		postTmpMap[post.ParentId] = posts
	}
	num_post = cnt
	postIndexMap = postTmpMap
	return nil
}

/*func UpdateFile(filePath string,post Post) error{
	buf, err:=json.Marshal(post)
	if err!=nil{
		return nil
	}
	open,err:=os.Open(filePath+"post",os.WRONLY)
	if err!=nil{
		return err
	}
	defer open.close()
	write:=bufio.NewWriter(open)
	write.WriteString(buf)
	return nil
}*/
