package running

import (
	"context"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"sync"
	"time"

	"gitlab.com/distributed_lab/kit/kv"
)

// Server creates http.Server and makes it to ListenAndServe
func Server(
	ctx context.Context,
	log *logan.Entry,
	config ServerConfig,
	handler http.Handler,
) {
	var server *http.Server

	// Starting server asynchronously
	go UntilSuccess(ctx, log, "listening_server", func(ctx context.Context) (bool, error) {
		server = &http.Server{
			Addr:         config.Address,
			Handler:      handler,
			WriteTimeout: config.RequestWriteTimeout,
		}

		log.WithField("server", config).Infof("Starting Server listening on %s.", config.Address)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			return false, errors.Wrap(err, "failed to ListenAndServe (Server stopped with error)")
		}

		if IsCancelled(ctx) {
			// To avoid 'unsuccessful' log.
			return true, nil
		}

		// Something really strange - Server stopped with `http.ErrServerClosed`,
		// however ctx is not cancelled. Should actually never happen, but just in case.
		return false, nil
	}, time.Second, time.Hour)

	<-ctx.Done()

	log.Infof("Will start stopping the server after (%s).", config.ShutdownPause.String())
	time.Sleep(config.ShutdownPause)

	log.Info("Stopping the server.")
	if config.ShutdownTimeout != 0 {
		// Force shut down the server after config.ShutdownTimeout.
		shutdownCtx, _ := context.WithTimeout(context.Background(), config.ShutdownTimeout)
		server.Shutdown(shutdownCtx)
	} else {
		// Don't shutdown the server until all current requests are finished
		server.Shutdown(context.Background())
	}
	log.Info("Server stopped cleanly.")
}

// TODO Add optional bearer token into config and check Authorization for all requests
type ServerConfig struct {
	Address             string        `fig:"address,required"`
	RequestWriteTimeout time.Duration `fig:"request_write_timeout,required"`

	ShutdownPause   time.Duration `fig:"shutdown_pause"`   // optional
	ShutdownTimeout time.Duration `fig:"shutdown_timeout"` // optional
}

func (s ServerConfig) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"address":               s.Address,
		"request_write_timeout": s.RequestWriteTimeout,

		"shutdown_pause":   s.ShutdownPause,
		"shutdown_timeout": s.ShutdownTimeout,
	}
}

type Serverer interface {
	Server() ServerConfig
}

func NewServerer(getter kv.Getter) Serverer {
	return &serverer{
		getter: getter,
	}
}

type serverer struct {
	getter kv.Getter
	once   sync.Once
	value  ServerConfig
	err    error
}

func (s *serverer) Server() ServerConfig {
	s.once.Do(func() {
		var config ServerConfig

		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(s.getter, "server")).
			Please()
		if err != nil {
			s.err = errors.Wrap(err, "failed to figure out server")
			return
		}

		s.value = config
	})

	if s.err != nil {
		panic(s.err)
	}

	return s.value
}
