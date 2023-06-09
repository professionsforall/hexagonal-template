package main

import (
	"github.com/professionsforall/hexagonal-template/cmd"
	"github.com/professionsforall/hexagonal-template/pkg/config"
	"github.com/professionsforall/hexagonal-template/pkg/log"
)

func main() {
	log.Apply()
	if err := config.Apply(); err != nil {
		log.Logger.Panicf("error while readeing apllying config: %v", err)
	}
	cmd.Execute()
}
