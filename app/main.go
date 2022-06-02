package main

import (
    "fmt"
    "net/http"

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

	logrus.Panic(http.ListenAndServe(":8080", router))
}
