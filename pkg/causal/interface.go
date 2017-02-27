package causal

import (
	"github.com/matzhouse/causal/pkg/alerters"
	"github.com/matzhouse/causal/pkg/watchers"
)

// Watcher has a Run function that enables it
// to report an alert
type Watcher interface {
	Run() (*watchers.Result, error)
}

// Alerter has an alert function that enables
// it to fire an alert
type Alerter interface {
	Alert(*alerters.Alert) error
}
