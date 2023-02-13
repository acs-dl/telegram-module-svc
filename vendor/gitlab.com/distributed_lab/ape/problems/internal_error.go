package problems

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
)

func InternalError() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusInternalServerError),
		Status: fmt.Sprintf("%d", http.StatusInternalServerError),
	}
}
