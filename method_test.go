package main

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

// TestS_AA 成员方法单个打桩
func TestS_A(t *testing.T) {
	s := &S{}

	// 公共成员方法
	patch := gomonkey.ApplyMethod(reflect.TypeOf(s), "B", func(_ *S) int {
		return 1
	})
	// 私有成员方法
	patch.ApplyPrivateMethod(reflect.TypeOf(s), "b", func(_ *S) int {
		return 2
	})
	defer patch.Reset()

	assert.Equal(t, 3, s.A())
}

// TestS_AA 成员方法连续打桩
func TestS_AA(t *testing.T) {
	s := &S{}

	// 私有成员方法
	patch := gomonkey.ApplyFuncSeq((*S).b, []gomonkey.OutputCell{
		{Values: gomonkey.Params{1}},
		{Values: gomonkey.Params{2}},
	})
	// 公共成员方法
	patch.ApplyMethodSeq(reflect.TypeOf(s), "B", []gomonkey.OutputCell{
		{Values: gomonkey.Params{1}},
		{Values: gomonkey.Params{2}},
	})
	defer patch.Reset()

	assert.Equal(t, 6, s.AA())
}
