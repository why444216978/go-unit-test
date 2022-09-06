package main

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	convey.Convey("TestSend", t, func() {
		convey.Convey("success", func() {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder(http.MethodGet, "https://127.0.0.1:8080", httpmock.NewStringResponder(http.StatusOK, ""))

			err := Send()
			assert.Nil(t, err)
		})
	})
}
