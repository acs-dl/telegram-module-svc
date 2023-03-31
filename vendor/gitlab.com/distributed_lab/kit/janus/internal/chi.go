package internal

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

type Chi struct {
	r chi.Router
}

func NewChi(r chi.Router) *Chi {
	return &Chi{r}
}

func (c *Chi) Services() ([]Service, error) {
	var services []Service
	walk := func(method, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		if len(route) > 1 {
			route = strings.TrimRight(route, "/")
		}
		methods := []string{method}
		name := GetName(route, method)
		services = append(services, Service{
			Name:   name,
			Active: true,
			Proxy: Proxy{
				AppendPath: true,
				ListenPath: route,
				Methods:    methods,
			},
		})
		return nil
	}
	err := chi.Walk(c.r, walk)
	if err != nil {
		return nil, errors.Wrap(err, "failed to walk router")
	}
	return services, nil
}
