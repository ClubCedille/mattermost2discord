package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerDiscordMessage(t *testing.T) {
	DiscordToken = "test"
	DiscordChannel = "test"
	TriggerWordMattermost = "2disc"
	ts := httptest.NewServer(SetupServer())

	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	payload := MattermostPayload{Text: "2disc hello", Username: "hello"}
	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post(fmt.Sprintf("%s/v1/discord-message/", ts.URL),
		"application/json",
		bytes.NewReader(jsonData))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestServerHealthCheck(t *testing.T) {
	ts := httptest.NewServer(SetupServer())

	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/health/", ts.URL))
	var jsonData map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&jsonData)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "UP", jsonData["status"])
}
