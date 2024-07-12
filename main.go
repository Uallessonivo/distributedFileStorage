package main

import (
	"github/Uallessonivo/distributedFileStorage/p2p"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
