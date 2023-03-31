package problems

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
)

func TooManyRequests() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusTooManyRequests),
		Status: fmt.Sprintf("%d", http.StatusTooManyRequests),
	}
}
