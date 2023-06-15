package indextest

import (
	"app/cmd/utils/testtool"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	router, w, _ := testtool.NewTestServerWithRouter()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"ok\"}", w.Body.String())
}
