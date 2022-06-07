package main

import (
    "fmt"
    "net/http"
	"flag"

    "github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

var (
	branch = "bspark"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	resp := fmt.Sprintf("Welcome to %s!", branch)
	w.Write([]byte(resp))
	w.WriteHeader(http.StatusOK)
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
	listen:= flag.String("listen", "0.0.0.0:8080", "listen address")

	logrus.Panic(http.ListenAndServe(*listen, router))
}
