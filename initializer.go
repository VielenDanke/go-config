package config

import "io"

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

// WithParsingFile initialize option passed to config file for it's opening and unmarshalling based on fileType
func WithParsingFile(filePath string, fileType FileType) configOption {
	return func(cfg interface{}) error {
		return ParseFile(filePath, fileType, cfg)
	}
}

// WithParsingEnv initialize option for parsing ENV
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
