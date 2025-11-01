package commons

import "errors"

const (
	PlayerNotFound = "PlayerNotFound"
	PlayerLocked   = "PlayerLocked"
	PlayerRejected = "PlayerRejected"

	MsgTryLater = "We ran into an issue, please try later!"
)

var (
	ErrPlayerNotFound = errors.New(PlayerNotFound)
	ErrPlayerLocked   = errors.New(PlayerLocked)
	ErrPlayerRejected = errors.New(PlayerRejected)
)
