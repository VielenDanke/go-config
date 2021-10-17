package config_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vielendanke/go-config"
)

type TestConfig struct {
	First      string           `json:"first" env:"first"`
	Second     int              `json:"second" env:"second"`
	InnerThird *InnerTestConfig `json:"inner_third"`
}

type InnerTestConfig struct {
	FirstInner string `json:"first_inner" env:"first_inner"`
}

func TestParseBytes_Success(t *testing.T) {
	// prepare
	cfg := &TestConfig{
		First:  "first",
		Second: 2,
		InnerThird: &InnerTestConfig{
			FirstInner: "Inner test field",
		},
	}
	data, _ := json.Marshal(cfg)

	cfgForParse := &TestConfig{
		InnerThird: &InnerTestConfig{},
	}

	// make test
	resErr := config.ParseBytes(data, config.JSON, cfgForParse)

	// assertions
	assert.Nil(t, resErr)
	assert.Equal(t, cfgForParse.First, cfg.First)
	assert.Equal(t, cfg.Second, cfgForParse.Second)
	assert.Equal(t, cfg.InnerThird.FirstInner, cfgForParse.InnerThird.FirstInner)
}
