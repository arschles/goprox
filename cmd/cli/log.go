package main

import "fmt"

func printf(fmtStr string, args ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(fmtStr, args...))
}
