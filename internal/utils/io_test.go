package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHTTPURLIsTrueForHTTPS(t *testing.T) {
	sut := "https://test.com"
	assert.True(t, IsHTTPURL(sut))
}

func TestIsHTTPURLIsFalseForHTTP(t *testing.T) {
	sut := "http://test.com"
	assert.False(t, IsHTTPURL(sut))
}

func TestIsHTTPURLIsFalseForFileURL(t *testing.T) {
	sut := "./go.mod"
	assert.False(t, IsHTTPURL(sut))
}
