package pkg

import (
	"fmt"
	"io"

	"github.com/maikelh/go-arazzo/pkg/arazzo"
)

type FileType string

const (
	JSON FileType = "json"
	YAML FileType = "yaml"
)

func ParseFile(reader io.Reader, filetype FileType) (*arazzo.Document, error) {
	switch filetype {
	case JSON:
		return parseJSON(reader)
	case YAML:
		return parseYAML(reader)
	}
	return nil, fmt.Errorf("unsupported file type: %s", filetype)
}

func parseYAML(reader io.Reader) (*arazzo.Document, error) {
	panic("unimplemented")
}

func parseJSON(reader io.Reader) (*arazzo.Document, error) {
	panic("unimplemented")
}
