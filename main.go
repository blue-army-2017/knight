package main

import (
	"net/http"

	"github.com/blue-army-2017/knight/util"
)

const port string = "8080"

func main() {
	logger := util.GetLogger()

	mux := http.NewServeMux()

	logger.Infof("server started on port %s", port)
	logger.Fatal(http.ListenAndServe(":"+port, mux))
}
