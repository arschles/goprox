package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func logAndErr(w http.ResponseWriter, code int, fmtStr string, vals ...interface{}) {
	str := fmt.Sprintf(fmtStr, vals...)
	log.Println(str)
	http.Error(w, str, code)
}
