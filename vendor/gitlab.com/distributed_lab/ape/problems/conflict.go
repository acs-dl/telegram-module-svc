package problems

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
)

func Conflict() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusConflict),
		Status: fmt.Sprintf("%d", http.StatusConflict),
	}
}
