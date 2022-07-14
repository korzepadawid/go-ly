package main

import (
	"log"

	"github.com/korzepadawid/go-ly/pkg/util"
)

func main() {
	log.Println("Starting server.")
	util.InitRedis()
}
