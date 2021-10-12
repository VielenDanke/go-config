package config

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type FileType int

const (
	JSON FileType = iota
	YAML
	XML
)

func unmarshallJSON(reader io.Reader, cfg interface{}) error {
	return json.NewDecoder(reader).Decode(&cfg)
}

func unmarshallXML(reader io.Reader, cfg interface{}) error {
	return xml.NewDecoder(reader).Decode(&cfg)
}

func unmarshallYAML(reader io.Reader, cfg interface{}) error {
	return yaml.NewDecoder(reader).Decode(&cfg)
}

func ParseFile(filePath string, fileType FileType, cfg interface{}) (err error) {
	f, fErr := os.Open(filePath)

	if fErr != nil {
		err = fErr
		return
	}
	defer CloseResources(f)

	switch fileType {
	case JSON:
		err = unmarshallJSON(f, cfg)
	case YAML:
		err = unmarshallYAML(f, cfg)
	case XML:
		err = unmarshallXML(f, cfg)
	default:
		err = errors.New("unknown file type")
	}
	return err
}
