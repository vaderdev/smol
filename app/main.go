package main

import (
	"github.com/vaderdev/smol/app/model"
	"github.com/vaderdev/smol/app/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
