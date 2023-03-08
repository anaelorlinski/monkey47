package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvOrStringNoEnv(t *testing.T) {
	flag := EnvOrString("monkey-env-test1", "test1")
	assert.Equal(t, "test1", flag)
}

func TestEnvOrStrHaveEnv(t *testing.T) {
	os.Setenv("monkey-env-test1", "test2")
	defer os.Unsetenv("monkey-env-test1")

	flag := EnvOrString("monkey-env-test1", "test1")
	assert.Equal(t, "test2", flag)
}
