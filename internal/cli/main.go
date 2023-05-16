package cli

import (
	"gitlab.com/distributed_lab/logan/v3"
	"os"
	"os/signal"
	"syscall"

	"github.com/acs-dl/telegram-module-svc/internal/config"
	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/registrator"
	"github.com/acs-dl/telegram-module-svc/internal/service"
	"github.com/alecthomas/kingpin"
	"gitlab.com/distributed_lab/kit/kv"
)

func Run(args []string) bool {
	log := logan.New()

	cfg := config.New(kv.MustFromEnv())
	log = cfg.Log()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
			err := registrator.NewRegistrar(cfg).UnregisterModule()
			if err != nil {
				log.WithError(err).Errorf("failed to unregister module %s", data.ModuleName)
			}
			log.Infof("unregistered module %s", data.ModuleName)
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		sig := <-signalChannel
		log.Infof("service was interrupted by signal `%s`", sig.String())
		err := registrator.NewRegistrar(cfg).UnregisterModule()
		if err != nil {
			log.WithError(err).Errorf("failed to unregister module %s", data.ModuleName)
			os.Exit(1)
		}
		log.Infof("unregistered module %s", data.ModuleName)
		os.Exit(0)
	}()

	app := kingpin.New(data.ModuleName, "")

	runCmd := app.Command("run", "run command")
	serviceCmd := runCmd.Command("service", "run service") // you can insert custom help

	migrateCmd := app.Command("migrate", "migrate command")
	migrateUpCmd := migrateCmd.Command("up", "migrate db up")
	migrateDownCmd := migrateCmd.Command("down", "migrate db down")

	// custom commands go here...

	cmd, err := app.Parse(args[1:])
	if err != nil {
		log.WithError(err).Error("failed to parse arguments")
		return false
	}

	switch cmd {
	case serviceCmd.FullCommand():
		service.Run(cfg)
	case migrateUpCmd.FullCommand():
		err = MigrateUp(cfg)
	case migrateDownCmd.FullCommand():
		err = MigrateDown(cfg)
	// handle any custom commands here in the same way
	default:
		log.Errorf("unknown command %s", cmd)
		return false
	}
	if err != nil {
		log.WithError(err).Error("failed to exec cmd")
		return false
	}
	return true
}
