package config

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type FileType int

const (
	JSON FileType = iota
	YAML
	XML
)

func unmarshallJSON(data []byte, cfg interface{}) error {
	return json.Unmarshal(data, cfg)
}

func unmarshallXML(data []byte, cfg interface{}) error {
	return xml.Unmarshal(data, cfg)
}

func unmarshallYAML(data []byte, cfg interface{}) error {
	return yaml.Unmarshal(data, &cfg)
}

// ParseEnv parse environment. Searchin for tag 'env' in structure. Also if field contains tag default - it's using it as a value for field
func ParseEnv(cfg interface{}) error {
	return nil
}

// ParseBytes parse input slice of bytes to cfg interface{} based on fileType (YAML, JSON, XML)
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
func ParseReader(reader io.Reader, fileType FileType, cfg interface{}) (err error) {
	if data, readErr := io.ReadAll(bufio.NewReader(reader)); readErr == nil {
		return ParseBytes(data, fileType, cfg)
	} else {
		return readErr
	}
}

// ParseFile parse file based on filePath and fileType (YAML, JSON, XML). If file is not exists - returns an error
// Underneath using ParseReader function
func ParseFile(filePath string, fileType FileType, cfg interface{}) (err error) {
	f, fErr := os.Open(filePath)

	if fErr != nil {
		err = fErr
		return
	}
	defer CloseResources(f)

	return ParseReader(f, fileType, cfg)
}
