package main

import (
    "fmt"
    "net/http"
	"flag"

    "github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

var (
	branch = "main"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprintf(w, "Welcome to %s!\n", branch)
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
	listen:= flag.String("listen", "0.0.0.0:8080", "listen address")

	logrus.Panic(http.ListenAndServe(*listen, router))
}
