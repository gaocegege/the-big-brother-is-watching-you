package source

import (
	"errors"

	"github.com/gaocegege/the-big-brother-is-watching-you/source/watcher"
)

// Manager is the object to manage all the sources
type Manager struct {
	sources []Source
}

// NewManager returns a new Manager Object
func NewManager(mockfilePath string) (*Manager, error) {
	m := &Manager{
		sources: make([]Source, 0),
	}

	if mockfilePath != "" {
		mock, err := watcher.NewMock(mockfilePath)
		if err != nil {
			return nil, err
		}

		m.registerSource(mock)
		return m, nil
	}

	return nil, errors.New("No watcher registered to manager.")
}

// GetSources returns all the sources registered to manager
func (m *Manager) GetSources() []Source {
	return m.sources
}

// registerSource is the function to register Source to Manager
func (m *Manager) registerSource(s Source) {
	m.sources = append(m.sources, s)
}
