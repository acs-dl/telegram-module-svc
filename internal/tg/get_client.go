package tg

import "github.com/xelaj/mtproto/telegram"

func (t *tg) GetClient() *telegram.Client {
	return t.client
}
