package logger

import (
	"log"
	"time"

	"github.com/matzhouse/causal/pkg/alerters"
)

// PackageName can be referenced for an identifier to this functionality
var PackageName = "log"

// Logger is a simple log output Alerter for testing
type Logger struct{}

// Alert takes an Alert value and prints out the details using the
// standard library log
func (l *Logger) Alert(a *alerters.Alert) (err error) {

	log.Printf("Received '%s' at %s from %s \n", a.Message, a.Time.Format(time.RFC1123), a.Name)

	return nil
}
