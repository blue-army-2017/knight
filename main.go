package main

import (
	"net/http"

	"github.com/blue-army-2017/knight/controller"
	"github.com/blue-army-2017/knight/util"
	"go.uber.org/zap"
)

var l *zap.SugaredLogger

func init() {
	l = util.GetLogger()
}

func main() {
	mux, err := controller.GetRoutesMux()
	if err != nil {
		l.Fatal(err)
	}

	l.Infof("server has started on port %s", util.Config.Port)
	if err := http.ListenAndServe(":"+util.Config.Port, mux); err != nil {
		l.Fatal(err)
	}
}
