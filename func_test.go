package main

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

// TestA 函数，单次打桩
func TestA(t *testing.T) {
	patch := gomonkey.ApplyFunc(B, func() int {
		return 1
	})
	defer patch.Reset()

	assert.Equal(t, 1, A())
}

// TestAA 函数，连续打桩
func TestAA(t *testing.T) {
	patch := gomonkey.ApplyFuncSeq(B, []gomonkey.OutputCell{
		{Values: gomonkey.Params{1}},
		{Values: gomonkey.Params{2}},
	})
	defer patch.Reset()

	assert.Equal(t, 3, AA())
}
