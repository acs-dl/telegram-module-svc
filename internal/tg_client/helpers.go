package tg_client

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gotd/td/tg"
)

func enter(whatToEnter string) string {
	fmt.Printf("Enter %s :", whatToEnter)
	entered, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	entered = strings.ReplaceAll(entered, "\n", "")

	return entered
}

func removeDuplicateUser(arr []tg.UserClass) []tg.User {
	allKeys := make(map[string]bool)
	var list []tg.User
	for _, item := range arr {
		if _, value := allKeys[item.(*tg.User).Username]; !value {
			allKeys[item.(*tg.User).Username] = true
			list = append(list, *item.(*tg.User))
		}
	}
	return list
}
