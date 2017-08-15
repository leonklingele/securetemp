package securetemp

import (
	"errors"
)

func createRAMDisk(size int) (string, func(), error) {
	return "", func() {}, errors.New("windows is currently unsupported")
}
