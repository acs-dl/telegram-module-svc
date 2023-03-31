package pgdb

import (
	"database/sql"
	"github.com/lib/pq"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Databaser interface {
	DB() *DB
	RawDB() *sql.DB
	NewListener() *pq.Listener
}

type databaser struct {
	getter kv.Getter
	once   comfig.Once
}

func NewDatabaser(getter kv.Getter) Databaser {
	return &databaser{
		getter: getter,
	}
}

type databaserCfg struct {
	URL                      string        `fig:"url,required"`
	MaxOpenConnections       int           `fig:"max_open_connection"`
	MaxIdleConnections       int           `fig:"max_idle_connections"`
	ListenerMinRetryDuration time.Duration `fig:"listener_min_retry_duration"`
	ListenerMaxRetryDuration time.Duration `fig:"listener_max_retry_duration"`
}

func (d *databaser) readConfig() databaserCfg {
	config := databaserCfg{
		MaxOpenConnections:       12,
		MaxIdleConnections:       12,
		ListenerMinRetryDuration: time.Second,
		ListenerMaxRetryDuration: time.Minute,
	}
	err := figure.Out(&config).
		From(kv.MustGetStringMap(d.getter, "db")).
		Please()
	if err != nil {
		panic(errors.Wrap(err, "failed to figure out"))
	}

	return config
}

func (d *databaser) DB() *DB {
	return d.once.Do(func() interface{} {
		config := d.readConfig()

		db, err := Open(Opts{
			URL:                config.URL,
			MaxOpenConnections: config.MaxOpenConnections,
			MaxIdleConnections: config.MaxIdleConnections,
		})
		if err != nil {
			panic(errors.Wrap(err, "failed to open database"))
		}

		return db
	}).(*DB)
}

//NewListener - returns new listener for notify events
func (d *databaser) NewListener() *pq.Listener {
	config := d.readConfig()
	listener := pq.NewListener(config.URL, config.ListenerMinRetryDuration, config.ListenerMaxRetryDuration, nil)
	return listener
}

func (d *databaser) RawDB() *sql.DB {
	return d.DB().RawDB()
}
