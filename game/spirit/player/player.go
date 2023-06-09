package player

import (
	"Nvwa/server/nvwa_socket"
	"Nvwa/session"
)

type Player struct {
	session *session.Session
}

func NewPlayer(conn nvwa_socket.Connection) *Player {
	return nil
}
