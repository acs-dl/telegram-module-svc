package pgdb

import (
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

// OffsetPageParams defines page params for offset-based pagination.
type SortedOffsetPageParams struct {
	Limit      uint64   `page:"limit" default:"15"`
	Sort       []string `url:"sort"`
	PageNumber uint64   `page:"number"`
}

// ApplyTo returns a new SelectBuilder after applying the paging effects of  `p` to `sql`.
/**
 * columns should be a map of <attribute>:<table-column>
 * map[string]string{}{
 * 	"views": "table1.views",
 *	"created_at": "table3.created_at"
 * }
 * results in: ORDER BY table1.views desc, table3.created_at desc
 */

func (p *SortedOffsetPageParams) ApplyTo(sql squirrel.SelectBuilder, columns map[string]string) squirrel.SelectBuilder {
	if p.Limit == 0 {
		p.Limit = 15
	}

	offset := p.Limit * p.PageNumber

	sql = sql.Limit(p.Limit).Offset(offset)

	for _, sort := range p.Sort {
		trimmed := strings.TrimSpace(sort)
		column, ok := columns[strings.TrimPrefix(string(trimmed), "-")]

		if !ok {
			panic(errors.Errorf("unknown sort parameter: %s", sort))
		}

		order := OrderTypeAsc

		if strings.HasPrefix(string(trimmed), "-") {
			order = OrderTypeDesc
		}

		sql = sql.OrderBy(fmt.Sprintf("%s %s", column, order))
	}

	return sql
}
