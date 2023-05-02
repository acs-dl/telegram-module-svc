package models

import (
	"gitlab.com/distributed_lab/acs/telegram-module/resources"
)

func newLink(link string, isExists bool) resources.Link {
	return resources.Link{
		Key: resources.Key{
			ID:   link,
			Type: resources.LINKS,
		},
		Attributes: resources.LinkAttributes{
			Link:     link,
			IsExists: &isExists,
		},
	}
}

func NewLinkResponse(link string, isExists bool) resources.LinkResponse {
	return resources.LinkResponse{
		Data: newLink(link, isExists),
	}
}
