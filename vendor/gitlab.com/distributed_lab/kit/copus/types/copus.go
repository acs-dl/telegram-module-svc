package types

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3"
)

// Copuser creates new instance of Copus from config
type Copuser interface {
	Copus() Copus
}

// Copus used to register service endpoints inside external API Gateway
type Copus interface {
	WithLog(log *logan.Entry) Copus
	RegisterGojiEndpoint(endpoint, method string) error
	RegisterChi(r chi.Router) error
}
