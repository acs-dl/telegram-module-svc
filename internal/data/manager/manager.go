package manager

import (
	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/data/postgres"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Manager struct {
	db *pgdb.DB

	responses   data.Responses
	permissions data.Permissions
	users       data.Users
	links       data.Links
}

func NewManager(db *pgdb.DB) *Manager {
	return &Manager{
		db:          db,
		responses:   postgres.NewResponsesQ(db),
		permissions: postgres.NewPermissionsQ(db),
		users:       postgres.NewUsersQ(db),
		links:       postgres.NewLinksQ(db),
	}
}

func (m *Manager) Transaction(fn func() error) error {
	return m.db.Transaction(fn)
}
