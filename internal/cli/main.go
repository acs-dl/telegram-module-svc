package cli

import (
	"github.com/alecthomas/kingpin"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

func Run(args []string) bool {
	log := logan.New()

	cfg := config.New(kv.MustFromEnv())
	log = cfg.Log()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
		}
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
