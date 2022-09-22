package main

import (
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
