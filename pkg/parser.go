package pkg

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/maikelh/go-arazzo/pkg/arazzo"
	"gopkg.in/yaml.v3"
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
	var doc arazzo.Document
	data, err := io.ReadAll(reader)
	if err != nil {
		slog.Error("failed to read YAML file", "error", err)
		return nil, err
	}
	err = yaml.Unmarshal(data, &doc)
	if err != nil {
		slog.Error("failed to parse YAML file", "error", err)
		return nil, err
	}
	return &doc, nil
}

func parseJSON(reader io.Reader) (*arazzo.Document, error) {
	panic("unimplemented")
}
