package service

import (
	"errors"
	"github.com/dzc1997/Go_test/repository"
	"sync"
)

// PageInfo 页面信息结构体
type PageInfo struct {
	Topic    *repository.Topic  //话题
	PostList []*repository.Post //帖子
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}

func NewQueryPageInfoFlow(topId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topId,
	}
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo

	topic *repository.Topic
	posts []*repository.Post
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	//校验ID
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	//从repository层获取话题数据与回帖列表数据
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	//
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

//对ID进行校验
func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

//通过repository层的查询方法获取话题数据和回帖列表数据
//获取话题数据与获取回帖列表数据互不依赖，采用协程方式更快速
func (f *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		//获取topic信息
		topic := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		f.topic = topic
	}()
	go func() {
		defer wg.Done()
		//获取post列表
		posts := repository.NewPostDaoInstance().QueryPostsByParentId(f.topicId)
		f.posts = posts
	}()
	wg.Wait()
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	f.pageInfo = &PageInfo{
		Topic:    f.topic,
		PostList: f.posts,
	}
	return nil
}
