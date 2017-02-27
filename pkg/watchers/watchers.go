package watchers

import (
	"errors"
	"time"
)

// ErrNotImplemented can be returned on a Run command to say that it won't work
// on a machine type
var ErrNotImplemented = errors.New("Run not implemented on this OS or Architecture")

// State describes the state of the Run on a watcher
type State int

// Constants for states
const (
	OK      State = iota // 1 - an alert should not be fired on an OK state
	FAIL                 // 2 - an alert should be fired on a FAIL state
	UNKNOWN              // 3
)

// Result shows the result of a watcher run
type Result struct {
	Timestamp time.Time
	Message   string
	State     State
}

func (r *Result) String() string {
	return r.Message
}
