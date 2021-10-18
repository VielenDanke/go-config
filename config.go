package config

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"os"
	"reflect"
	"strconv"

	"gopkg.in/yaml.v2"
)

type FileType int

const (
	envTag = "goenv"
)

const (
	JSON FileType = iota
	YAML
	XML
)

type configOption func(cfg interface{}) error

// WithParsingBytes initialize option with passing bytes for unmarshalling based on fileType
func WithParsingBytes(data []byte, fileType FileType) configOption {
	return func(cfg interface{}) error {
		return ParseBytes(data, fileType, cfg)
	}
}

// WithParsingReader initialize option with passing io.Reader for unmarshalling based on fileType
func WithParsingReader(reader io.Reader, fileType FileType) configOption {
	return func(cfg interface{}) error {
		return ParseReader(reader, fileType, cfg)
	}
}

// WithParsingFile initialize option path to config file for it's opening and unmarshalling based on fileType
func WithParsingFile(filePath string, fileType FileType) configOption {
	return func(cfg interface{}) error {
		return ParseFile(filePath, fileType, cfg)
	}
}

func WithParsingEnv() configOption {
	return func(cfg interface{}) error {
		return ParseEnv(cfg)
	}
}

// NewConfig initializing cfg struct with various of options
func NewConfig(cfg interface{}, opts ...configOption) (err error) {
	for _, v := range opts {
		err = v(cfg)
		if err != nil {
			return
		}
	}
	return
}

/*
ParseEnv method takes config interface as argument
runs over the fields of incoming interface, if field contains tag goenv:"value"
after it is searching for "value" in os Environment
if find - inject os environment into field, if not - do nothing

Important note: cfg and all inner struct field should be initialized as pointers

Example:
type Test struct {
   Inner *InnerTest
}

type InnerTest struct {
   Field string `goenv:"field"`
}

cfg := &Test{Inner: &InnerTest{}}
*/
func ParseEnv(cfg interface{}) error {
	if cfg == nil {
		return nil
	}
	v := reflect.ValueOf(cfg)
	if v.Kind() == reflect.Ptr {
		el := v.Elem()
		for i := 0; i < el.NumField(); i++ {
			if el.Field(i).Kind() == reflect.Ptr {
				err := ParseEnv(el.Field(i).Interface())
				if err != nil {
					return err
				}
			} else {
				t := reflect.TypeOf(cfg).Elem()
				tagenv := t.Field(i).Tag.Get(envTag)
				env := os.Getenv(tagenv)
				if env == "" {
					continue
				}
				switch el.Field(i).Kind() {
				case reflect.String:
					el.Field(i).SetString(env)
				case reflect.Int:
					num, _ := strconv.Atoi(env)
					el.Field(i).SetInt(int64(num))
				case reflect.Bool:
					b, _ := strconv.ParseBool(env)
					el.Field(i).SetBool(b)
				default:
					return errors.New("not implemented go type")
				}
			}
		}
	}
	return nil
}

// ParseBytes parse input slice of bytes to cfg interface{} based on fileType (YAML, JSON, XML)
// cfg should be passed as pointer
func ParseBytes(data []byte, fileType FileType, cfg interface{}) (err error) {
	switch fileType {
	case JSON:
		err = unmarshallJSON(data, cfg)
	case YAML:
		err = unmarshallYAML(data, cfg)
	case XML:
		err = unmarshallXML(data, cfg)
	default:
		err = errors.New("unknown file type")
	}
	return err
}

// ParseReader parse input io.Reader to cfg interface{} based on fileType (YAML, JSON, XML)
// Underneath using ParseBytes function
// cfg should be passed as pointer
func ParseReader(reader io.Reader, fileType FileType, cfg interface{}) (err error) {
	if data, readErr := io.ReadAll(bufio.NewReader(reader)); readErr == nil {
		return ParseBytes(data, fileType, cfg)
	} else {
		return readErr
	}
}

// ParseFile parse file based on filePath and fileType (YAML, JSON, XML). If file is not exists - returns an error
// Underneath using ParseReader function
// cfg should be passed as pointer
func ParseFile(filePath string, fileType FileType, cfg interface{}) (err error) {
	f, fErr := os.Open(filePath)

	if fErr != nil {
		err = fErr
		return
	}
	defer CloseResources(f)

	return ParseReader(f, fileType, cfg)
}

func unmarshallJSON(data []byte, cfg interface{}) error {
	return json.Unmarshal(data, cfg)
}

func unmarshallXML(data []byte, cfg interface{}) error {
	return xml.Unmarshal(data, cfg)
}

func unmarshallYAML(data []byte, cfg interface{}) error {
	return yaml.Unmarshal(data, cfg)
}
