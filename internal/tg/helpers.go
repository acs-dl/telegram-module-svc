package tg

import (
	"bufio"
	"fmt"
	pkgErrors "github.com/pkg/errors"
	"github.com/xelaj/mtproto"
	"github.com/xelaj/mtproto/telegram"
	"os"
	"strings"
)

func enter(whatToEnter string) string {
	fmt.Print(whatToEnter, ": ")
	entered, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	entered = strings.ReplaceAll(entered, "\n", "")

	return entered
}

func IsSessionRegistered(client *telegram.Client) (bool, error) {
	_, err := client.UsersGetFullUser(&telegram.InputUserSelf{})
	if err == nil {
		return true, nil
	}
	var errCode *mtproto.ErrResponseCode
	if pkgErrors.As(err, &errCode) {
		if errCode.Message == "AUTH_KEY_UNREGISTERED" {
			return false, nil
		}
		return false, err
	} else {
		return false, err
	}
}

func removeDuplicateUser(arr []telegram.User) []telegram.User {
	allKeys := make(map[string]bool)
	var list []telegram.User
	for _, item := range arr {
		if _, value := allKeys[item.(*telegram.UserObj).Username]; !value {
			allKeys[item.(*telegram.UserObj).Username] = true
			list = append(list, item)
		}
	}
	return list
}
