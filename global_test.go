package main

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/prashantv/gostub"
)

func TestSizeStub(t *testing.T) {
	tests := []struct {
		name string
		want int
		f    func() *gostub.Stubs
	}{
		{name: "size > 10", want: 10, f: func() *gostub.Stubs {
			return gostub.Stub(&size, 11)
		}},
		{name: "size <= 10", want: 3, f: func() *gostub.Stubs {
			return gostub.Stub(&size, 3)
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stub := tt.f()
			if got := Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
			stub.Reset()
		})
	}
}

func TestSizeMonkey(t *testing.T) {
	tests := []struct {
		name string
		want int
		f    func() *gomonkey.Patches
	}{
		{name: "size > 10", want: 10, f: func() *gomonkey.Patches {
			return gomonkey.ApplyGlobalVar(&size, 11)
		}},
		{name: "size <= 10", want: 3, f: func() *gomonkey.Patches {
			return gomonkey.ApplyGlobalVar(&size, 3)
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stub := tt.f()
			if got := Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
			stub.Reset()
		})
	}
}
