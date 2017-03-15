package main

import (
	"log"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}
