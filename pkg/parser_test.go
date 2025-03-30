package pkg

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/maikelh/go-arazzo/pkg/arazzo"
	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		filename string
		fileType FileType
		want     *arazzo.Document
		wantErr  bool
	}{
		{
			name:     "valid YAML",
			input:    "key: value",
			fileType: YAML,
			want:     &arazzo.Document{},
			wantErr:  false,
		},
		{
			name:     "invalid YAML",
			input:    "invalid: : yaml",
			fileType: YAML,
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "unsupported file type",
			input:    "some content",
			fileType: "invalid",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "valid YAML",
			filename: "test/oauth.arazzo.yaml",
			fileType: YAML,
			want:     nil,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var reader io.Reader
			var err error
			if tt.filename != "" {
				workingDir, err := os.Getwd()
				fullPath := filepath.Join(workingDir, tt.filename)
				reader, err = os.Open(fullPath)
				if err != nil {
					t.Fatalf("failed to open file: %v", err)
					return
				}
			} else {
				reader = strings.NewReader(tt.input)
			}
			got, err := ParseFile(reader, tt.fileType)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}
