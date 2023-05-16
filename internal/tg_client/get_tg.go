package tg_client

import "github.com/gotd/td/telegram"

func (t *tgInfo) GetTg() *tgInfo {
	return t
}

func (t *tgInfo) GetSuperClient() *telegram.Client {
	return t.superUserClient
}

func (t *tgInfo) GetUsualClient() *telegram.Client {
	return t.userClient
}
