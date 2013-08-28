package handlers

import (
	"fmt"
)

var (
	cfg = utils.NewConfig()
)

func Nickserv(conn *irc.Conn, line *irc.Line) {
	conn.Privmsg(fmt.Sprintf("nickserv identify %s", cfg.Options["gobot"]["password"]))
}
