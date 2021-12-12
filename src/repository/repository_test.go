package Repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLatestVersion(t *testing.T) {
	filename := getLatestVersion("../../configuration")
	assert.Equal(t, "proxy.yaml", filename)
}
