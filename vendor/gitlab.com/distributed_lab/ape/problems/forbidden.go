package problems

import (
    "fmt"
    "net/http"

    "github.com/google/jsonapi"
)

func Forbidden() *jsonapi.ErrorObject {
    return &jsonapi.ErrorObject{
        Title:  http.StatusText(http.StatusForbidden),
        Status: fmt.Sprintf("%d", http.StatusForbidden),
    }
}
