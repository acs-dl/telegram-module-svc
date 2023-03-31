package pgdb

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

// Sorts is a slice of sorting params that can be applied directly to sql query.
type Sorts []Sort

// Sort is a string that describes one sort parameter
// defined by client.
type Sort string

// Desc returns whether sorting by specific column
// should be performed in descending order.
func (s Sort) Desc() bool {
	return strings.HasPrefix(string(s), "-")
}

func (s Sort) order() string {
	if s.Desc() {
		return OrderTypeDesc
	}

	return OrderTypeAsc
}

// ApplyTo applies sorts to a prepared sql statement. Takes a map of <query-param>:<column> names (without minuses).
// Like:
// 	 map[string]string{}{
// 		"created_at": "books.created_at",
// 		"author.name": "authors.name",
//   }
// Panics if sort parameter in a slice isn't provided in "columns" map.
func (sorts Sorts) ApplyTo(stmt squirrel.SelectBuilder, columns map[string]string) squirrel.SelectBuilder {
	for _, sort := range sorts {
		column, ok := columns[strings.TrimPrefix(string(sort), "-")]
		if !ok {
			panic(errors.Errorf("unknown sort parameter: %s", sort))
		}

		stmt = stmt.OrderBy(fmt.Sprintf("%s %s", column, sort.order()))
	}

	return stmt
}
