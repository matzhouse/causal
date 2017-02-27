package alerters

import (
	"github.com/matzhouse/causal/pkg/watchers"
	"time"
)

// Alert is a struct to pass a message from an Watcher to
// an alerter.
type Alert struct {
	Name    string
	Type    string
	Message string
	Time    time.Time
	Result  *watchers.Result
}
