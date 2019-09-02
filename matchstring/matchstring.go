package matchstring

import (
	"regexp"
)

// Match09AZ match string 0-9, a-z
func Match09azAZ(s string) (bool, error) {
	match, err := regexp.MatchString("^[0-9a-zA-Z]+$", s)
	return match, err
}
