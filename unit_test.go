package main

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	a := 1
	b := 2
	want := -1

	actual := Compare(a, b)
	if actual != want {
		t.Errorf("want %d, actual %d", want, actual)
	}
}

func TestCompareTDT(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "a > b", args: args{2, 1}, want: 1},
		{name: "a == b", args: args{2, 2}, want: 0},
		{name: "a < b", args: args{1, 2}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareConvey(t *testing.T) {
	convey.Convey("TestCompareConvey", t, func() {
		convey.Convey("a > b", func() {
			assert.Equal(t, 1, Compare(2, 1))
		})
		convey.Convey("a == b", func() {
			assert.Equal(t, 0, Compare(2, 2))
		})
		convey.Convey("a < b", func() {
			assert.Equal(t, -1, Compare(1, 2))
		})
	})
}
