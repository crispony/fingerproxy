package main

import (
	"log"

	"github.com/gospider007/proxy"
)

// goreleaser build --snapshot --clean
func main() {
	cli, err := proxy.NewClient(proxy.ClientOption{
		Addr: "127.0.0.1:443",
	})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Print("listening on: ", cli.Addr())
	log.Print(cli.Run())
}
