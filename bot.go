package main

import (
	gobot_handlers "github.com/but3k4/gobot/handlers"
	gobot_utils "github.com/but3k4/gobot/utils"
	irc "github.com/fluffle/goirc/client"
)

var (
	cfg = utils.NewConfig()
)

func main() {
	c := irc.SimpleClient(cfg.Options["gobot"]["nick"])
	c.EnableStateTracking()
	c.SSL = true

	c.AddHandler(irc.CONNECTED, gobot_handlers.Nickserv)
	quit := make(chan bool)
	c.AddHandler(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) { quit <- true })

	if err := c.Connect(cfg.Options["gobot"]["server"]); err != nil {
		fmt.Printf("Connection error: %s\n", err.String())
	}
	<-quit
}
