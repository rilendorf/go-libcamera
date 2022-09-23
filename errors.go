package libcamera

import (
	"errors"
)

var (
	InvalidRunDuration = errors.New("Run duration invalid, possibly nil value (0)")
)
