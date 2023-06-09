package handlers

import (
	"net/http"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"gitlab.com/distributed_lab/ape"
)

func GetUserRolesMap(w http.ResponseWriter, r *http.Request) {
	result := newModuleRolesResponse()

	result.Data.Attributes["super_admin"] = data.Roles[data.Owner]
	result.Data.Attributes["admin"] = data.Roles[data.Admin]
	result.Data.Attributes["user"] = data.Roles[data.Member]

	ape.Render(w, result)
}
