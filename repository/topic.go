package repository

import (
	"sync"
)

// Topic 话题结构体
type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}
type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	//保证只执行一次，类似单例模式
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

// QueryTopicById 通过话题ID查询话题
func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}
