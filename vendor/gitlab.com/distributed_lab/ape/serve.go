package ape

import (
	"context"
	"net"
	"net/http"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
)

//ServeOpts - defines serve configuration options that must be specified by developer
type ServeOpts struct {
	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body. Default is 5s.
	ReadTimeout time.Duration
	// WriteTimeout, unless connection is HTTPS (in our case never as we are always behind nginx),
	// covers the time from the end of the request header read to the end of the response write
	// (a.k.a. the lifetime of the ServeHTTP). Default 10s
	WriteTimeout time.Duration
	// ShutdownTimeout is maximum duration for waiting for server to shutdown. Default 15s.
	ShutdownTimeout time.Duration
}

//ServeConfig - defines external dependencies of serve
type ServeConfig interface {
	Log() *logan.Entry
	Listener() net.Listener
}

//Serve - accepts incoming connections on the Listener. Similar to http.Server, but with default sane timeouts suitable
// for generic REST API and graceful shutdown.
func Serve(ctx context.Context, handler http.Handler, cfg ServeConfig, opts ServeOpts) {
	if opts.ShutdownTimeout == 0 {
		opts.ShutdownTimeout = 15 * time.Second
	}

	if opts.ReadTimeout == 0 {
		opts.ReadTimeout = 5 * time.Second
	}

	if opts.WriteTimeout == 0 {
		opts.WriteTimeout = 10 * time.Second
	}

	log := cfg.Log().WithField("service", "http_serve")

	srv := &http.Server{
		ReadTimeout:  opts.ReadTimeout,
		WriteTimeout: opts.WriteTimeout,
		Handler:      handler,
	}

	serveDone := make(chan struct{})
	go func() {
		defer close(serveDone)
		err := srv.Serve(cfg.Listener())
		if err == http.ErrServerClosed {
			log.Info("stopped accepting new connections")
			return
		}

		if err != nil {
			cfg.Log().WithError(err).Error("died")
		}
	}()

	select {
	case <-serveDone:
		// error is already logged
		return
	case <-ctx.Done():
		log.WithField("shutdown_timeout", opts.ShutdownTimeout).Info("received signal to shutdown gracefully")
		shutDownCtx, cancel := context.WithTimeout(context.Background(), opts.ShutdownTimeout)
		defer cancel()
		if err := srv.Shutdown(shutDownCtx); err != nil {
			log.WithError(err).Error("failed to shutdown gracefully")
			return
		}

		log.Info("shutdown gracefully")
	}

}
