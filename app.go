package main

import (
	"github.com/nbanitama/chat/util"

	"github.com/nbanitama/chat/server"
)

func main() {
	util.Debug("Hello Chat")
	server.Start()
}
