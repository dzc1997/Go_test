package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloTom(t *testing.T) {
	//测试HelloTom()函数
	output := HelloTom()
	//期望输出"Tom"
	expectOutput := "Tom"
	assert.Equal(t, expectOutput, output)
}
