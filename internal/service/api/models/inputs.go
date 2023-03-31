package models

import (
	"gitlab.com/distributed_lab/acs/telegram-module/resources"
)

func NewInputsModel() resources.Inputs {
	result := resources.Inputs{
		Key: resources.Key{
			ID:   "0",
			Type: resources.INPUTS,
		},
		Attributes: resources.InputsAttributes{
			Username:    "string",
			Phone:       "string",
			Link:        "string",
			AccessLevel: "number",
		},
	}

	return result
}

func NewInputsResponse() resources.InputsResponse {
	return resources.InputsResponse{
		Data: NewInputsModel(),
	}
}
