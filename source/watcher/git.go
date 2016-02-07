package watcher

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gaocegege/the-big-brother-is-watching-you/common"
	
	"github.com/bitly/go-simplejson"
)

const (
	githubAPI = "https://api.github.com/users/%s/events"
)

// Git is the github source
type Git struct {
	host     string
	username string
}

// GitRecord is a record from the api.github.com
type GitRecord struct {
	Payload  string `json:"payload,omitempty"`
	CreatedAt string `json:created_at,omitempty`
}

// Payload is the field in GitRecord
type Payload struct {
	Action string `json:"action,omitempty"`
}

// NewGit returns a new Git object
func NewGit(username string) (*Git, error) {
	return &Git{
		host:     common.GithubOrigin,
		username: username,
	}, nil
}

// FetchFromOrigin implements the Source interface
func (g *Git) FetchFromOrigin(vendorID string, t time.Time) ([]common.Record, error) {
	githubURL := fmt.Sprintf(githubAPI, g.username)
	res, err := http.Get(githubURL)
	if err != nil {
		log.Fatal(err)
	}
	// binRes, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	v, err := simplejson.NewFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// records, err := v.Array()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Printf("Created_at: %s", v.GetIndex(0).Get("created_at"))
	
	return nil, nil
}

// GetHostName implements the Source interface
func (g *Git) GetHostName() string {
	return g.host
}
