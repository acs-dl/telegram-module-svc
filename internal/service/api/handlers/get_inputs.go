package handlers

import (
	"net/http"

	"github.com/acs-dl/telegram-module-svc/internal/service/api/models"
	"gitlab.com/distributed_lab/ape"
)

func GetInputs(w http.ResponseWriter, r *http.Request) {
	ape.Render(w, models.NewInputsResponse())
}
