package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfigVaultClient(t *testing.T) {
	// prepare
	setEnvErr := os.Setenv("VAULT_ADDR", "http://localhost:8200")

	// make test
	cli, cliErr := configVaultClient()

	// clean
	rErr := os.Unsetenv("VAULT_ADDR")

	// assertions
	assert.NotNil(t, cli)
	assert.Nil(t, setEnvErr)
	assert.Nil(t, cliErr)
	assert.Nil(t, rErr)
}

func TestConfigVaultClient_Fails_EnvDoesNotExists(t *testing.T) {
	// make test
	cli, cliErr := configVaultClient()

	// assertions
	assert.Nil(t, cli)
	assert.NotNil(t, cliErr)
}