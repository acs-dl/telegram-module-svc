package cop

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gitlab.com/distributed_lab/kit/copus/types"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CopConfig struct {
	Endpoint      string `fig:"endpoint,required"`
	Upstream      string `fig:"upstream,required"`
	ServiceName   string `fig:"service_name,required"`
	ServicePort   string `fig:"service_port,required"`
	ServicePrefix string `fig:"service_prefix"`
}

type Cop struct {
	disabled bool
	client   Client
	config   CopConfig
	log      *logan.Entry
}

func NewNoOp() *Cop {
	return &Cop{
		disabled: true,
	}
}

func New(config CopConfig) *Cop {
	return &Cop{
		config: config,
		log:    logan.New().Out(ioutil.Discard),
		client: Client{
			Endpoint: config.Endpoint,
		},
	}
}

func (c *Cop) WithLog(log *logan.Entry) types.Copus {
	c.log = log
	return c
}

func (c *Cop) RegisterGojiEndpoint(endpoint, method string) error {
	if c.disabled {
		return nil
	}

	service := Service{Data: ServiceData{
		ID:   c.config.ServiceName,
		Type: "traefik-service",
		Attributes: ServiceAttributes{
			Name: c.config.ServiceName,
			Url:  c.config.Upstream,
			Port: c.config.ServicePort,
			Rule: fmt.Sprintf("(Path(`%s`)&&Method(`%s`))", getRouteForGoji(endpoint), method),
		},
	}}

	err := c.client.AddService(service)
	if err != nil {
		return errors.Wrap(err, "failed to add service")
	}

	return nil
}

func (c *Cop) RegisterChi(r chi.Router) error {
	if c.disabled {
		return nil
	}

	rule, err := c.GetRule(r)
	if err != nil {
		return errors.Wrap(err, "failed to get rule")
	}

	service := Service{Data: ServiceData{
		ID:   c.config.ServiceName,
		Type: "traefik-service",
		Attributes: ServiceAttributes{
			Name: c.config.ServiceName,
			Url:  c.config.Upstream,
			Port: c.config.ServicePort,
			Rule: rule,
		},
	}}

	err = c.client.AddService(service)
	if err != nil {
		return errors.Wrap(err, "failed to add service")
	}

	return nil
}

func (c *Cop) GetRule(r chi.Router) (string, error) {
	if c.config.ServicePrefix != "" {
		rule := fmt.Sprintf("PathPrefix(`%s`)", c.config.ServicePrefix)
		return rule, nil
	}

	var routes []string
	walk := func(method, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		if len(route) > 1 {
			route = strings.TrimRight(route, "/")
		}

		routes = append(routes, fmt.Sprintf("(Path(`%s`)&&Method(`%s`))", route, method))
		return nil
	}

	err := chi.Walk(r, walk)
	if err != nil {
		return "", errors.Wrap(err, "failed to walk router")
	}
	return strings.Join(routes, "||"), nil
}
func getRouteForGoji(endpoint string) string {
	sp := strings.Split(endpoint[1:], "/")
	var result string
	for _, s := range sp {
		if strings.HasPrefix(s, ":") {
			s = fmt.Sprintf("{%s}", strings.TrimPrefix(s, ":"))
		}
		result = fmt.Sprintf("%s/%s", result, s)
	}
	return result
}
