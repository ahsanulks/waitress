package handler

import (
	"bufio"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type fakeResponse struct {
	*httptest.ResponseRecorder
}

func (crr *fakeResponse) CloseNotify() <-chan bool {
	return nil
}

func (crr *fakeResponse) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return nil, nil, nil
}

func (crr *fakeResponse) Pusher() http.Pusher {
	return nil
}

func (crr *fakeResponse) Size() int {
	return 0
}

func (crr *fakeResponse) Status() int {
	return 0
}

func (crr *fakeResponse) WriteHeaderNow() {}

func (crr *fakeResponse) Written() bool {
	return true
}

func Test_render(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		response string
	}{
		{
			name:     "message",
			data:     "this is ok",
			response: `{"message":"this is ok"}`,
		},
		{
			name:     "error",
			data:     errors.New("unexpected error"),
			response: `{"error":"unexpected error"}`,
		},
		{
			name:     "nil",
			data:     nil,
			response: ``,
		},
		{
			name: "struct",
			data: struct {
				ID int `json:"id"`
			}{ID: 123},
			response: `{"id":123}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				rr = httptest.NewRecorder()
				c  = &gin.Context{Writer: &fakeResponse{rr}}
			)

			render(c, tt.data, 200)
			if tt.response != "" {
				assert.JSONEq(t, tt.response, rr.Body.String())
			} else {
				assert.Equal(t, tt.response, rr.Body.String())
			}
		})
	}
}
