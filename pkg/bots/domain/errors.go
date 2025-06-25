package bots

import "errors"

var ErrUserNotFound = errors.New("user-not-found")
var ErrUserNotSubscribed = errors.New("user-not-subscribed")
var ErrAgentNotExists = errors.New("agent-not-exists")
