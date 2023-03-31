package urlval

import (
	"strings"
)

// Sort is a string that describes one sort parameter
// defined by client.
type Sort string

// Desc returns whether sorting by specific column
// should be performed in descending order.
func (s Sort) Desc() bool {
	return strings.HasPrefix(string(s), "-")
}

// Key of the sort parameter.
func (s Sort) Key() string {
	return strings.TrimPrefix(string(s), "-")
}
