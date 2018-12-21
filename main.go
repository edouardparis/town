package main

import (
	"context"

	"git.iiens.net/edouardparis/town/web/server"
)

func main() {
	server.Run(context.Background())
}
