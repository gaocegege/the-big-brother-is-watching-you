package watcher

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gaocegege/the-big-brother-is-watching-you/common"
)

// Mock is the mock source, which use file as source
type Mock struct {
	filePath string
	host     string
	counter  int
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
		host:     common.MockOrigin,
		counter:  1,
	}, nil
}

// FetchFromOrigin implements the Source interface
func (m *Mock) FetchFromOrigin(vendorID string, beginTime time.Time) ([]common.Record, error) {
	var records = make([]common.Record, 0)
	var record = common.Record{
		VendorID:   vendorID,
		Content:    fmt.Sprintf("For dev: %d", m.counter),
		CreateTime: time.Now(),
		URL:        fmt.Sprintf("%s/%d", m.host, m.counter),
	}

	records = append(records, record)
	m.counter = m.counter + 1
	return records, nil
}

// GetHostName implements the Source interface
func (m *Mock) GetHostName() string {
	return m.host
}
