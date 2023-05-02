package models

import (
	"gitlab.com/distributed_lab/acs/telegram-module/resources"
)

func NewEstimatedTimeModel(estimatedTime string) resources.EstimatedTime {
	return resources.EstimatedTime{
		Key: resources.Key{
			ID:   string(resources.ESTIMATED_TIME),
			Type: resources.ESTIMATED_TIME,
		},
		Attributes: resources.EstimatedTimeAttributes{
			EstimatedTime: estimatedTime,
		},
	}
}

func NewEstimatedTimeResponse(estimatedTime string) resources.EstimatedTimeResponse {
	return resources.EstimatedTimeResponse{
		Data: NewEstimatedTimeModel(estimatedTime),
	}
}
