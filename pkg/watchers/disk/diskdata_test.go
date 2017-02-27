package disk

import (
	"bytes"
	"testing"
)

var dfOutput = `Filesystem                          Size   Used  Avail Capacity iused      ifree %iused  Mounted on
/dev/disk2                         106Gi   56Gi   50Gi    53% 1128115 4293839164    0%   /
/dev/disk3                         222Gi  183Gi   40Gi    83% 1915296 4293051983    0%   /Volumes/hd1`

func TestDisk_processDFOutput(t *testing.T) {

	b := bytes.NewReader([]byte(dfOutput))

	fullness, err := processDFOutput(b, "/dev/disk2")

	if err != nil {
		t.Log("error found when none expected", err)
		t.Fail()
	}

	if fullness != 53 {
		t.Log("Unexpected value returned -", fullness)
		t.Fail()
	}

}
