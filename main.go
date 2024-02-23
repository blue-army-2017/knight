package main

import (
	"net/http"

	"github.com/blue-army-2017/knight/controller"
	"github.com/blue-army-2017/knight/util"
)

const port string = "8080"

func main() {
	mux := controller.GetRoutesMux()

	util.LogInfo("server has started", "port", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		util.LogFatal(err.Error())
	}
}
