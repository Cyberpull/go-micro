package gosrv

import (
	"strings"

	"cyberpull.com/gotk/errors"
)

const (
	requestState string = "REQUEST"
)

func parseState(v string) (name string, value bool, err error) {
	chunks := strings.SplitN(v, separator, 2)

	if chunks[0] != statePrefix {
		err = errors.New("Invalid " + statePrefix)
	}

	name, value = parseStateValue(chunks[1])

	return
}

func parseStateValue(v string) (name string, value bool) {
	chunks := strings.SplitN(v, "=", 2)

	name = chunks[0]

	if len(chunks) == 2 && chunks[1] == "YES" {
		value = true
	}

	return
}
