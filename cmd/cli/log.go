package main

import (
	"fmt"
	"log"
)

func debugf(fmtStr string, args ...interface{}) {
	if flagDebug {
		log.Printf("%s", fmt.Sprintf(fmtStr, args...))
	}
}
func printf(fmtStr string, args ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(fmtStr, args...))
}
