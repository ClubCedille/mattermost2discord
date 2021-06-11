package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoggerIsSameInstance(T *testing.T) {
	logger := loggerInstance()
	logger.Info("Testing logger")

  if assert.NotNil(T, logger) {

    // now we know that object isn't nil, we are safe to make
    // further assertions without causing any errors
    assert.Equal(T, loggerInstance(), logger)
	}
}
