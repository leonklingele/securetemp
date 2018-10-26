package securetemp

import (
	"errors"
)

func createRAMDisk(size int) (string, CleanupFunc, error) {
	return "", func() {}, errors.New("windows is currently unsupported")
}
