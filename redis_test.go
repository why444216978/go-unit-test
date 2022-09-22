package main

import (
	"errors"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_handleRedis(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	c := NewMockCmdable(ctl)
	c.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return(redis.NewStringResult("redis", nil))

	res, err := handleRedis(c)
	assert.Nil(t, err)
	assert.Equal(t, "redis", res)
}

func Test_handleOrder(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	c := NewMockCmdable(ctl)
	m1 := c.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return(redis.NewStringResult("redis", errors.New("error")))
	m2 := c.EXPECT().Get(gomock.Any(), gomock.Any()).Times(1).Return(redis.NewStringResult("redis", nil))
	gomock.InOrder(m1, m2)

	res, err := handleOrder(c)
	assert.Nil(t, err)
	assert.Equal(t, "redis", res)
}
