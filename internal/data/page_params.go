package data

import (
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
	"strconv"

	"github.com/acs-dl/telegram-module-svc/resources"
	"gitlab.com/distributed_lab/kit/pgdb"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	pageParamLimit  = "page[limit]"
	pageParamNumber = "page[number]"
	pageParamCursor = "page[cursor]"
	pageParamOrder  = "page[order]"
)

// OrderType - represents sorting order of the query
type OrderType string

const (
	// OrderAsc - ascending order
	OrderAsc OrderType = "asc"
	// OrderDesc - descending order
	OrderDesc OrderType = "desc"
)

const (
	OrderAscending         = "asc"
	OrderDescending        = "desc"
	defaultLimit    uint64 = 15
	maxLimit        uint64 = 100
)

// Invert - inverts order by
func (o OrderType) Invert() OrderType {
	switch o {
	case OrderDesc:
		return OrderAsc
	case OrderAsc:
		return OrderDesc
	default:
		panic(errors.From(errors.New("unexpected order type"), logan.F{
			"order_type": o,
		}))
	}
}

func GetOffsetLinksForPGParams(r *http.Request, p pgdb.OffsetPageParams) *resources.Links {
	result := resources.Links{
		Next: getOffsetLink(r, p.PageNumber+1, p.Limit, OrderType(p.Order)),
		Self: getOffsetLink(r, p.PageNumber, p.Limit, OrderType(p.Order)),
	}

	return &result
}

func getOffsetLink(r *http.Request, pageNumber, limit uint64, order OrderType) string {
	u := r.URL
	query := u.Query()
	query.Set(pageParamNumber, strconv.FormatUint(pageNumber, 10))
	query.Set(pageParamLimit, strconv.FormatUint(limit, 10))
	query.Set(pageParamOrder, string(order))
	u.RawQuery = query.Encode()
	return u.String()
}
