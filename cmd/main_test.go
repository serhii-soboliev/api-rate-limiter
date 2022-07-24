package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	rt "com.sbk/api-rate-limiter/api/v1/router"

	"github.com/stretchr/testify/assert"
)

func TestFixedWindowShouldBanFourthRequest(t *testing.T) {
	r := rt.InitalizeRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/resource/fixedwindow/123", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
    r.ServeHTTP(w, req)
    assert.Equal(t, 200, w.Code)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, 403, w.Code)
    assert.Equal(t, "\"Too much reqeuests for 123\"", w.Body.String())
}
