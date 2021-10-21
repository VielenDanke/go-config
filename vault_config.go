package config

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/vault/api"
)

/*
configVaultClient function create default api.Client for Vault
Important note: address is fetching from Environment using os.Getenv function
Environment key - VAULT_ADDR
Returning api.Client, error
*/
func configVaultClient() (*api.Client, error) {
	vaultEnv := os.Getenv("VAULT_ADDR")
	if vaultEnv == "" {
		return nil, errors.New("VAULT_ADDR Environment variable is required")
	}
	cfg := &api.Config{
		Address: vaultEnv,
		HttpClient: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
		MinRetryWait: time.Millisecond * 1000,
		MaxRetryWait: time.Millisecond * 1500,
		MaxRetries:   2,
		Timeout:      time.Second * 60,
		Backoff:      retryablehttp.LinearJitterBackoff,
		CheckRetry:   retryablehttp.DefaultRetryPolicy,
	}
	return api.NewClient(cfg)
}

/*
fetchVaultConfig receives api.Client - creating by configVaultClient function
path - to Vault secret directory (pattern for Vault v2 - /secret/data/{secret_name})
token - for accessing Vault
If fetched api.Secret is nil - return error
if read from Vault failed - return error
Returning: map[string]interface{}, error
*/
func fetchVaultConfig(cli *api.Client, path, token string) (*api.Secret, error) {
	cli.SetToken(token)
	data, dataErr := cli.Logical().Read(path)
	if dataErr != nil {
		return nil, dataErr
	}
	if data == nil {
		return nil, fmt.Errorf("path %s not found", path)
	}
	return data, nil
}

/*
FetchVaultSecret process fetching secret using incoming arguments (path, token)
configVaultClient and fetchVaultConfig functions are using inside
Returning: map[string]interface{}, error
*/
func FetchVaultSecret(path, token string) (*api.Secret, error) {
	cli, configErr := configVaultClient()
	if configErr != nil {
		return nil, configErr
	}
	return fetchVaultConfig(cli, path, token)
}

/*
FetchVaultSecretEnv process fetching secret from vault using env variable VAULT_TOKEN for token and VAULT_SECRET_PATH for path to secret
FetchVaultSecret function is using inside
Returning: *api.Secret, error
*/
func FetchVaultSecretEnv() (*api.Secret, error) {
	path, token, err := fetchVaultEnv()
	if err != nil {
		return nil, err
	}
	return FetchVaultSecret(path, token)
}

/*
FetchBytesVaultSecretData process fetching secret bytes using incoming arguments (path, token)
FetchVaultSecret function is using inside
Returning: []byte, error
*/
func FetchBytesVaultSecretData(path, token string) ([]byte, error) {
	secret, err := FetchVaultSecret(path, token)
	if err != nil {
		return nil, err
	}
	data, mErr := json.Marshal(secret.Data["data"])
	if mErr != nil {
		return nil, mErr
	}
	return data, nil
}

/*
FetchBytesVaultSecretDataEnv process fetching secret bytes env variable VAULT_TOKEN for token and VAULT_SECRET_PATH for path to secret
FetchBytesVaultSecretData function is using inside
Returning: []byte, error
*/
func FetchBytesVaultSecretDataEnv() ([]byte, error) {
	path, token, err := fetchVaultEnv()
	if err != nil {
		return nil, err
	}
	return FetchBytesVaultSecretData(path, token)
}

func fetchVaultEnv() (path, token string, err error) {
	token = os.Getenv("VAULT_TOKEN")
	path = os.Getenv("VAULT_SECRET_PATH")

	if len(token) == 0 {
		err = errors.New("token is not set in the environment variable VAULT_TOKEN")
		return
	}
	if len(path) == 0 {
		err = errors.New("path is not set in the enviromnent variable VAULT_SECRET_PATH")
		return
	}
	return
}