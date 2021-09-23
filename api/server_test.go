package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestServer(t *testing.T) {
	suite.Run(t, new(DiscordTestSuite))
}

func (suite *DiscordTestSuite) TestServerDiscordMessageError() {

	ts := httptest.NewServer(SetupServer())

	// Shut down the server and block until all requests have gone through
	defer ts.Close()
	jsonData, _ := json.Marshal(suite.testPayload)
	resp, err := http.Post(fmt.Sprintf("%s/v1/discord-message", ts.URL),
		"application/json",
		bytes.NewReader(jsonData))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, resp.StatusCode)
}

func (suite *DiscordTestSuite) TestServerHealthCheck() {

	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/healthz", ts.URL))
	var jsonData map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&jsonData)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)
	assert.Equal(suite.T(), "UP", jsonData["status"])
}
