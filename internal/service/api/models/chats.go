package models

import (
	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/helpers"
	"github.com/acs-dl/telegram-module-svc/resources"
)

func NewChatModel(chat data.Chat) resources.Chat {
	id, accessHash := helpers.SubmoduleIdentifiersToString(chat.Id, chat.AccessHash)

	result := resources.Chat{
		Key: resources.NewKeyInt64(chat.Id, resources.CHATS),
		Attributes: resources.ChatAttributes{
			Title:               chat.Title,
			SubmoduleId:         id,
			SubmoduleAccessHash: accessHash,
			MembersAmount:       chat.MembersAmount,
			Photo:               chat.PhotoLink,
		},
	}

	return result
}

func NewChatResponse(chat data.Chat) resources.ChatResponse {
	return resources.ChatResponse{
		Data: NewChatModel(chat),
	}
}

func NewChatListModel(chats []data.Chat) []resources.Chat {
	newChats := make([]resources.Chat, 0)

	for _, chat := range chats {
		newChats = append(newChats, NewChatModel(chat))
	}

	return newChats
}
