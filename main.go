package main

import (
	"net/http"

	"github.com/blue-army-2017/knight/controller"
	"github.com/blue-army-2017/knight/util"
	"go.uber.org/zap"
)

const port string = "8080"

var l *zap.SugaredLogger

func init() {
	l = util.GetLogger()
}

func main() {
	mux := controller.GetRoutesMux()

	l.Infof("server has started on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		l.Fatal(err)
	}
}
