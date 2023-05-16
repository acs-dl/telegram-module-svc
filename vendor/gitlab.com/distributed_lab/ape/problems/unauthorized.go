package problems

import (
	"fmt"
	"github.com/google/jsonapi"
	"net/http"
)

func Unauthorized() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusUnauthorized),
		Status: fmt.Sprintf("%d", http.StatusUnauthorized),
	}
}
