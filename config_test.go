package config_test

import (
	"encoding/xml"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vielendanke/go-config"
)

type TestConfig struct {
	XMLName    xml.Name         `xml:"TestConfig"`
	First      string           `json:"first" env:"first" yaml:"first" xml:"First" goenv:"FIRST"`
	Second     int              `json:"second" env:"second" yaml:"second" xml:"Second" goenv:"SECOND"`
	InnerThird *InnerTestConfig `json:"inner_third" yaml:"innerThird" xml:"InnerTestConfig"`
}

type InnerTestConfig struct {
	XMLName    xml.Name `xml:"InnerTestConfig"`
	FirstInner string   `json:"first_inner" env:"first_inner" yaml:"firstInner" xml:"FirstInner" goenv:"FIRST_INNER"`
}

func TestNewConfigWithFileNOpt_Success(t *testing.T) {
	// prepare
	cfgForParse := &TestConfig{
		InnerThird: &InnerTestConfig{},
	}

	// make test
	err := config.NewConfig(cfgForParse, config.WithParsingFile("test.json", config.JSON))

	// assertions
	assert.Nil(t, err)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestNewConfigWithFileNOpt_EnvOpt_Success(t *testing.T) {
	// prepare
	cfgForParse := &TestConfig{
		InnerThird: &InnerTestConfig{},
	}
	data := []byte(`
	{
		"first":"first",
		"inner_third": {
			"first_inner":"first_inner"
		}
	}
	`)
	os.Setenv("SECOND", "2")
	defer func() {
		os.Remove("SECOND")
	}()

	// make test
	err := config.NewConfig(cfgForParse, config.WithParsingBytes(data, config.JSON), config.WithParsingEnv())

	// assertions
	assert.Nil(t, err)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestParseEnv_Success(t *testing.T) {
	// prepare
	os.Setenv("FIRST", "first")
	os.Setenv("SECOND", "2")
	os.Setenv("FIRST_INNER", "first_inner")

	defer func() {
		os.Remove("FIRST")
		os.Remove("SECOND")
		os.Remove("FIRST_INNER")
	}()

	cfgForParse := &TestConfig{
		InnerThird: &InnerTestConfig{},
	}

	// make test
	err := config.ParseEnv(cfgForParse)

	// assertions
	assert.Nil(t, err)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestNewConfigWithBytesOpt_Success(t *testing.T) {
	// prepare
	cfgForParse := &TestConfig{
		InnerThird: &InnerTestConfig{},
	}
	f, _ := os.Open("test.json")
	data, _ := io.ReadAll(f)

	// make test
	err := config.NewConfig(cfgForParse, config.WithParsingBytes(data, config.JSON))

	// assertions
	assert.Nil(t, err)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestNewConfigWithReaderOpt_Success(t *testing.T) {
	// prepare
	cfgForParse := &TestConfig{
		InnerThird: &InnerTestConfig{},
	}
	f, _ := os.Open("test.json")

	// make test
	err := config.NewConfig(cfgForParse, config.WithParsingReader(f, config.JSON))

	// assertions
	assert.Nil(t, err)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestParseBytesJSON_Success(t *testing.T) {
	// prepare
	f, _ := os.Open("test.json")

	data, _ := io.ReadAll(f)

	cfgForParse := &TestConfig{
		InnerThird: &InnerTestConfig{},
	}

	// make test
	resErr := config.ParseBytes(data, config.JSON, cfgForParse)

	// assertions
	assert.Nil(t, resErr)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestParseReaderJSON_Success(t *testing.T) {
	// prepare
	f, _ := os.Open("test.json")

	cfgForParse := &TestConfig{
		InnerThird: &InnerTestConfig{},
	}

	// make test
	resErr := config.ParseReader(f, config.JSON, cfgForParse)

	// assertions
	assert.Nil(t, resErr)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestParseFileJSON_Success(t *testing.T) {
	// prepare
	cfgForParse := &TestConfig{
		InnerThird: &InnerTestConfig{},
	}

	// make test
	resErr := config.ParseFile("test.json", config.JSON, cfgForParse)

	// assertions
	assert.Nil(t, resErr)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestParseBytesYAML_Success(t *testing.T) {
	// prepare
	f, _ := os.Open("test.yaml")

	data, _ := io.ReadAll(f)

	cfgForParse := &TestConfig{
		InnerThird: &InnerTestConfig{},
	}

	// make test
	resErr := config.ParseBytes(data, config.YAML, cfgForParse)

	// assertions
	assert.Nil(t, resErr)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestParseReaderYAML_Success(t *testing.T) {
	// prepare
	f, _ := os.Open("test.yaml")

	cfgForParse := &TestConfig{
		InnerThird: &InnerTestConfig{},
	}

	// make test
	resErr := config.ParseReader(f, config.YAML, cfgForParse)

	// assertions
	assert.Nil(t, resErr)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestParseFileYAML_Success(t *testing.T) {
	// prepare
	cfgForParse := &TestConfig{
		InnerThird: &InnerTestConfig{},
	}

	// make test
	resErr := config.ParseFile("test.yaml", config.YAML, cfgForParse)

	// assertions
	assert.Nil(t, resErr)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestParseBytesXML_Success(t *testing.T) {
	// prepare
	f, _ := os.Open("test.xml")

	data, _ := io.ReadAll(f)

	cfgForParse := &TestConfig{
		First:      "f",
		Second:     2,
		InnerThird: &InnerTestConfig{FirstInner: "asd"},
	}

	// make test
	resErr := config.ParseBytes(data, config.XML, cfgForParse)

	// assertions
	assert.Nil(t, resErr)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestParseReaderXML_Success(t *testing.T) {
	// prepare
	f, _ := os.Open("test.xml")

	cfgForParse := &TestConfig{
		First:      "f",
		Second:     2,
		InnerThird: &InnerTestConfig{FirstInner: "asd"},
	}

	// make test
	resErr := config.ParseReader(f, config.XML, cfgForParse)

	// assertions
	assert.Nil(t, resErr)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}

func TestParseFileXML_Success(t *testing.T) {
	// prepare
	cfgForParse := &TestConfig{
		First:      "f",
		Second:     2,
		InnerThird: &InnerTestConfig{FirstInner: "asd"},
	}

	// make test
	resErr := config.ParseFile("test.xml", config.XML, cfgForParse)

	// assertions
	assert.Nil(t, resErr)
	assert.Equal(t, "first", cfgForParse.First)
	assert.Equal(t, 2, cfgForParse.Second)
	assert.Equal(t, "first_inner", cfgForParse.InnerThird.FirstInner)
}
