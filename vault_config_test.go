package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	testPath = "/secret/data/test"
	testToken = "root"
	testHost = "http://127.0.0.1:8200"
)

const (
	vaultAddr = "VAULT_ADDR"
	vaultSecretPath = "VAULT_SECRET_PATH"
	vaultToken = "VAULT_TOKEN"
)

func TestConfigVaultClient(t *testing.T) {
	// prepare
	setEnvErr := os.Setenv(vaultAddr, testHost)

	// make test
	cli, cliErr := configVaultClient()

	// clean
	rErr := os.Unsetenv(vaultAddr)

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

func TestFetchVaultSecret(t *testing.T) {
	// prepare
	setEnvErr := os.Setenv(vaultAddr, testHost)

	// make test
	cfg, cfgFetchErr := FetchVaultSecret(testPath, testToken)

	// assertions
	assert.Nil(t, cfgFetchErr)
	assert.NotNil(t, cfg)
	assert.Nil(t, setEnvErr)
}

func TestFetchVaultSecret_TokenIsEmpty(t *testing.T) {
	// prepare
	setEnvErr := os.Setenv(vaultAddr, testHost)

	// make test
	cfg, cfgFetchErr := FetchVaultSecret(testPath, "")

	// assertions
	assert.Nil(t, cfg)
	assert.NotNil(t, cfgFetchErr)
	assert.Nil(t, setEnvErr)
}

func TestFetchVaultSecret_PathIsEmpty(t *testing.T) {
	// prepare
	setEnvErr := os.Setenv(vaultAddr, testHost)

	// make test
	cfg, cfgFetchErr := FetchVaultSecret(testPath, "")

	// assertions
	assert.Nil(t, cfg)
	assert.NotNil(t, cfgFetchErr)
	assert.Nil(t, setEnvErr)
}

func TestFetchVaultSecretEnv(t *testing.T) {
	// prepare
	setEnvErr := os.Setenv(vaultAddr, testHost)
	setEnvErr = os.Setenv(vaultToken, testToken)
	setEnvErr = os.Setenv(vaultSecretPath, testPath)

	// make test
	secret, fetchErr := FetchVaultSecretEnv()

	// clean
	setEnvErr = os.Unsetenv(vaultAddr)
	setEnvErr = os.Unsetenv(vaultToken)
	setEnvErr = os.Unsetenv(vaultSecretPath)

	// assertions
	assert.Nil(t, fetchErr)
	assert.NotNil(t, secret)
	assert.Nil(t, setEnvErr)
}

func TestFetchVaultSecretEnv_VaultAddrIsEmpty(t *testing.T) {
	// prepare
	setEnvErr := os.Setenv(vaultToken, testToken)
	setEnvErr = os.Setenv(vaultSecretPath, testPath)

	// make test
	secret, fetchErr := FetchVaultSecretEnv()

	// clean
	setEnvErr = os.Unsetenv(vaultToken)
	setEnvErr = os.Unsetenv(vaultSecretPath)

	// assertions
	assert.NotNil(t, fetchErr)
	assert.Nil(t, secret)
	assert.Nil(t, setEnvErr)
}

func TestFetchVaultSecretEnv_VaultTokenIsEmpty(t *testing.T) {
	// prepare
	setEnvErr := os.Setenv(vaultAddr, testHost)
	setEnvErr = os.Setenv(vaultSecretPath, testPath)

	// make test
	secret, fetchErr := FetchVaultSecretEnv()

	// clean
	setEnvErr = os.Unsetenv(vaultAddr)
	setEnvErr = os.Unsetenv(vaultSecretPath)

	// assertions
	assert.NotNil(t, fetchErr)
	assert.Nil(t, secret)
	assert.Nil(t, setEnvErr)
}

func TestFetchBytesVaultSecretData(t *testing.T) {
	// prepare
	envErr := os.Setenv(vaultAddr, testHost)

	// make test
	data, err := FetchBytesVaultSecretData(testPath, testToken)

	// clean
	envErr = os.Unsetenv(vaultAddr)

	// assertions
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Nil(t, envErr)
}

func TestFetchBytesVaultSecretData_VaultAddrEnvIsEmpty(t *testing.T) {
	// make test
	data, err := FetchBytesVaultSecretData(testPath, testToken)

	// assertions
	assert.NotNil(t, err)
	assert.Nil(t, data)
}

func TestFetchBytesVaultSecretData_PathIsEmpty(t *testing.T) {
	// make test
	data, err := FetchBytesVaultSecretData("", testToken)

	// assertions
	assert.NotNil(t, err)
	assert.Nil(t, data)
}

func TestFetchBytesVaultSecretData_TokenIsEmpty(t *testing.T) {
	// make test
	data, err := FetchBytesVaultSecretData(testPath, "")

	// assertions
	assert.NotNil(t, err)
	assert.Nil(t, data)
}

func TestFetchBytesVaultSecretDataEnv(t *testing.T) {
	// prepare
	setEnvErr := os.Setenv(vaultAddr, testHost)
	setEnvErr = os.Setenv(vaultToken, testToken)
	setEnvErr = os.Setenv(vaultSecretPath, testPath)

	// make test
	secret, fetchErr := FetchBytesVaultSecretDataEnv()

	// clean
	setEnvErr = os.Unsetenv(vaultAddr)
	setEnvErr = os.Unsetenv(vaultToken)
	setEnvErr = os.Unsetenv(vaultSecretPath)

	// assertions
	assert.Nil(t, setEnvErr)
	assert.Nil(t, fetchErr)
	assert.NotNil(t, secret)
}

func TestFetchBytesVaultSecretDataEnv_VaultAddrEnvIsEmpty(t *testing.T) {
	// prepare
	setEnvErr := os.Setenv(vaultToken, testToken)
	setEnvErr = os.Setenv(vaultSecretPath, testPath)

	// make test
	secret, fetchErr := FetchBytesVaultSecretDataEnv()

	// clean
	setEnvErr = os.Unsetenv(vaultToken)
	setEnvErr = os.Unsetenv(vaultSecretPath)

	// assertions
	assert.Nil(t, setEnvErr)
	assert.NotNil(t, fetchErr)
	assert.Nil(t, secret)
}

func TestFetchBytesVaultSecretDataEnv_VaultTokenEnvIsEmpty(t *testing.T) {
	// prepare
	setEnvErr := os.Setenv(vaultAddr, testHost)
	setEnvErr = os.Setenv(vaultSecretPath, testPath)

	// make test
	secret, fetchErr := FetchBytesVaultSecretDataEnv()

	// clean
	setEnvErr = os.Unsetenv(vaultAddr)
	setEnvErr = os.Unsetenv(vaultSecretPath)

	// assertions
	assert.Nil(t, setEnvErr)
	assert.NotNil(t, fetchErr)
	assert.Nil(t, secret)
}

func TestFetchBytesVaultSecretDataEnv_VaultSecretPathEnvIsEmpty(t *testing.T) {
	// prepare
	setEnvErr := os.Setenv(vaultAddr, testHost)
	setEnvErr = os.Setenv(vaultToken, testToken)

	// make test
	secret, fetchErr := FetchBytesVaultSecretDataEnv()

	// clean
	setEnvErr = os.Unsetenv(vaultAddr)
	setEnvErr = os.Unsetenv(vaultToken)

	// assertions
	assert.Nil(t, setEnvErr)
	assert.NotNil(t, fetchErr)
	assert.Nil(t, secret)
}