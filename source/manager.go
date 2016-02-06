package source

import (
	"github.com/gaocegege/the-big-brother-is-watching-you/source/watcher"
)

// Manager is the object to manage all the sources
type Manager struct {
	sources []Source
}

// NewManager returns a new Manager Object
func NewManager(mockMode bool, githubUsername string) (*Manager, error) {
	m := &Manager{
		sources: make([]Source, 0),
	}

	// add mock to manager if mock mode is true
	if mockMode == true {
		mock, err := watcher.NewMock()
		if err != nil {
			return nil, err
		}
		m.registerSource(mock)
	}

	// add github to manager if the username is given
	if githubUsername != "" {
		git, err := watcher.NewGit(githubUsername)
		if err != nil {
			return nil, err
		}
		m.registerSource(git)
	}

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
