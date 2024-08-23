package main

import (
	"log"

	"github.com/tabrizihamid84/foreverstore/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":5000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

}
