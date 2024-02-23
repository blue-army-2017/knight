package controller

import (
	"fmt"
	"net/http"
)

func getHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UP")
}
