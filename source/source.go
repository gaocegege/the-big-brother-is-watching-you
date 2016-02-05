package source

import (
	"time"

	"github.com/gaocegege/the-big-brother-is-watching-you/common"
)

// Source is the interface, represents a vendor
type Source interface {
	FetchFromOrigin(beginTime time.Time) ([]common.Record, error)
}
