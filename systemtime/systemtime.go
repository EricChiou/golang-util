package systemtime

import (
	"time"
)

// GetSysTimeInSec get system time in second
func GetSysTimeInSec() int {
	return int(time.Now().Unix())
}

func GetSysTimeInNanoSec() int {
	return int(time.Now().UnixNano())
}
