// +build !windows

package disk

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

// ErrDeviceNotFound should be returned when a device cannot be found
// when doing the df command
var ErrDeviceNotFound = errors.New("device cannot be found")

func getDeviceFullness(device string) (fullness int, err error) {

	// use df -h as it's slightly smaller than df
	c := exec.Command("df", "-h")

	stdout, err := c.StdoutPipe()
	if err != nil {
		return 0, err
	}

	if err := c.Start(); err != nil {
		return 0, err
	}

	fullness, err = processDFOutput(stdout, device)

	if err != nil {
		return 0, err
	}

	if err := c.Wait(); err != nil {
		return 0, err
	}

	return fullness, err

}

func processDFOutput(in io.Reader, device string) (fullness int, err error) {

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {

		lineSplit := strings.Fields(scanner.Text())

		if lineSplit[0] == device {
			return strconv.Atoi(strings.Replace(lineSplit[4], "%", "", -1))
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}

	return 0, ErrDeviceNotFound

}
