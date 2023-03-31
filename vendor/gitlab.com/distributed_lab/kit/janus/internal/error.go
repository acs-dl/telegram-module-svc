package internal

type Status int
type Response []byte

type Error struct {
	msg    string
	status Status
	body   Response
}

func (e Error) Error() string {
	return e.msg
}

func (e Error) Status() int {
	return int(e.status)
}

func (e Error) Body() []byte {
	return e.body
}

func E(msg string, args ...interface{}) error {
	e := Error{
		msg: msg,
	}

	for _, arg := range args {
		switch arg := arg.(type) {
		case Status:
			e.status = arg
		case Response:
			e.body = arg
		}
	}
	return e
}
