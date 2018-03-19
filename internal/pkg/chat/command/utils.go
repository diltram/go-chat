package command

import (
	"strings"
)

// RemoveCmd function removes from a text line send by user prefix containing
// call to command like /channel or other. In case if text after the operation
// starts with space (should be always) then remove it from beginning.
func RemoveCmd(name string, line string) string {
	line = strings.TrimLeft(line, name)
	if strings.HasPrefix(line, " ") {
		line = line[1:]
	}

	return line
}
