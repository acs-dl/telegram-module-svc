package models

import (
	"strconv"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/resources"
)

func NewUserModel(user data.User, id int) resources.User {
	result := resources.User{
		Key: resources.Key{
			ID:   strconv.Itoa(id),
			Type: resources.USER,
		},
		Attributes: resources.UserAttributes{
			UserId:      user.Id,
			Username:    user.Username,
			Phone:       user.Phone,
			Module:      data.ModuleName,
			CreatedAt:   &user.CreatedAt,
			Submodule:   user.Submodule,
			AccessLevel: &user.AccessLevel,
		},
	}

	return result
}

func NewUserResponse(user data.User) resources.UserResponse {
	return resources.UserResponse{
		Data: NewUserModel(user, 0),
	}
}

func NewUserListResponse(users []data.User, offset uint64) UserListResponse {
	return UserListResponse{
		Data: NewUsersList(users, offset),
	}
}

func NewUsersList(users []data.User, offset uint64) []resources.User {
	result := make([]resources.User, len(users))
	for i, user := range users {
		result[i] = NewUserModel(user, i+int(offset))
	}
	return result
}

type UserListResponse struct {
	Meta  Meta             `json:"meta"`
	Data  []resources.User `json:"data"`
	Links *resources.Links `json:"links"`
}

type Meta struct {
	TotalCount int64 `json:"total_count"`
}
