package service

import (
	"github.com/dzc1997/Go_test/repository"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	repository.Init("../data/")
	os.Exit(m.Run())
}
func TestQueryPageInfo(t *testing.T) {
	pageInfo, _ := QueryPageInfo(1)
	assert.NotEqual(t, nil, pageInfo)
	assert.Equal(t, 5, len(pageInfo.PostList))
}
