package disk

import (
	"fmt"
	"time"

	"github.com/matzhouse/causal/pkg/watchers"
)

// DiskWatcher contains the data to watch a disk device and
// alert when a disk becomes too full
type Watcher struct {
	device       string
	percentLimit int
}

// New returns a DiskWatcher struct. The device should be the device label
// when doing df (in Linux environments). Windows not implemented yet.
func New(device string, percentLimit int) *Watcher {

	return &Watcher{
		device:       device,
		percentLimit: percentLimit,
	}

}

// Run runs the disk watcher and reports the size of the disk
func (dw *Watcher) Run() (r *watchers.Result, err error) {

	// Create a result to return
	r = &watchers.Result{
		Timestamp: time.Now(),
		Message:   "",
		State:     watchers.OK,
	}

	// Check the fullness of the devices
	fullness, err := getDeviceFullness(dw.device)

	if err != nil {
		r.Message = fmt.Sprintf("error in diskwatcher: %s", err)
		return r, err
	}

	// Set the message with the percent full data
	r.Message = fmt.Sprintf("disk %s is %d%% full", dw.device, fullness)

	// Check if the fullness of the disk is above what is expected
	if fullness > dw.percentLimit {
		r.State = watchers.FAIL
	}

	return

}
