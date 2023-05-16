package problems

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
)

func NotFound() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusNotFound),
		Status: fmt.Sprintf("%d", http.StatusNotFound),
	}
}
