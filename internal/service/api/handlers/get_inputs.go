package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/models"
	"gitlab.com/distributed_lab/ape"
)

func GetInputs(w http.ResponseWriter, r *http.Request) {
	ape.Render(w, models.NewInputsResponse())
}
