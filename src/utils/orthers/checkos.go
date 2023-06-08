package others

import (
	"runtime"
)

func CheckOS() string {
	os := runtime.GOOS
	return os
}
