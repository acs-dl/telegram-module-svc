package tg

func (t *tg) getUserStatus(id int32, hashID *int64, userId int64) (string, error) {
	users, err := t.getAllUsers(id, hashID)
	if err != nil {
		t.log.WithError(err).Errorf("failed to get all users")
		return "", err
	}

	for _, user := range users {
		if user.TelegramId == userId {
			return user.AccessLevel, nil
		}
	}

	return "", nil
}
