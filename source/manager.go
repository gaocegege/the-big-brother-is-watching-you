package source

import (
	"github.com/gaocegege/the-big-brother-is-watching-you/source/watcher"
)

// Manager is the object to manage all the sources
type Manager struct {
	sources []Source
}

// NewManager returns a new Manager Object
func NewManager() (*Manager, error) {
	m := &Manager{
		sources: make([]Source, 0),
	}

	mock, err := watcher.NewMock()
	if err != nil {
		return nil, err
	}

	m.registerSource(mock)
	return m, nil
}

// GetSources returns all the sources registered to manager
func (m *Manager) GetSources() []Source {
	return m.sources
}

// registerSource is the function to register Source to Manager
func (m *Manager) registerSource(s Source) {
	m.sources = append(m.sources, s)
}
