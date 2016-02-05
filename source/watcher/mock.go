package watcher

import (
	"errors"
	"os"
	"time"

	"github.com/gaocegege/the-big-brother-is-watching-you/common"
)

// Mock is the mock source, which use file as source
type Mock struct {
	filePath string
}

// NewMock returns a Mock Object
func NewMock(filePath string) (*Mock, error) {
	file, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	if file.IsDir() {
		return nil, errors.New("In mock mode, filePath must be a file")
	}
	return &Mock{
		filePath: filePath,
	}, nil
}

// FetchFromOrigin implements the Source interface
func (m *Mock) FetchFromOrigin(beginTime time.Time) ([]common.Record, error) {
	return nil, nil
}
