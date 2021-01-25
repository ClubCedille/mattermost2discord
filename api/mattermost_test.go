package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MattermostFunc(t *testing.T) {
	ts := httptest.NewServer(SetupServer())

	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	resp, err := http.Get(fmt.Sprintf("%s/", ts.URL))

	// Make test assertions
	assert.Nil(t, err)
	assert.True(t, resp.StatusCode == http.StatusOK)
}
