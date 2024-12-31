package common

import (
	"errors"
)

var InvalidPasswordErr = errors.New("invalid_password")
var UserNotExistsErr = errors.New("user_not_exists")
var ChatNotExistsErr = errors.New("chat_not_exists")
var InvalidCodeErr = errors.New("invalid_code")
