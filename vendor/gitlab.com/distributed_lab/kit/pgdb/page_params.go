package pgdb

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

const (
	// OrderTypeAsc means result should be sorted in ascending order.
	OrderTypeAsc = "asc"
	// OrderTypeDesc means result should be sorted in descending order.
	OrderTypeDesc = "desc"
)

// OffsetPageParams defines page params for offset-based pagination.
type OffsetPageParams struct {
	Limit      uint64 `page:"limit" default:"15" json:"limit"`
	Order      string `page:"order" default:"desc" json:"order"`
	PageNumber uint64 `page:"number" json:"number"`
}

// ApplyTo returns a new SelectBuilder after applying the paging effects of  `p` to `sql`.
func (p *OffsetPageParams) ApplyTo(sql squirrel.SelectBuilder, cols ...string) squirrel.SelectBuilder {
	if p.Limit == 0 {
		p.Limit = 15
	}
	if p.Order == "" {
		p.Order = OrderTypeDesc
	}

	offset := p.Limit * p.PageNumber

	sql = sql.Limit(p.Limit).Offset(offset)

	switch p.Order {
	case OrderTypeAsc:
		for _, col := range cols {
			sql = sql.OrderBy(fmt.Sprintf("%s %s", col, "asc"))
		}
	case OrderTypeDesc:
		for _, col := range cols {
			sql = sql.OrderBy(fmt.Sprintf("%s %s", col, "desc"))
		}
	default:
		panic(fmt.Errorf("unexpected order type: %v", p.Order))
	}

	return sql
}

//CursorPageParams - page params of the db query
type CursorPageParams struct {
	Cursor uint64 `page:"cursor" json:"cursor"`
	Order  string `page:"order" json:"order"`
	Limit  uint64 `page:"limit" json:"limit"`
}

// ApplyTo returns a new SelectBuilder after applying the paging effects of
// `p` to `sql`.  This method provides the default case for paging: int64
// cursor-based paging by an id column.
func (p *CursorPageParams) ApplyTo(sql squirrel.SelectBuilder, col string) squirrel.SelectBuilder {
	if p.Limit == 0 {
		p.Limit = 15
	}
	if p.Order == "" {
		p.Order = OrderTypeDesc
	}

	sql = sql.Limit(p.Limit)

	switch p.Order {
	case OrderTypeAsc:
		if p.Cursor != 0 {
			sql = sql.Where(fmt.Sprintf("%s > ?", col), p.Cursor)
		}

		sql = sql.OrderBy(fmt.Sprintf("%s asc", col))
	case OrderTypeDesc:
		if p.Cursor != 0 {
			sql = sql.
				Where(fmt.Sprintf("%s < ?", col), p.Cursor)
		}
		sql = sql.OrderBy(fmt.Sprintf("%s desc", col))
	default:
		panic(fmt.Errorf("unexpected order type: %v", p.Order))
	}

	return sql
}
