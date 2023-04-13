package processor

import (
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) convertUserFromInterfaceAndCheck(userInterface any) (*data.User, error) {
	user, ok := userInterface.(*data.User)
	if !ok {
		return nil, errors.Errorf("wrong response type while getting users from api")
	}
	if user == nil {
		return nil, errors.Errorf("something wrong with user from api")
	}

	return user, nil
}
