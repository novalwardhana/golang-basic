package main

import (
	novalAgung "github.com/novalagung/gubrak"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Part 61 Go Vendoring")

	example := novalAgung.RandomInt(10, 20)
	log.Info("example= ", example)
}
