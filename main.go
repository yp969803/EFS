package main

import (
	"fmt"
	"log"

	"github.com/efs/p2p"
)

func main() {
	tr:= p2p.NewTCPTrancport(":3000")
	if err:= tr.ListenAndAcccept() ; err!= nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Hello world")
	select{}
}